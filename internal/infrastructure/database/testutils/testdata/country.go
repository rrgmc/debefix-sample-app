package testdata

import (
	"fmt"
	"strings"

	"github.com/rrgmc/debefix"
	"github.com/rrgmc/debefix-sample-app/internal/domain/model"
	"github.com/rrgmc/debefix/filter"
)

func GetCountryList(options ...TestDataOption) (filter.FilterDataRefIDResult[model.Country], error) {
	ret, err := filterDataRefID[model.Country]("countries", func(row debefix.Row) (model.Country, error) {
		return mapToStruct[model.Country](row.Fields)
	}, func(sort string, a, b filter.FilterItem[model.Country]) int {
		switch sort {
		case "name":
			return strings.Compare(a.Item.Name, b.Item.Name)
		default:
			panic(fmt.Sprintf("unknown sort '%s' for 'Country' testdata", sort))
		}
	}, options...)
	if err != nil {
		return filter.FilterDataRefIDResult[model.Country]{}, err
	}
	return ret, nil
}

func GetCountry(options ...TestDataOption) (model.Country, error) {
	ret, err := GetCountryList(options...)
	if err != nil {
		return model.Country{}, err
	}
	if len(ret.Data) != 1 {
		return model.Country{}, fmt.Errorf("incorrect amount of data returned for 'Country': expected %d got %d", 1, len(ret.Data))
	}
	return ret.Data[0], nil
}
