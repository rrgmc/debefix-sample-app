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
		"Limit":  "min=10,max=500",
	}, entity.UserFilter{})
	validatordeps.Validate.RegisterStructValidationMapRules(map[string]string{
		"Name":      "required,min=3,max=100",
		"Email":     "required,email,min=3,max=200",
		"CountryID": "required,uuid4",
	}, entity.UserAdd{}, entity.UserUpdate{})
}

type UserValidator interface {
	ValidateUserFilter(ctx context.Context, userFilter entity.UserFilter) error
	ValidateUserAdd(ctx context.Context, user entity.UserAdd) error
	ValidateUserUpdate(ctx context.Context, user entity.UserUpdate) error
}

type userValidator struct {
}

func NewUserValidator() UserValidator {
	return &userValidator{}
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
	return nil
}

func (t userValidator) ValidateUserUpdate(ctx context.Context, user entity.UserUpdate) error {
	err := validatordeps.Validate.StructCtx(ctx, user)
	if err != nil {
		return domain.NewError(domain.ValidationError, err)
	}
	return nil
}
