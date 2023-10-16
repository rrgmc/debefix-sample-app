package testdata

import (
	"fmt"
	"slices"
	"time"

	"github.com/RangelReale/debefix"
	"github.com/RangelReale/debefix-sample-app/internal/testutils/fixtures"
	"github.com/google/go-cmp/cmp"
	"github.com/iancoleman/strcase"
	"github.com/mitchellh/mapstructure"
	"gotest.tools/v3/assert/opt"
)

func resolveData[T any](tableID string, f func(row debefix.Row) (T, error), options ...TestDataOption) []T {
	optns := parseOptions(options...)

	var ret []T
	_, err := fixtures.
		MustResolveFixtures(fixtures.WithTags(optns.resolveTags)).
		ExtractRows(func(table *debefix.Table, row debefix.Row) (bool, error) {
			if table.ID == tableID {
				include := optns.filterAll

				// filter refID
				if len(optns.filterRefIDs) > 0 {
					include = slices.Contains(optns.filterRefIDs, row.Config.RefID)
				}

				// filter fields
				if len(optns.filterFields) > 0 {
					isFilter := 0
					for filterField, filterValue := range optns.filterFields {
						fieldValue, isField := optns.filterFields[filterField]
						if !isField {
							return false, fmt.Errorf("field '%s' does not exists")
						}
						if cmp.Equal(filterValue, fieldValue, opt.TimeWithThreshold(time.Hour)) {
							isFilter++
						}
					}
					include = isFilter == len(optns.filterFields)
				}

				// filter func
				if optns.filterRow != nil {
					isRow, err := optns.filterRow(row)
					if err != nil {
						return false, fmt.Errorf("error filtering row: %s", err)
					}
					include = isRow
				}

				if include {
					data, err := f(row)
					if err != nil {
						return false, err
					}
					ret = append(ret, data)
				}
			}
			return true, nil
		})
	if err != nil {
		panic(fmt.Sprintf("error loading fixture for '%s`: %s", tableID, err))
	}

	return ret
}

func mapToStruct[T any](input any) (T, error) {
	var ret T
	dec, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Result: &ret,
		MatchName: func(mapKey, fieldName string) bool {
			return strcase.ToSnake(fieldName) == mapKey
		},
	})
	if err != nil {
		return ret, nil
	}
	err = dec.Decode(input)
	return ret, err
}
