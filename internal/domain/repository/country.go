package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/rrgmc/debefix-sample-app/internal/domain/model"
)

type CountryRepository interface {
	GetCountryList(ctx context.Context, filter model.CountryFilter) ([]model.Country, error)
	GetCountryByID(ctx context.Context, CountryID uuid.UUID) (model.Country, error)
}
