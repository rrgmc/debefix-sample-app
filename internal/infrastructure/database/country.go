package database

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/rrgmc/debefix-sample-app/internal/domain/model"
	"github.com/rrgmc/debefix-sample-app/internal/domain/repository"
	"github.com/stephenafamo/bob"
	"github.com/stephenafamo/bob/dialect/psql"
	"github.com/stephenafamo/bob/dialect/psql/sm"
	"github.com/stephenafamo/scan"
)

type countryRepository struct {
	db            *sql.DB
	q             scan.Queryer
	countryMapper scan.Mapper[model.Country]
}

func NewCountryRepository(db *sql.DB) repository.CountryRepository {
	countrysMapper := scan.StructMapper[model.Country]()

	return &countryRepository{
		db:            db,
		q:             bob.NewQueryer(db),
		countryMapper: countrysMapper,
	}
}

func getDefaultCountryColumns() []any {
	return []any{"country_id", "name"}
}

func (t countryRepository) GetCountryList(ctx context.Context, filter model.CountryFilter) ([]model.Country, error) {
	query := psql.Select(
		sm.Columns(getColumnsWithAlias(getDefaultCountryColumns(), "d")...),
		sm.From("countries").As("d"),
		sm.OrderBy(`d.name`),
		sm.Limit(filter.Limit),
		sm.Offset(filter.Offset),
	)

	queryStr, args, err := query.Build()
	if err != nil {
		return nil, err
	}

	list, err := scan.All(ctx, t.q, t.countryMapper, queryStr, args...)
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (t countryRepository) GetCountryByID(ctx context.Context, countryID uuid.UUID) (model.Country, error) {
	query := psql.Select(
		sm.Columns(getColumnsWithAlias(getDefaultCountryColumns(), "d")...),
		sm.From("countries").As("d"),
		sm.Where(psql.Raw(`d.country_id = ?`, psql.Arg(countryID))),
	)

	queryStr, args, err := query.Build()
	if err != nil {
		return model.Country{}, err
	}

	item, err := scan.One(ctx, t.q, t.countryMapper, queryStr, args...)
	if err != nil {
		return model.Country{}, err
	}

	return item, nil
}
