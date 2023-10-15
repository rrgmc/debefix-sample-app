package testdata

import (
	"fmt"
	"slices"

	"github.com/RangelReale/debefix"
	"github.com/RangelReale/debefix-sample-app/internal/testutils/fixtures"
)

func resolveData[T any](refIDs []string, tableID string, f func(row debefix.Row) T, options ...TestDataOption) []T {
	optns := parseOptions(options...)

	loaded := map[string]T{}
	_, err := fixtures.
		MustResolveFixtures(fixtures.WithTags(optns.tags)).
		ExtractRows(func(table *debefix.Table, row debefix.Row) (bool, error) {
			if table.ID == tableID {
				if slices.Contains(refIDs, row.Config.RefID) {
					loaded[row.Config.RefID] = f(row)
				}
			}
			return true, nil
		})
	if err != nil {
		panic(fmt.Sprintf("error loading fixture for '%s`: %s", tableID, err))
	}

	var ret []T
	for _, refID := range refIDs {
		customer, ok := loaded[refID]
		if !ok {
			panic(fmt.Sprintf("fixture refID '%s' for '%s' not found", refID, tableID))
		}
		ret = append(ret, customer)
	}

	return ret
}
