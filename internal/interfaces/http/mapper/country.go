package mapper

import (
	"github.com/rrgmc/debefix-sample-app/internal/domain/entity"
	"github.com/rrgmc/debefix-sample-app/internal/interfaces/http/payload"
)

// From entity

func CountryFromEntity(country entity.Country) payload.Country {
	return payload.Country{
		CountryID: country.CountryID.String(),
		Name:      country.Name,
	}
}

func CountryListFromEntity(countryList []entity.Country) []payload.Country {
	var list []payload.Country
	for _, item := range countryList {
		list = append(list, CountryFromEntity(item))
	}
	return list
}

// To entity

func CountryFilterToEntity(countryFilter payload.CountryFilter) entity.CountryFilter {
	return entity.CountryFilter{
		Offset: countryFilter.Offset,
		Limit:  countryFilter.Limit,
	}
}
