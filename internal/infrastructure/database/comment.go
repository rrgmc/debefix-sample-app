package database

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/rrgmc/debefix-sample-app/internal/domain/model"
	"github.com/rrgmc/debefix-sample-app/internal/domain/repository"
	"github.com/rrgmc/debefix-sample-app/internal/infrastructure/database/internal/dbmodel"
	"github.com/rrgmc/debefix-sample-app/internal/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) repository.CommentRepository {
	return &commentRepository{
		db: db,
	}
}

func (t commentRepository) GetCommentList(ctx context.Context, filter model.CommentFilter) ([]model.Comment, error) {
	var list []dbmodel.Comment
	result := t.db.WithContext(ctx).
		Order("created_at").
		Offset(filter.Offset).
		Limit(filter.Limit).
		Find(&list)
	if result.Error != nil {
		return nil, result.Error
	}
	return dbmodel.CommentListToEntity(list), nil
}

func (t commentRepository) GetCommentByID(ctx context.Context, commentID uuid.UUID) (model.Comment, error) {
	var item dbmodel.Comment
	result := t.db.WithContext(ctx).
		First(&item, commentID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return model.Comment{}, utils.ErrResourceNotFound
		}
		return model.Comment{}, result.Error
	}
	return item.ToEntity(), nil
}

func (t commentRepository) AddComment(ctx context.Context, comment model.Comment) (model.Comment, error) {
	item := dbmodel.CommentFromEntity(comment)
	item.CreatedAt = time.Now()
	item.UpdatedAt = time.Now()

	result := t.db.
		Clauses(clause.Returning{}).
		Select("post_id", "user_id", "text", "created_at", "updated_at").
		Create(&item)
	if result.Error != nil {
		return model.Comment{}, result.Error
	}

	return item.ToEntity(), nil
}

func (t commentRepository) UpdateCommentByID(ctx context.Context, commentID uuid.UUID, comment model.Comment) (model.Comment, error) {
	item := dbmodel.CommentFromEntity(comment)
	item.UpdatedAt = time.Now()

	result := t.db.
		Clauses(clause.Returning{}).
		Select("post_id", "user_id", "text", "created_at", "updated_at").
		Where("comment_id = ?", commentID).
		Updates(&item)
	if result.Error != nil {
		return model.Comment{}, result.Error
	}
	if result.RowsAffected != 1 {
		return model.Comment{}, utils.ErrResourceNotFound
	}

	return item.ToEntity(), nil

}

func (t commentRepository) DeleteCommentByID(ctx context.Context, commentID uuid.UUID) error {
	result := t.db.
		Delete(dbmodel.Comment{}, commentID)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected != 1 {
		return utils.ErrResourceNotFound
	}
	return nil
}
