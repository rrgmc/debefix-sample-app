package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pkg/errors"
	"github.com/rrgmc/debefix-sample-app/internal/infrastructure/database/testutils/fixtures"
	pgxUUID "github.com/vgarvardt/pgx-google-uuid/v5"
)

func main() {
	err := run()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error seeding data: %s\n", err)
		os.Exit(1)
	}
}

func run() error {
	connConfig, err := pgx.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		return errors.Errorf("error connecting to database: %s", err)
	}

	db := stdlib.OpenDB(*connConfig, stdlib.OptionAfterConnect(func(ctx context.Context, conn *pgx.Conn) error {
		pgxUUID.Register(conn.TypeMap())
		return nil
	}))
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return errors.Errorf("error connecting to database: %s", err)
	}

	_, err = fixtures.DBSeedFixtures(db, fixtures.WithTags([]string{"base", "seed"}))
	return err
}
