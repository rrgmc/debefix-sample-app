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

func (t User) ToEntity() model.User {
	return model.User{
		UserID:    t.UserID,
		Name:      t.Name,
		Email:     t.Email,
		CountryID: t.CountryID,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}
}

func UserListToEntity(list []User) []model.User {
	var ret []model.User
	for _, item := range list {
		ret = append(ret, item.ToEntity())
	}
	return ret
}

func UserFromEntity(item model.User) User {
	return User{
		UserID:    item.UserID,
		Name:      item.Name,
		Email:     item.Email,
		CountryID: item.CountryID,
		CreatedAt: item.CreatedAt,
		UpdatedAt: item.UpdatedAt,
	}
}
