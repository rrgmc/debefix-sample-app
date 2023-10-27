package validator

import (
	"context"
	"errors"
	"fmt"

	"github.com/rrgmc/debefix-sample-app/internal/domain"
	"github.com/rrgmc/debefix-sample-app/internal/domain/entity"
	"github.com/rrgmc/debefix-sample-app/internal/domain/service"
	"github.com/rrgmc/debefix-sample-app/internal/domain/validator/validatordeps"
)

func init() {
	validatordeps.Validate.RegisterStructValidationMapRules(map[string]string{
		"Offset": "min=0",
		"Limit":  "min=1,max=500",
	}, entity.UserFilter{})
	validatordeps.Validate.RegisterStructValidationMapRules(map[string]string{
		"Name":      "required,min=3,max=100",
		"Email":     "required,email,min=3,max=200",
		"CountryID": "required",
	}, entity.UserAdd{}, entity.UserUpdate{})
}

type UserValidator interface {
	ValidateUserFilter(ctx context.Context, userFilter entity.UserFilter) error
	ValidateUserAdd(ctx context.Context, user entity.UserAdd) error
	ValidateUserUpdate(ctx context.Context, user entity.UserUpdate) error
}

type userValidator struct {
	countryService service.CountryService
}

func NewUserValidator(countryService service.CountryService) UserValidator {
	return &userValidator{
		countryService: countryService,
	}
}

func (t userValidator) ValidateUserFilter(ctx context.Context, userFilter entity.UserFilter) error {
	err := validatordeps.Validate.StructCtx(ctx, userFilter)
	if err != nil {
		return domain.NewError(domain.ValidationError, err)
	}
	return nil
}

func (t userValidator) ValidateUserAdd(ctx context.Context, user entity.UserAdd) error {
	err := validatordeps.Validate.StructCtx(ctx, user)
	if err != nil {
		return domain.NewError(domain.ValidationError, err)
	}

	found, err := t.countryService.ExistsCountryByID(ctx, user.CountryID)
	if err != nil {
		return domain.NewError(domain.ValidationError, fmt.Errorf("country_id: %w", err))
	} else if !found {
		return domain.NewError(domain.ValidationError, errors.New("country_id: not found"))
	}

	return nil
}

func (t userValidator) ValidateUserUpdate(ctx context.Context, user entity.UserUpdate) error {
	return t.ValidateUserAdd(ctx, user)
}
