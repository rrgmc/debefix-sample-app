package fixtures

import (
	"database/sql"
	"fmt"

	"github.com/RangelReale/debefix"
	sql2 "github.com/RangelReale/debefix/db/sql"
	"github.com/RangelReale/debefix/db/sql/postgres"
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
			fmt.Printf("Loading table '%s'...\n", tableName)
		}))
}
