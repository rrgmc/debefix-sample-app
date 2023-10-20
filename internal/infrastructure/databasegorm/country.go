package databasegorm

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/rrgmc/debefix-sample-app/internal/domain/model"
	"github.com/rrgmc/debefix-sample-app/internal/domain/repository"
	"github.com/rrgmc/debefix-sample-app/internal/infrastructure/databasegorm/internal/dbmodel"
	"github.com/rrgmc/debefix-sample-app/internal/utils"
	"gorm.io/gorm"
)

type countryRepository struct {
	db *gorm.DB
}

func NewCountryRepository(db *gorm.DB) repository.CountryRepository {
	return &countryRepository{
		db: db,
	}
}

func (t countryRepository) GetCountryList(ctx context.Context, filter model.CountryFilter) ([]model.Country, error) {
	var list []dbmodel.Country
	result := t.db.WithContext(ctx).
		Order("name").
		Offset(filter.Offset).
		Limit(filter.Limit).
		Find(&list)
	if result.Error != nil {
		return nil, result.Error
	}
	return dbmodel.CountryListToEntity(list), nil
}

func (t countryRepository) GetCountryByID(ctx context.Context, countryID uuid.UUID) (model.Country, error) {
	var item dbmodel.Country
	result := t.db.WithContext(ctx).
		First(&item, countryID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return model.Country{}, utils.ErrResourceNotFound
		}
		return model.Country{}, result.Error
	}
	return item.ToEntity(), nil
}
