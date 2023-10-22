package fixtures

import (
	"database/sql"
	"fmt"

	"github.com/rrgmc/debefix"
	"github.com/rrgmc/debefix-sample-app/internal/utils"
	sql2 "github.com/rrgmc/debefix/db/sql"
	"github.com/rrgmc/debefix/db/sql/postgres"
)

func DBSeedFixtures(db *sql.DB, options ...ResolveFixtureOption) (*debefix.Data, error) {
	var optns resolveFixturesOptions
	for _, opt := range options {
		opt(&optns)
	}
	optns.tags = utils.EnsureSliceContains(optns.tags, []string{"base"})

	sourceData := fixtures
	if len(optns.mergeData) > 0 {
		var err error
		sourceData, err = MergeData(sourceData, optns.mergeData)
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
