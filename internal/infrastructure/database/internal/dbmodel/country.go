package dbmodel

import (
	"github.com/google/uuid"
	"github.com/rrgmc/debefix-sample-app/internal/domain/entity"
)

// data

type Country struct {
	CountryID uuid.UUID `gorm:"primaryKey"`
	Name      string
}

func (Country) TableName() string {
	return "countries"
}

func (m Country) ToEntity() entity.Country {
	return entity.Country{
		CountryID: m.CountryID,
		Name:      m.Name,
	}
}

func CountryListToEntity(list []Country) []entity.Country {
	var ret []entity.Country
	for _, item := range list {
		ret = append(ret, item.ToEntity())
	}
	return ret
}

func CountryFromEntity(m entity.Country) Country {
	return Country{
		CountryID: m.CountryID,
		Name:      m.Name,
	}
}
