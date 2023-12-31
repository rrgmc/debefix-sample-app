package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/rrgmc/debefix-sample-app/internal/domain/entity"
)

type CountryRepository interface {
	GetCountryList(ctx context.Context, rctx Context, filter entity.CountryFilter) ([]entity.Country, error)
	GetCountryByID(ctx context.Context, rctx Context, CountryID uuid.UUID) (entity.Country, error)
	ExistsCountryByID(ctx context.Context, rctx Context, CountryID uuid.UUID) (bool, error)
}
