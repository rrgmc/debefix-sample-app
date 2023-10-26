package database

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/rrgmc/debefix-sample-app/internal/domain"
	"github.com/rrgmc/debefix-sample-app/internal/domain/entity"
	"github.com/rrgmc/debefix-sample-app/internal/domain/repository"
	"github.com/rrgmc/debefix-sample-app/internal/infrastructure/database/internal/dbmodel"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type userRepository struct {
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{}
}

func (t userRepository) GetUserList(ctx context.Context, rctx repository.Context, filter entity.UserFilter) ([]entity.User, error) {
	db, err := getDB(rctx)
	if err != nil {
		return nil, err
	}

	var list []dbmodel.User
	result := db.
		WithContext(ctx).
		Order("name").
		Offset(filter.Offset).
		Limit(filter.Limit).
		Find(&list)
	if result.Error != nil {
		return nil, domain.NewError(domain.RepositoryError, result.Error)
	}
	return dbmodel.UserListToEntity(list), nil
}

func (t userRepository) GetUserByID(ctx context.Context, rctx repository.Context, userID uuid.UUID) (entity.User, error) {
	db, err := getDB(rctx)
	if err != nil {
		return entity.User{}, err
	}

	var item dbmodel.User
	result := db.
		WithContext(ctx).
		First(&item, userID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return entity.User{}, domain.NewError(domain.NotFound)
		}
		return entity.User{}, domain.NewError(domain.RepositoryError, result.Error)
	}
	return item.ToEntity(), nil
}

func (t userRepository) AddUser(ctx context.Context, rctx repository.Context, user entity.User) (entity.User, error) {
	var item dbmodel.User

	err := doInUnitOfWork(ctx, rctx, func(db *gorm.DB) error {
		item = dbmodel.UserFromEntity(user)
		item.CreatedAt = time.Now()
		item.UpdatedAt = time.Now()

		result := db.
			WithContext(ctx).
			Clauses(clause.Returning{}).
			Select("name", "email", "country_id", "created_at", "updated_at").
			Create(&item)
		if result.Error != nil {
			return domain.NewError(domain.RepositoryError, result.Error)
		}
		return nil
	})
	if err != nil {
		return entity.User{}, err
	}

	return item.ToEntity(), nil
}

func (t userRepository) UpdateUserByID(ctx context.Context, rctx repository.Context, userID uuid.UUID, user entity.User) (entity.User, error) {
	var item dbmodel.User

	err := doInUnitOfWork(ctx, rctx, func(db *gorm.DB) error {
		item = dbmodel.UserFromEntity(user)
		item.UpdatedAt = time.Now()

		result := db.
			WithContext(ctx).
			Clauses(clause.Returning{}).
			Select("name", "email", "country_id", "created_at", "updated_at").
			Where("user_id = ?", userID).
			Updates(&item)
		if result.Error != nil {
			return domain.NewError(domain.RepositoryError, result.Error)
		}
		if result.RowsAffected != 1 {
			return domain.NewError(domain.NotFound)
		}
		return nil
	})
	if err != nil {
		return entity.User{}, err
	}

	return item.ToEntity(), nil

}

func (t userRepository) DeleteUserByID(ctx context.Context, rctx repository.Context, userID uuid.UUID) error {
	return doInUnitOfWork(ctx, rctx, func(db *gorm.DB) error {
		result := db.
			WithContext(ctx).
			Delete(dbmodel.User{}, userID)
		if result.Error != nil {
			return domain.NewError(domain.RepositoryError, result.Error)
		}
		if result.RowsAffected != 1 {
			return domain.NewError(domain.NotFound)
		}
		return nil
	})
}
