package testdata

import (
	"fmt"
	"reflect"

	"github.com/google/uuid"
	"github.com/iancoleman/strcase"
	"github.com/mitchellh/mapstructure"
	"github.com/rrgmc/debefix"
	"github.com/rrgmc/debefix-sample-app/internal/testutils/fixtures"
	"github.com/rrgmc/debefix/filter"
)

func filterData[T any](tableID string, f func(row debefix.Row) (T, error),
	options ...TestDataOption) (result []T, sort string, err error) {
	optns := parseOptions(options...)

	ret, err := filter.FilterData[T](
		fixtures.MustResolveFixtures(
			fixtures.WithTags(optns.resolveTags),
			fixtures.WithResolvedData(optns.resolvedData)),
		tableID, f, optns.filterDataOptions...)
	if err != nil {
		return nil, "", fmt.Errorf("error loading fixture for '%s`: %s", tableID, err)
	}
	return ret, optns.sort, nil
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
