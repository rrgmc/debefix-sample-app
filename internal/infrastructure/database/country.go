package database

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/rrgmc/debefix-sample-app/internal/domain"
	"github.com/rrgmc/debefix-sample-app/internal/domain/entity"
	"github.com/rrgmc/debefix-sample-app/internal/domain/repository"
	"github.com/rrgmc/debefix-sample-app/internal/infrastructure/database/internal/dbmodel"
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

func (t countryRepository) GetCountryList(ctx context.Context, filter entity.CountryFilter) ([]entity.Country, error) {
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

func (t countryRepository) GetCountryByID(ctx context.Context, countryID uuid.UUID) (entity.Country, error) {
	var item dbmodel.Country
	result := t.db.WithContext(ctx).
		First(&item, countryID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return entity.Country{}, domain.NewError(domain.NotFound)
		}
		return entity.Country{}, result.Error
	}
	return item.ToEntity(), nil
}
