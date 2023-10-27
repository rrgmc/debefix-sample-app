package validator

import (
	"context"

	"github.com/rrgmc/debefix-sample-app/internal/domain"
	"github.com/rrgmc/debefix-sample-app/internal/domain/entity"
	"github.com/rrgmc/debefix-sample-app/internal/domain/validator/validatordeps"
)

func init() {
	validatordeps.Validate.RegisterStructValidationMapRules(map[string]string{
		"Offset": "min=0",
		"Limit":  "min=1,max=500",
	}, entity.CountryFilter{})
}

type CountryValidator interface {
	ValidateCountryFilter(ctx context.Context, countryFilter entity.CountryFilter) error
}

type countryValidator struct {
}

func NewCountryValidator() CountryValidator {
	return &countryValidator{}
}

func (t countryValidator) ValidateCountryFilter(ctx context.Context, countryFilter entity.CountryFilter) error {
	err := validatordeps.Validate.StructCtx(ctx, countryFilter)
	if err != nil {
		return domain.NewError(domain.ValidationError, err)
	}
	return nil
}
