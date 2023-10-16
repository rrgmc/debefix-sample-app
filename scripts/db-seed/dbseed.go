package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/rrgmc/debefix-sample-app/internal/testutils/fixtures"
)

func main() {
	err := run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error seeding data: %s", err)
		os.Exit(1)
	}
}

func run() error {
	db, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		return fmt.Errorf("error connecting to database: %w\n", err)
	}
	defer db.Close()

	return fixtures.DBSeedFixtures(db)
}
