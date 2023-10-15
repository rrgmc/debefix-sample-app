package app

import (
	"database/sql"
	"fmt"

	"github.com/RangelReale/debefix-sample-app/pkg/config"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type App struct {
	config config.Config
}

func NewApp(cfg config.Config) *App {
	return &App{
		config: cfg,
	}
}

func (a *App) Run() error {
	db, err := sql.Open("pgx", a.config.Storage.DatabaseURL)
	if err != nil {
		return fmt.Errorf("error connecting to database: %w\n", err)
	}
	defer db.Close()

	return nil
}
