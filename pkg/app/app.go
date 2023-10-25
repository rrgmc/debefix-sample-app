package app

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/jackc/pgx/v5"
	stdlib "github.com/jackc/pgx/v5/stdlib"
	"github.com/pkg/errors"
	"github.com/rrgmc/debefix-sample-app/internal/domain/service"
	"github.com/rrgmc/debefix-sample-app/internal/infrastructure/database"
	http2 "github.com/rrgmc/debefix-sample-app/internal/interfaces/http"
	"github.com/rrgmc/debefix-sample-app/pkg/config"
	pgxUUID "github.com/vgarvardt/pgx-google-uuid/v5"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type App struct {
	config config.Config
	db     *sql.DB
}

func NewApp(ctx context.Context, cfg config.Config) (*App, error) {
	app := &App{
		config: cfg,
	}
	return app, app.init(ctx)
}

func (a *App) init(ctx context.Context) error {
	connConfig, err := pgx.ParseConfig(a.config.Storage.DatabaseURL)
	if err != nil {
		return err
	}

	a.db = stdlib.OpenDB(*connConfig, stdlib.OptionAfterConnect(func(ctx context.Context, conn *pgx.Conn) error {
		pgxUUID.Register(conn.TypeMap())
		return nil
	}))

	err = a.db.PingContext(ctx)
	if err != nil {
		return errors.Errorf("error connecting to database: %s", err)
	}

	return nil
}

func (a *App) Run(ctx context.Context) error {
	defer a.db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: a.db,
	}), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info),
	})

	tagRepository := database.NewTagRepository(gormDB)

	tagService := service.NewTagService(tagRepository)

	httpRouter := http2.NewHTTPHandler(tagService)
	err = http.ListenAndServe(":3980", httpRouter)
	if err != nil {
		panic(err)
	}

	return nil
}
