package entity

import (
	"time"

	"github.com/google/uuid"
)

// data

type User struct {
	UserID    uuid.UUID
	Name      string
	Email     string
	CountryID uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}

// data change

type UserAdd struct {
	Name      string
	Email     string
	CountryID uuid.UUID
}

type UserUpdate = UserAdd

// helpers

type UserFilter struct {
	Offset int
	Limit  int
}
