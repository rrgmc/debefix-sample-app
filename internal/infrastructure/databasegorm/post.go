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

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) repository.PostRepository {
	return &postRepository{
		db: db,
	}
}

func (t postRepository) GetPostList(ctx context.Context, filter model.PostFilter) ([]model.Post, error) {
	var list []dbmodel.Post
	result := t.db.WithContext(ctx).
		Order("title").
		Offset(filter.Offset).
		Limit(filter.Limit).
		Find(&list)
	if result.Error != nil {
		return nil, result.Error
	}
	return dbmodel.PostListToEntity(list), nil
}

func (t postRepository) GetPostByID(ctx context.Context, postID uuid.UUID) (model.Post, error) {
	var item dbmodel.Post
	result := t.db.WithContext(ctx).
		// Model(dbmodel.Post{}).
		Preload("Tags").
		First(&item, postID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return model.Post{}, utils.ErrResourceNotFound
		}
		return model.Post{}, result.Error
	}
	return item.ToEntity(), nil
}

func (t postRepository) AddPost(ctx context.Context, post model.Post) (model.Post, error) {
	item := dbmodel.PostFromEntity(post)
	item.CreatedAt = time.Now()
	item.UpdatedAt = time.Now()

	result := t.db.
		Clauses(clause.Returning{}).
		Select("title", "text", "user_id", "created_at", "updated_at").
		Create(&item)
	if result.Error != nil {
		return model.Post{}, result.Error
	}

	return item.ToEntity(), nil
}

func (t postRepository) UpdatePostByID(ctx context.Context, postID uuid.UUID, post model.Post) (model.Post, error) {
	item := dbmodel.PostFromEntity(post)
	item.UpdatedAt = time.Now()

	result := t.db.
		Clauses(clause.Returning{}).
		Select("title", "text", "user_id", "created_at", "updated_at").
		Where("post_id = ?", postID).
		Updates(&item)
	if result.Error != nil {
		return model.Post{}, result.Error
	}
	if result.RowsAffected != 1 {
		return model.Post{}, utils.ErrResourceNotFound
	}

	return item.ToEntity(), nil

}

func (t postRepository) DeletePostByID(ctx context.Context, postID uuid.UUID) error {
	result := t.db.
		Delete(dbmodel.Post{}, postID)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected != 1 {
		return utils.ErrResourceNotFound
	}
	return nil
}
