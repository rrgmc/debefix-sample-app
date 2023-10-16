package fixtures

import (
	"database/sql"
	"fmt"

	"github.com/rrgmc/debefix"
	sql2 "github.com/rrgmc/debefix/db/sql"
	"github.com/rrgmc/debefix/db/sql/postgres"
)

func DBSeedFixtures(db *sql.DB, options ...ResolveFixtureOption) error {
	optns := &resolveFixturesOptions{
		tags: []string{"base"},
	}
	for _, opt := range options {
		opt(optns)
	}

	return postgres.Resolve(sql2.NewSQLQueryInterface(db), fixtures,
		debefix.WithResolveTags(optns.tags),
		debefix.WithResolveProgress(func(tableID, tableName string) {
			if optns.output {
				fmt.Printf("Loading table '%s'...\n", tableName)
			}
		}))
}
