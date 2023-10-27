package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/rrgmc/debefix-sample-app/internal/domain/entity"
)

type CountryService interface {
	GetCountryList(ctx context.Context, filter entity.CountryFilter) ([]entity.Country, error)
	GetCountryByID(ctx context.Context, countryID uuid.UUID) (entity.Country, error)
}
