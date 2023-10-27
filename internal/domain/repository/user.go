package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/rrgmc/debefix-sample-app/internal/domain/entity"
)

type UserRepository interface {
	GetUserList(ctx context.Context, rctx Context, filter entity.UserFilter) ([]entity.User, error)
	GetUserByID(ctx context.Context, rctx Context, userID uuid.UUID) (entity.User, error)
	AddUser(ctx context.Context, rctx Context, user entity.UserAdd) (entity.User, error)
	UpdateUserByID(ctx context.Context, rctx Context, userID uuid.UUID, user entity.UserUpdate) (entity.User, error)
	DeleteUserByID(ctx context.Context, rctx Context, userID uuid.UUID) error
}
