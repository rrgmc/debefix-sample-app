package dbmodel

import (
	"time"

	"github.com/google/uuid"
)

// data

type User struct {
	UserID    uuid.UUID `gorm:"primaryKey"`
	Name      string
	Email     string
	CountryID uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (User) TableName() string {
	return "users"
}
