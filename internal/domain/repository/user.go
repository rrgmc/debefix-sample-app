package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/rrgmc/debefix-sample-app/internal/domain/entity"
)

type UserRepository interface {
	GetUserList(ctx context.Context, filter entity.UserFilter) ([]entity.User, error)
	GetUserByID(ctx context.Context, userID uuid.UUID) (entity.User, error)
	AddUser(ctx context.Context, user entity.User) (entity.User, error)
	UpdateUserByID(ctx context.Context, userID uuid.UUID, user entity.User) (entity.User, error)
	DeleteUserByID(ctx context.Context, userID uuid.UUID) error
}
