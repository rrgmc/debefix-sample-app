package testdata

import (
	"fmt"

	"github.com/iancoleman/strcase"
	"github.com/mitchellh/mapstructure"
	"github.com/rrgmc/debefix"
	"github.com/rrgmc/debefix-sample-app/internal/testutils/fixtures"
	"github.com/rrgmc/debefix/filter"
)

func filterData[T any](tableID string, f func(row debefix.Row) (T, error),
	options ...TestDataOption) (result []T, sort string) {
	optns := parseOptions(options...)

	ret, err := filter.FilterData[T](fixtures.MustResolveFixtures(fixtures.WithTags(optns.resolveTags)),
		tableID, f, optns.filterDataOptions...)
	if err != nil {
		panic(fmt.Sprintf("error loading fixture for '%s`: %s", tableID, err))
	}
	return ret, optns.sort
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
