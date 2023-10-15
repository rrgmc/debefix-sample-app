package dbtest

import (
	"context"
	"database/sql"
	stderrors "errors"
	"fmt"
	"log"
	"time"

	"github.com/docker/go-connections/nat"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/pkg/errors"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

// DBForTest spins up a postgres container, creates the test database on it, migrates it, and returns
// the db and a close function. The "name" parameter will become the test database name with
// a "_test" suffix as required by testfixtures.
func DBForTest(name string, opts ...DBForTestOption) (db *sql.DB, closeFunc func() error, err error) {
	if name == "" {
		return nil, nil, stderrors.New("test name is required")
	}

	var optns dbForTestOptions
	for _, opt := range opts {
		opt(&optns)
	}

	ctx := context.Background()
	// container and database
	container, db, err := CreateDBTestContainer(ctx, name)
	if err != nil {
		return nil, nil, err
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
		return db, closeFunc, err
	}

	err = mig.Up()
	if err != nil {
		return db, closeFunc, err
	}

	// TODO: load fixtures

	return db, closeFunc, nil
}

// DBMigrationTest tests the complete migration (up and down).
func DBMigrationTest(name string) (err error) {
	db, closeFunc, err := DBForTest(name)
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
	fixturesPath string
}

type DBForTestOption func(*dbForTestOptions)

func WithDBForTestFixturesPath(fixturesPath string) DBForTestOption {
	return func(o *dbForTestOptions) {
		o.fixturesPath = fixturesPath
	}
}

// CreateDBTestContainer spins up a Postgres database container
func CreateDBTestContainer(ctx context.Context, name string) (testcontainers.Container, *sql.DB, error) {
	dbName := fmt.Sprintf("%s_test", name) // database name must end with "_test" or will be rejected

	env := map[string]string{
		"POSTGRES_PASSWORD": "password",
		"POSTGRES_USER":     "postgres",
		"POSTGRES_DB":       dbName,
	}
	dockerPort := "5432/tcp"

	req := testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image:        "postgres:13",
			ExposedPorts: []string{dockerPort},
			Cmd:          []string{"postgres", "-c", "fsync=off"},
			Env:          env,
			WaitingFor: wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(30 * time.Second),
		},
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

	log.Printf("postgres container ready and running at %s:%s\n", host, mappedPort.Port())

	url := fmt.Sprintf("postgres://postgres:password@%s:%s/%s?sslmode=disable", host, mappedPort.Port(), dbName)
	db, err := sql.Open("postgres", url)
	if err != nil {
		return container, db, errors.Errorf("failed to establish database connection: %s", err)
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
