package app

import (
	"context"

	"github.com/jackc/pgx/v5"
	stdlib "github.com/jackc/pgx/v5/stdlib"
	"github.com/pkg/errors"
	"github.com/rrgmc/debefix-sample-app/pkg/config"
	pgxUUID "github.com/vgarvardt/pgx-google-uuid/v5"
)

type App struct {
	config config.Config
}

func NewApp(cfg config.Config) *App {
	return &App{
		config: cfg,
	}
}

func (a *App) Run(ctx context.Context) error {
	connConfig, err := pgx.ParseConfig(a.config.Storage.DatabaseURL)
	if err != nil {
		return err
	}

	db := stdlib.OpenDB(*connConfig, stdlib.OptionAfterConnect(func(ctx context.Context, conn *pgx.Conn) error {
		pgxUUID.Register(conn.TypeMap())
		return nil
	}))
	defer db.Close()

	err = db.PingContext(ctx)
	if err != nil {
		return errors.Errorf("error connecting to database: %s", err)
	}

	// db, err := sql.Open("pgx", a.config.Storage.DatabaseURL)
	// if err != nil {
	// 	return fmt.Errorf("error connecting to database: %w\n", err)
	// }
	// defer db.Close()

	return nil
}
