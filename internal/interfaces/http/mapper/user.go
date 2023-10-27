package mapper

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/rrgmc/debefix-sample-app/internal/domain"
	"github.com/rrgmc/debefix-sample-app/internal/domain/entity"
	"github.com/rrgmc/debefix-sample-app/internal/interfaces/http/payload"
)

// From entity

func UserFromEntity(user entity.User) payload.User {
	return payload.User{
		UserID:    user.UserID.String(),
		Name:      user.Name,
		Email:     user.Email,
		CountryID: user.CountryID.String(),
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
	}
}

func UserListFromEntity(userList []entity.User) []payload.User {
	var list []payload.User
	for _, item := range userList {
		list = append(list, UserFromEntity(item))
	}
	return list
}

// To entity

func UserAddToEntity(user payload.UserAdd) (entity.UserAdd, error) {
	countryID, err := uuid.Parse(user.CountryID)
	if err != nil {
		return entity.UserAdd{}, domain.NewError(domain.ValidationError, errors.New("country_id: invalid uuid value"))
	}

	return entity.UserAdd{
		Name:      user.Name,
		Email:     user.Email,
		CountryID: countryID,
	}, nil
}

func UserUpdateToEntity(user payload.UserUpdate) (entity.UserUpdate, error) {
	return UserAddToEntity(user)
}

func UserFilterToEntity(userFilter payload.UserFilter) entity.UserFilter {
	return entity.UserFilter{
		Offset: userFilter.Offset,
		Limit:  userFilter.Limit,
	}
}
