package fixtures

import (
	"database/sql"
	"fmt"

	"github.com/rrgmc/debefix"
	"github.com/rrgmc/debefix-sample-app/internal/util"
	sql2 "github.com/rrgmc/debefix/db/sql"
	"github.com/rrgmc/debefix/db/sql/postgres"
)

func DBSeedFixtures(db *sql.DB, options ...ResolveFixtureOption) (*debefix.Data, error) {
	var optns resolveFixturesOptions
	for _, opt := range options {
		opt(&optns)
	}
	optns.tags = util.EnsureSliceContains(optns.tags, []string{"base"})

	sourceData := fixtures
	if optns.mergeData != nil {
		var err error
		sourceData, err = debefix.MergeData(fixtures, optns.mergeData)
		if err != nil {
			return nil, err
		}
	}

	return postgres.Resolve(sql2.NewSQLQueryInterface(db), sourceData,
		debefix.WithResolveTags(optns.tags),
		debefix.WithResolveProgress(func(tableID, tableName string) {
			if optns.output {
				fmt.Printf("Loading table '%s'...\n", tableName)
			}
		}))
}
