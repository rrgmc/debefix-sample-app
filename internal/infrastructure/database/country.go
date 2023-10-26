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
}

func NewCountryRepository() repository.CountryRepository {
	return &countryRepository{}
}

func (t countryRepository) GetCountryList(ctx context.Context, rctx repository.Context, filter entity.CountryFilter) ([]entity.Country, error) {
	db, err := getDB(rctx)
	if err != nil {
		return nil, err
	}

	var list []dbmodel.Country
	result := db.
		WithContext(ctx).
		Order("name").
		Offset(filter.Offset).
		Limit(filter.Limit).
		Find(&list)
	if result.Error != nil {
		return nil, result.Error
	}
	return dbmodel.CountryListToEntity(list), nil
}

func (t countryRepository) GetCountryByID(ctx context.Context, rctx repository.Context, countryID uuid.UUID) (entity.Country, error) {
	db, err := getDB(rctx)
	if err != nil {
		return entity.Country{}, err
	}

	var item dbmodel.Country
	result := db.WithContext(ctx).
		WithContext(ctx).
		First(&item, countryID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return entity.Country{}, domain.NewError(domain.NotFound)
		}
		return entity.Country{}, result.Error
	}
	return item.ToEntity(), nil
}
