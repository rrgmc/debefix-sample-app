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

type postRepository struct {
}

func NewPostRepository() repository.PostRepository {
	return &postRepository{}
}

func (t postRepository) GetPostList(ctx context.Context, rctx repository.Context, filter entity.PostFilter) ([]entity.Post, error) {
	db, err := getDB(rctx)
	if err != nil {
		return nil, err
	}

	var list []dbmodel.Post
	result := db.
		WithContext(ctx).
		Preload("Tags", func(db *gorm.DB) *gorm.DB {
			return db.Order("tags.name")
		}).
		Order("title").
		Offset(filter.Offset).
		Limit(filter.Limit).
		Find(&list)
	if result.Error != nil {
		return nil, domain.NewError(domain.RepositoryError, result.Error)
	}
	return dbmodel.PostListToEntity(list), nil
}

func (t postRepository) GetPostByID(ctx context.Context, rctx repository.Context, postID uuid.UUID) (entity.Post, error) {
	db, err := getDB(rctx)
	if err != nil {
		return entity.Post{}, err
	}

	var item dbmodel.Post
	result := db.
		WithContext(ctx).
		Preload("Tags", func(db *gorm.DB) *gorm.DB {
			return db.Order("tags.name")
		}).
		First(&item, postID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return entity.Post{}, domain.NewError(domain.NotFound)
		}
		return entity.Post{}, domain.NewError(domain.RepositoryError, result.Error)
	}
	return item.ToEntity(), nil
}

func (t postRepository) AddPost(ctx context.Context, rctx repository.Context, post entity.Post) (entity.Post, error) {
	var item dbmodel.Post

	err := doInUnitOfWork(ctx, rctx, func(db *gorm.DB) error {
		item = dbmodel.PostFromEntity(post)
		item.CreatedAt = time.Now()
		item.UpdatedAt = time.Now()

		result := db.
			WithContext(ctx).
			Clauses(clause.Returning{}).
			Select("title", "text", "user_id", "created_at", "updated_at", "Tags").
			Create(&item)
		if result.Error != nil {
			return domain.NewError(domain.RepositoryError, result.Error)
		}
		return nil
	})
	if err != nil {
		return entity.Post{}, err
	}

	return item.ToEntity(), nil
}

func (t postRepository) UpdatePostByID(ctx context.Context, rctx repository.Context, postID uuid.UUID, post entity.Post) (entity.Post, error) {
	var item dbmodel.Post

	err := doInUnitOfWork(ctx, rctx, func(db *gorm.DB) error {
		item = dbmodel.PostFromEntity(post)
		item.UpdatedAt = time.Now()

		result := db.
			Clauses(clause.Returning{}).
			Select("title", "text", "user_id", "created_at", "updated_at").
			Where("post_id = ?", postID).
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
		return entity.Post{}, err
	}

	return item.ToEntity(), nil
}

func (t postRepository) DeletePostByID(ctx context.Context, rctx repository.Context, postID uuid.UUID) error {
	return doInUnitOfWork(ctx, rctx, func(db *gorm.DB) error {
		result := db.
			Delete(dbmodel.Post{}, postID)
		if result.Error != nil {
			return domain.NewError(domain.RepositoryError, result.Error)
		}
		if result.RowsAffected != 1 {
			return domain.NewError(domain.NotFound)
		}
		return nil
	})
}
