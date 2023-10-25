package testdata

import (
	"fmt"
	"strings"

	"github.com/rrgmc/debefix"
	"github.com/rrgmc/debefix-sample-app/internal/domain/entity"
	"github.com/rrgmc/debefix/filter"
)

func GetCountryList(options ...TestDataOption) (filter.FilterDataRefIDResult[entity.Country], error) {
	ret, err := filterDataRefID[entity.Country]("countries", func(row debefix.Row) (entity.Country, error) {
		return mapToStruct[entity.Country](row.Fields)
	}, func(sort string, a, b filter.FilterItem[entity.Country]) int {
		switch sort {
		case "name":
			return strings.Compare(a.Item.Name, b.Item.Name)
		default:
			panic(fmt.Sprintf("unknown sort '%s' for 'Country' testdata", sort))
		}
	}, options...)
	if err != nil {
		return filter.FilterDataRefIDResult[entity.Country]{}, err
	}
	return ret, nil
}

func GetCountry(options ...TestDataOption) (entity.Country, error) {
	ret, err := GetCountryList(options...)
	if err != nil {
		return entity.Country{}, err
	}
	if len(ret.Data) != 1 {
		return entity.Country{}, fmt.Errorf("incorrect amount of data returned for 'Country': expected %d got %d", 1, len(ret.Data))
	}
	return ret.Data[0], nil
}
