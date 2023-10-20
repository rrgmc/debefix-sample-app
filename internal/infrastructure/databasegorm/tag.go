package databasegorm

import (
	"context"

	"github.com/google/uuid"
	"github.com/rrgmc/debefix-sample-app/internal/domain/model"
	"github.com/rrgmc/debefix-sample-app/internal/domain/repository"
	"github.com/rrgmc/debefix-sample-app/internal/infrastructure/databasegorm/dbmodel"
	"gorm.io/gorm"
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
	return nil, nil
}

func (t tagRepository) GetTagByID(ctx context.Context, tagID uuid.UUID) (model.Tag, error) {
	var item dbmodel.Tag
	result := t.db.First(&item, tagID)
	if result.Error != nil {
		return model.Tag{}, result.Error
	}
	return item.ToEntity(), nil
}

func (t tagRepository) AddTag(ctx context.Context, tag model.Tag) (model.Tag, error) {
	return model.Tag{}, nil
}

func (t tagRepository) UpdateTagByID(ctx context.Context, tagID uuid.UUID, tag model.Tag) (model.Tag, error) {
	return model.Tag{}, nil
}

func (t tagRepository) DeleteTagByID(ctx context.Context, tagID uuid.UUID) error {
	return nil
}
