package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/rrgmc/debefix-sample-app/internal/domain/model"
)

type UserRepository interface {
	GetUserList(ctx context.Context, filter model.UserFilter) ([]model.User, error)
	GetUserByID(ctx context.Context, userID uuid.UUID) (model.User, error)
	AddUser(ctx context.Context, user model.User) (model.User, error)
	UpdateUserByID(ctx context.Context, userID uuid.UUID, user model.User) (model.User, error)
	DeleteUserByID(ctx context.Context, userID uuid.UUID) error
}
