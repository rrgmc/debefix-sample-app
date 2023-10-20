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

func (t Country) ToEntity() model.Country {
	return model.Country{
		CountryID: t.CountryID,
		Name:      t.Name,
	}
}

func CountryListToEntity(list []Country) []model.Country {
	var ret []model.Country
	for _, item := range list {
		ret = append(ret, item.ToEntity())
	}
	return ret
}

func CountryFromEntity(item model.Country) Country {
	return Country{
		CountryID: item.CountryID,
		Name:      item.Name,
	}
}
