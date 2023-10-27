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

type commentRepository struct {
}

func NewCommentRepository() repository.CommentRepository {
	return &commentRepository{}
}

func (t commentRepository) GetCommentList(ctx context.Context, rctx repository.Context, filter entity.CommentFilter) ([]entity.Comment, error) {
	db, err := getDB(rctx)
	if err != nil {
		return nil, err
	}

	var list []dbmodel.Comment
	q := db.
		WithContext(ctx).
		Order("created_at").
		Offset(filter.Offset).
		Limit(filter.Limit)

	if filter.PostID != nil {
		q = q.Where(`post_id = ?`, *filter.PostID)
	}
	if filter.UserID != nil {
		q = q.Where(`user_id = ?`, *filter.UserID)
	}

	result := q.
		Find(&list)
	if result.Error != nil {
		return nil, domain.NewError(domain.RepositoryError, result.Error)
	}
	return dbmodel.CommentListToEntity(list), nil
}

func (t commentRepository) GetCommentByID(ctx context.Context, rctx repository.Context, commentID uuid.UUID) (entity.Comment, error) {
	db, err := getDB(rctx)
	if err != nil {
		return entity.Comment{}, err
	}

	var item dbmodel.Comment
	result := db.
		WithContext(ctx).
		First(&item, commentID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return entity.Comment{}, domain.NewError(domain.NotFound)
		}
		return entity.Comment{}, domain.NewError(domain.RepositoryError, result.Error)
	}
	return item.ToEntity(), nil
}

func (t commentRepository) ExistsCommentByID(ctx context.Context, rctx repository.Context, commentID uuid.UUID) (bool, error) {
	db, err := getDB(rctx)
	if err != nil {
		return false, err
	}

	var item dbmodel.Comment
	result := db.
		WithContext(ctx).
		Select("comment_id").
		First(&item, commentID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, domain.NewError(domain.RepositoryError, result.Error)
	}
	return true, nil
}

func (t commentRepository) AddComment(ctx context.Context, rctx repository.Context, comment entity.CommentAdd) (entity.Comment, error) {
	var item dbmodel.Comment

	err := doInUnitOfWork(ctx, rctx, func(urctx repository.Context, db *gorm.DB) error {
		item = dbmodel.CommentAddFromEntity(comment)
		item.CreatedAt = time.Now()
		item.UpdatedAt = time.Now()

		result := db.
			WithContext(ctx).
			Clauses(clause.Returning{}).
			Select("post_id", "user_id", "text", "created_at", "updated_at").
			Create(&item)
		if result.Error != nil {
			return domain.NewError(domain.RepositoryError, result.Error)
		}
		return nil
	})
	if err != nil {
		return entity.Comment{}, err
	}

	return item.ToEntity(), nil
}

func (t commentRepository) UpdateCommentByID(ctx context.Context, rctx repository.Context,
	commentID uuid.UUID, comment entity.CommentUpdate) (entity.Comment, error) {
	var item dbmodel.Comment

	err := doInUnitOfWork(ctx, rctx, func(urctx repository.Context, db *gorm.DB) error {
		item = dbmodel.CommentUpdateFromEntity(comment)
		item.UpdatedAt = time.Now()

		result := db.
			WithContext(ctx).
			Clauses(clause.Returning{}).
			Select("post_id", "user_id", "text", "updated_at").
			Where("comment_id = ?", commentID).
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
		return entity.Comment{}, err
	}

	return item.ToEntity(), nil

}

func (t commentRepository) DeleteCommentByID(ctx context.Context, rctx repository.Context, commentID uuid.UUID) error {
	return doInUnitOfWork(ctx, rctx, func(urctx repository.Context, db *gorm.DB) error {
		result := db.
			WithContext(ctx).
			Delete(dbmodel.Comment{}, commentID)
		if result.Error != nil {
			return domain.NewError(domain.RepositoryError, result.Error)
		}
		if result.RowsAffected != 1 {
			return domain.NewError(domain.NotFound)
		}
		return nil
	})
}
