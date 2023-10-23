package dbmodel

import (
	"github.com/google/uuid"
	"github.com/rrgmc/debefix-sample-app/internal/domain/model"
)

// data

type Country struct {
	CountryID uuid.UUID `gorm:"primaryKey"`
	Name      string
}

func (Country) TableName() string {
	return "countries"
}

func (m Country) ToEntity() model.Country {
	return model.Country{
		CountryID: m.CountryID,
		Name:      m.Name,
	}
}

func CountryListToEntity(list []Country) []model.Country {
	var ret []model.Country
	for _, item := range list {
		ret = append(ret, item.ToEntity())
	}
	return ret
}

func CountryFromEntity(m model.Country) Country {
	return Country{
		CountryID: m.CountryID,
		Name:      m.Name,
	}
}
