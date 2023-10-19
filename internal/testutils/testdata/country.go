package testdata

import (
	"fmt"
	"strings"

	"github.com/rrgmc/debefix"
	"github.com/rrgmc/debefix-sample-app/internal/entity"
)

func GetCountryList(options ...TestDataOption) ([]entity.Country, error) {
	ret, err := filterData[entity.Country]("countries", func(row debefix.Row) (entity.Country, error) {
		return mapToStruct[entity.Country](row.Fields)
	}, func(sort string, a, b entity.Country) int {
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

func GetCountry(options ...TestDataOption) (entity.Country, error) {
	ret, err := GetCountryList(options...)
	if err != nil {
		return entity.Country{}, err
	}
	if len(ret) != 1 {
		return entity.Country{}, fmt.Errorf("incorrect amount of data returned for 'Country': expected %d got %d", 1, len(ret))
	}
	return ret[0], nil
}
