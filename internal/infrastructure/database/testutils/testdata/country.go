package testdata

import (
	"fmt"
	"strings"

	"github.com/rrgmc/debefix"
	"github.com/rrgmc/debefix-sample-app/internal/domain/model"
)

func GetCountryList(options ...TestDataOption) ([]model.Country, error) {
	ret, err := filterData[model.Country]("countries", func(row debefix.Row) (model.Country, error) {
		return mapToStruct[model.Country](row.Fields)
	}, func(sort string, a, b model.Country) int {
		switch sort {
		case "name":
			return strings.Compare(a.Name, b.Name)
		default:
			panic(fmt.Sprintf("unknown sort '%s' for 'Country' testdata", sort))
		}
	}, options...)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func GetCountry(options ...TestDataOption) (model.Country, error) {
	ret, err := GetCountryList(options...)
	if err != nil {
		return model.Country{}, err
	}
	if len(ret) != 1 {
		return model.Country{}, fmt.Errorf("incorrect amount of data returned for 'Country': expected %d got %d", 1, len(ret))
	}
	return ret[0], nil
}
