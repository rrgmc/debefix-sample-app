package databasegorm

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/rrgmc/debefix-sample-app/internal/domain/model"
	"github.com/rrgmc/debefix-sample-app/internal/domain/repository"
	"github.com/rrgmc/debefix-sample-app/internal/infrastructure/databasegorm/internal/dbmodel"
	"github.com/rrgmc/debefix-sample-app/internal/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (t userRepository) GetUserList(ctx context.Context, filter model.UserFilter) ([]model.User, error) {
	var list []dbmodel.User
	result := t.db.WithContext(ctx).
		Order("name").
		Offset(filter.Offset).
		Limit(filter.Limit).
		Find(&list)
	if result.Error != nil {
		return nil, result.Error
	}
	return dbmodel.UserListToEntity(list), nil
}

func (t userRepository) GetUserByID(ctx context.Context, userID uuid.UUID) (model.User, error) {
	var item dbmodel.User
	result := t.db.WithContext(ctx).
		First(&item, userID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return model.User{}, utils.ErrResourceNotFound
		}
		return model.User{}, result.Error
	}
	return item.ToEntity(), nil
}

func (t userRepository) AddUser(ctx context.Context, user model.User) (model.User, error) {
	item := dbmodel.UserFromEntity(user)
	item.CreatedAt = time.Now()
	item.UpdatedAt = time.Now()

	result := t.db.
		Clauses(clause.Returning{}).
		Select("name", "email", "country_id", "created_at", "updated_at").
		Create(&item)
	if result.Error != nil {
		return model.User{}, result.Error
	}

	return item.ToEntity(), nil
}

func (t userRepository) UpdateUserByID(ctx context.Context, userID uuid.UUID, user model.User) (model.User, error) {
	item := dbmodel.UserFromEntity(user)
	item.UpdatedAt = time.Now()

	result := t.db.
		Clauses(clause.Returning{}).
		Select("name", "email", "country_id", "created_at", "updated_at").
		Where("user_id = ?", userID).
		Updates(&item)
	if result.Error != nil {
		return model.User{}, result.Error
	}
	if result.RowsAffected != 1 {
		return model.User{}, utils.ErrResourceNotFound
	}

	return item.ToEntity(), nil

}

func (t userRepository) DeleteUserByID(ctx context.Context, userID uuid.UUID) error {
	result := t.db.
		Delete(dbmodel.User{}, userID)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected != 1 {
		return utils.ErrResourceNotFound
	}
	return nil
}
