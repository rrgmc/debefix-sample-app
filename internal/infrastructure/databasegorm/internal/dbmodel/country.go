package dbmodel

import (
	"github.com/google/uuid"
)

// data

type Country struct {
	CountryID uuid.UUID `gorm:"primaryKey"`
	Name      string
}

func (Country) TableName() string {
	return "country"
}
