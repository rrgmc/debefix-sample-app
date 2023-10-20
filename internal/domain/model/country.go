package model

import (
	"github.com/google/uuid"
)

// data

type Country struct {
	CountryID uuid.UUID
	Name      string
}

// helpers

type CountryFilter struct {
	Offset int
	Limit  int
}
