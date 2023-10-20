package databasegorm

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/rrgmc/debefix-sample-app/internal/domain/model"
	"github.com/rrgmc/debefix-sample-app/internal/domain/repository"
	"github.com/rrgmc/debefix-sample-app/internal/infrastructure/databasegorm/dbmodel"
	"github.com/rrgmc/debefix-sample-app/internal/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type tagRepository struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) repository.TagRepository {
	return &tagRepository{
		db: db,
	}
}

func (t tagRepository) GetTagList(ctx context.Context, filter model.TagFilter) ([]model.Tag, error) {
	var list []dbmodel.Tag
	result := t.db.WithContext(ctx).
		Where("").
		Order("name").
		Offset(filter.Offset).
		Limit(filter.Limit).
		Find(&list)
	if result.Error != nil {
		return nil, result.Error
	}
	return dbmodel.TagListToEntity(list), nil
}

func (t tagRepository) GetTagByID(ctx context.Context, tagID uuid.UUID) (model.Tag, error) {
	var item dbmodel.Tag
	result := t.db.WithContext(ctx).
		First(&item, tagID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return model.Tag{}, utils.ErrResourceNotFound
		}
		return model.Tag{}, result.Error
	}
	return item.ToEntity(), nil
}

func (t tagRepository) AddTag(ctx context.Context, tag model.Tag) (model.Tag, error) {
	item := dbmodel.TagFromEntity(tag)
	item.CreatedAt = time.Now()
	item.UpdatedAt = time.Now()

	result := t.db.
		Clauses(clause.Returning{}).
		Select("name", "created_at", "updated_at").
		Create(&item)
	if result.Error != nil {
		return model.Tag{}, result.Error
	}

	return item.ToEntity(), nil
}

func (t tagRepository) UpdateTagByID(ctx context.Context, tagID uuid.UUID, tag model.Tag) (model.Tag, error) {
	item := dbmodel.TagFromEntity(tag)
	item.UpdatedAt = time.Now()

	result := t.db.
		Clauses(clause.Returning{}).
		Select("name", "created_at", "updated_at").
		Where("tag_id = ?", tagID).
		Updates(&item)
	if result.Error != nil {
		return model.Tag{}, result.Error
	}
	if result.RowsAffected != 1 {
		return model.Tag{}, utils.ErrResourceNotFound
	}

	return item.ToEntity(), nil

}

func (t tagRepository) DeleteTagByID(ctx context.Context, tagID uuid.UUID) error {
	result := t.db.
		Delete(dbmodel.Tag{}, tagID)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected != 1 {
		return utils.ErrResourceNotFound
	}
	return nil
}
