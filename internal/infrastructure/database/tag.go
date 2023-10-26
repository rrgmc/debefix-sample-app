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

type tagRepository struct {
}

func NewTagRepository() repository.TagRepository {
	return &tagRepository{}
}

func (t tagRepository) GetTagList(ctx context.Context, rctx repository.Context, filter entity.TagFilter) ([]entity.Tag, error) {
	db, err := getDB(rctx)
	if err != nil {
		return nil, err
	}

	var list []dbmodel.Tag
	result := db.
		WithContext(ctx).
		Order("name").
		Offset(filter.Offset).
		Limit(filter.Limit).
		Find(&list)
	if result.Error != nil {
		return nil, domain.NewError(domain.RepositoryError, result.Error)
	}
	return dbmodel.TagListToEntity(list), nil
}

func (t tagRepository) GetTagByID(ctx context.Context, rctx repository.Context, tagID uuid.UUID) (entity.Tag, error) {
	db, err := getDB(rctx)
	if err != nil {
		return entity.Tag{}, err
	}

	var item dbmodel.Tag
	result := db.
		WithContext(ctx).
		First(&item, tagID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return entity.Tag{}, domain.NewError(domain.NotFound)
		}
		return entity.Tag{}, domain.NewError(domain.RepositoryError, result.Error)
	}
	return item.ToEntity(), nil
}

func (t tagRepository) AddTag(ctx context.Context, rctx repository.Context, tag entity.TagAdd) (entity.Tag, error) {
	var item dbmodel.Tag

	err := doInUnitOfWork(ctx, rctx, func(db *gorm.DB) error {
		item = dbmodel.TagAddFromEntity(tag)
		item.CreatedAt = time.Now()
		item.UpdatedAt = time.Now()

		result := db.
			WithContext(ctx).
			Clauses(clause.Returning{}).
			Select("name", "created_at", "updated_at").
			Create(&item)
		if result.Error != nil {
			return domain.NewError(domain.RepositoryError, result.Error)
		}
		return nil
	})
	if err != nil {
		return entity.Tag{}, nil
	}

	return item.ToEntity(), nil
}

func (t tagRepository) UpdateTagByID(ctx context.Context, rctx repository.Context, tagID uuid.UUID, tag entity.TagUpdate) (entity.Tag, error) {
	var item dbmodel.Tag

	err := doInUnitOfWork(ctx, rctx, func(db *gorm.DB) error {
		item = dbmodel.TagUpdateFromEntity(tag)
		item.UpdatedAt = time.Now()

		result := db.
			Clauses(clause.Returning{}).
			Select("name", "updated_at").
			Where("tag_id = ?", tagID).
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
		return entity.Tag{}, err
	}

	return item.ToEntity(), nil
}

func (t tagRepository) DeleteTagByID(ctx context.Context, rctx repository.Context, tagID uuid.UUID) error {
	err := doInUnitOfWork(ctx, rctx, func(db *gorm.DB) error {
		result := db.
			Delete(dbmodel.Tag{}, tagID)
		if result.Error != nil {
			return domain.NewError(domain.RepositoryError, result.Error)
		}
		if result.RowsAffected != 1 {
			return domain.NewError(domain.NotFound)
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
