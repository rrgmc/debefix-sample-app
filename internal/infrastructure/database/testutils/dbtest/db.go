package dbtest

import (
	"context"
	"database/sql"
	stderrors "errors"
	"fmt"
	"time"

	"github.com/docker/docker/api/types/mount"
	"github.com/docker/go-connections/nat"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pkg/errors"
	"github.com/rrgmc/debefix"
	"github.com/rrgmc/debefix-sample-app/internal/infrastructure/database/testutils/fixtures"
	"github.com/rrgmc/debefix-sample-app/internal/utils"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	pgxUUID "github.com/vgarvardt/pgx-google-uuid/v5"
)

func init() {
	testcontainers.Logger = nullLogger{}
}

// DBForTest spins up a postgres container, creates the test database on it, migrates it, and returns
// the db and a close function. The "name" parameter will become the test database name with
// a "_test" suffix as required by testfixtures.
func DBForTest(name string, opts ...DBForTestOption) (db *sql.DB, resolvedData *debefix.Data, closeFunc func() error, err error) {
	if name == "" {
		return nil, nil, nil, stderrors.New("test name is required")
	}

	var optns dbForTestOptions
	for _, opt := range opts {
		opt(&optns)
	}
	optns.fixturesTags = utils.EnsureSliceContains(optns.fixturesTags, []string{"base", "tests.base"})

	ctx := context.Background()
	// container and database
	container, db, err := CreateDBTestContainer(ctx, name, !optns.inDisk)
	if err != nil {
		return nil, nil, nil, err
	}

	closeFunc = func() error {
		return stderrors.Join(
			db.Close(),
			container.Terminate(ctx),
		)
	}

	defer func() {
		if err != nil {
			err = stderrors.Join(closeFunc(), err)
			closeFunc = nil
			db = nil
		}
	}()

	// migration
	mig, err := NewDBPGMigrator(db)
	if err != nil {
		return db, nil, closeFunc, err
	}

	err = mig.Up()
	if err != nil {
		return db, nil, closeFunc, err
	}

	// load fixtures
	resolvedData, err = fixtures.DBSeedFixtures(ctx, db,
		fixtures.WithTags(optns.fixturesTags),
		fixtures.WithOutput(optns.debugOutput),
		fixtures.WithMergeData(optns.mergeData),
	)
	if err != nil {
		return db, nil, closeFunc, err
	}

	return db, resolvedData, closeFunc, nil
}

// DBMigrationTest tests the complete migration (up and down).
func DBMigrationTest(name string) (err error) {
	db, _, closeFunc, err := DBForTest(name, WithDBForTestDebugOutput(true))
	if err != nil {
		return err
	}

	defer func() {
		err = stderrors.Join(closeFunc(), err)
	}()

	// test down migration
	mig, err := NewDBPGMigrator(db)
	if err != nil {
		return err
	}

	err = mig.Down()
	if err != nil {
		return fmt.Errorf("error in down migration: %w", err)
	}

	return nil
}

type dbForTestOptions struct {
	mergeData    []string
	fixturesTags []string
	debugOutput  bool
	inDisk       bool // if false, will create tmpfs
}

type DBForTestOption func(*dbForTestOptions)

func WithDBForTestMergeData(data []string) DBForTestOption {
	return func(o *dbForTestOptions) {
		o.mergeData = data
	}
}

func WithDBForTestFixturesTags(fixturesTags []string) DBForTestOption {
	return func(o *dbForTestOptions) {
		o.fixturesTags = fixturesTags
	}
}

func WithDBForTestDebugOutput(debugOutput bool) DBForTestOption {
	return func(o *dbForTestOptions) {
		o.debugOutput = debugOutput
	}
}

func WithDBForTestInDisk(inDisk bool) DBForTestOption {
	return func(o *dbForTestOptions) {
		o.inDisk = inDisk
	}
}

type nullLogger struct {
}

func (t nullLogger) Printf(format string, v ...interface{}) {
}

// CreateDBTestContainer spins up a Postgres database container
func CreateDBTestContainer(ctx context.Context, name string, inMemory bool) (testcontainers.Container, *sql.DB, error) {
	dbName := fmt.Sprintf("%s_test", name) // database name must end with "_test" or will be rejected

	env := map[string]string{
		"POSTGRES_PASSWORD": "password",
		"POSTGRES_USER":     "postgres",
		"POSTGRES_DB":       dbName,
	}
	dockerPort := "5432/tcp"

	var mounts testcontainers.ContainerMounts
	if inMemory {
		mounts = testcontainers.Mounts(
			testcontainers.ContainerMount{
				Source: testcontainers.DockerTmpfsMountSource{
					TmpfsOptions: &mount.TmpfsOptions{
						SizeBytes: 50 * 1024 * 1024,
						Mode:      0o644,
					},
				},
				Target: "/var/lib/postgresql/data",
			},
		)
	}

	req := testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image:        "postgres:13",
			ExposedPorts: []string{dockerPort},
			Cmd:          []string{"postgres", "-c", "fsync=off"},
			Env:          env,
			Mounts:       mounts,
			WaitingFor: wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(30 * time.Second),
		},
		Logger:  nullLogger{},
		Started: true,
	}
	container, err := testcontainers.GenericContainer(ctx, req)
	if err != nil {
		return container, nil, errors.Errorf("failed to start container: %s", err)
	}

	mappedPort, err := container.MappedPort(ctx, nat.Port(dockerPort))
	if err != nil {
		return container, nil, errors.Errorf("failed to get container external port: %s", err)
	}

	host, err := container.Host(ctx)
	if err != nil {
		return container, nil, errors.Errorf("failed to get container host: %s", err)
	}

	// log.Printf("postgres container ready and running at %s:%s\n", host, mappedPort.Port())

	url := fmt.Sprintf("postgres://postgres:password@%s:%s/%s?sslmode=disable", host, mappedPort.Port(), dbName)

	connConfig, err := pgx.ParseConfig(url)
	if err != nil {
		return container, nil, errors.Errorf("error connecting to database: %s", err)
	}

	db := stdlib.OpenDB(*connConfig, stdlib.OptionAfterConnect(func(ctx context.Context, conn *pgx.Conn) error {
		pgxUUID.Register(conn.TypeMap())
		return nil
	}))

	err = db.PingContext(ctx)
	if err != nil {
		return container, nil, errors.Errorf("error connecting to database: %s", err)
	}

	return container, db, nil
}

// NewDBPGMigrator creates a migrator instance
func NewDBPGMigrator(db *sql.DB) (*migrate.Migrate, error) {
	migrationsPath, err := MigrationsPath()
	if err != nil {
		return nil, err
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return nil, errors.Errorf("failed to create migrator driver: %s", err)
	}
	return migrate.NewWithDatabaseInstance(migrationsPath, "postgres", driver)
}
