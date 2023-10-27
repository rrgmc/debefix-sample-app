package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/rrgmc/debefix-sample-app/internal/domain/entity"
	"github.com/rrgmc/debefix-sample-app/internal/domain/repository"
	"github.com/rrgmc/debefix-sample-app/internal/domain/validator"
)

type CountryService interface {
	GetCountryList(ctx context.Context, filter entity.CountryFilter) ([]entity.Country, error)
	GetCountryByID(ctx context.Context, countryID uuid.UUID) (entity.Country, error)
}

type countryService struct {
	rctx              repository.Context
	countryRepository repository.CountryRepository
	countryValidator  validator.CountryValidator
}

func NewCountryService(rctx repository.Context, countryRepository repository.CountryRepository) CountryService {
	return &countryService{
		rctx:              rctx,
		countryRepository: countryRepository,
		countryValidator:  validator.NewCountryValidator(),
	}
}

func (d countryService) GetCountryList(ctx context.Context, filter entity.CountryFilter) ([]entity.Country, error) {
	err := d.countryValidator.ValidateCountryFilter(ctx, filter)
	if err != nil {
		return nil, err
	}
	return d.countryRepository.GetCountryList(ctx, d.rctx, filter)
}

func (d countryService) GetCountryByID(ctx context.Context, countryID uuid.UUID) (entity.Country, error) {
	return d.countryRepository.GetCountryByID(ctx, d.rctx, countryID)
}
