package testdata

import (
	"fmt"
	"reflect"
	"slices"

	"github.com/google/uuid"
	"github.com/iancoleman/strcase"
	"github.com/mitchellh/mapstructure"
	"github.com/rrgmc/debefix"
	"github.com/rrgmc/debefix-sample-app/internal/infrastructure/database/testutils/fixtures"
	"github.com/rrgmc/debefix/filter"
)

func filterData[T any](tableID string, f func(row debefix.Row) (T, error),
	sortCompare func(sort string, a, b T) int, options ...TestDataOption) ([]T, error) {
	optns := parseOptions(options...)

	var filterSortCompare func(a, b T) int
	if optns.sort != "" && sortCompare != nil {
		filterSortCompare = func(a, b T) int {
			return sortCompare(optns.sort, a, b)
		}
	}

	ret, err := filter.FilterData[T](filterDataGetData(optns),
		tableID, f, filterSortCompare, optns.filterDataOptions...)
	if err != nil {
		return nil, fmt.Errorf("error loading fixture for '%s`: %s", tableID, err)
	}
	return ret, nil
}

func filterDataRows[T any](tableID string, f func(row debefix.Row) (T, error),
	sortCompare func(sort string, a, b filter.FilterItem[T]) int, options ...TestDataOption) ([]filter.FilterItem[T], error) {
	optns := parseOptions(options...)

	var filterSortCompare func(a, b filter.FilterItem[T]) int
	if optns.sort != "" && sortCompare != nil {
		filterSortCompare = func(a, b filter.FilterItem[T]) int {
			return sortCompare(optns.sort, a, b)
		}
	}

	ret, err := filter.FilterDataRows[T](filterDataGetData(optns),
		tableID, f, filterSortCompare, optns.filterDataOptions...)
	if err != nil {
		return nil, fmt.Errorf("error loading fixture for '%s`: %s", tableID, err)
	}
	return ret, nil
}

func filterDataRefID[T any](tableID string, f func(row debefix.Row) (T, error),
	sortCompare func(sort string, a, b filter.FilterItem[T]) int, options ...TestDataOption) (filter.FilterDataRefIDResult[T], error) {
	optns := parseOptions(options...)

	var filterSortCompare func(a, b filter.FilterItem[T]) int
	if optns.sort != "" {
		if optns.sort == "refid" {
			filterSortCompare = func(a, b filter.FilterItem[T]) int {
				if slices.Index(optns.filterRefIDs, a.Row.Config.RefID) > slices.Index(optns.filterRefIDs, b.Row.Config.RefID) {
					return 1
				}
				return -1
			}
		} else if sortCompare != nil {
			filterSortCompare = func(a, b filter.FilterItem[T]) int {
				return sortCompare(optns.sort, a, b)
			}
		}
	}

	ret, err := filter.FilterDataRefID[T](filterDataGetData(optns),
		tableID, f, filterSortCompare, optns.filterDataOptions...)
	if err != nil {
		return filter.FilterDataRefIDResult[T]{}, fmt.Errorf("error loading fixture for '%s`: %s", tableID, err)
	}
	return ret, nil
}

func filterDataGetData(optns testDataOptions) *debefix.Data {
	return fixtures.MustResolveFixtures(
		fixtures.WithTags(optns.resolveTags),
		fixtures.WithResolvedData(optns.resolvedData))
}

func mapToStruct[T any](input any) (T, error) {
	var ret T
	dec, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Result: &ret,
		MatchName: func(mapKey, fieldName string) bool {
			return strcase.ToSnake(fieldName) == mapKey
		},
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			stringToUUIDHookFunc(),
		),
	})
	if err != nil {
		return ret, nil
	}
	err = dec.Decode(input)
	return ret, err
}

func stringToUUIDHookFunc() mapstructure.DecodeHookFunc {
	return func(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
		if f.Kind() != reflect.String {
			return data, nil
		}
		if t != reflect.TypeOf(uuid.UUID{}) {
			return data, nil
		}

		return uuid.Parse(data.(string))
	}
}
