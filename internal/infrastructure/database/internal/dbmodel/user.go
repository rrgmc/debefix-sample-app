package dbmodel

import (
	"time"

	"github.com/google/uuid"
	"github.com/rrgmc/debefix-sample-app/internal/domain/model"
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

func (m User) ToEntity() model.User {
	return model.User{
		UserID:    m.UserID,
		Name:      m.Name,
		Email:     m.Email,
		CountryID: m.CountryID,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func UserListToEntity(list []User) []model.User {
	var ret []model.User
	for _, item := range list {
		ret = append(ret, item.ToEntity())
	}
	return ret
}

func UserFromEntity(m model.User) User {
	return User{
		UserID:    m.UserID,
		Name:      m.Name,
		Email:     m.Email,
		CountryID: m.CountryID,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}
