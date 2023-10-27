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
	}, entity.PostFilter{})
	validatordeps.Validate.RegisterStructValidationMapRules(map[string]string{
		"Title": "required,min=3,max=100",
		"Text":  "required,min=3,max=1000",
	}, entity.PostAdd{}, entity.PostUpdate{})
}

type PostValidator interface {
	ValidatePostFilter(ctx context.Context, postFilter entity.PostFilter) error
	ValidatePostAdd(ctx context.Context, post entity.PostAdd) error
	ValidatePostUpdate(ctx context.Context, post entity.PostUpdate) error
}

type postValidator struct {
}

func NewPostValidator() PostValidator {
	return &postValidator{}
}

func (t postValidator) ValidatePostFilter(ctx context.Context, postFilter entity.PostFilter) error {
	err := validatordeps.Validate.StructCtx(ctx, postFilter)
	if err != nil {
		return domain.NewError(domain.ValidationError, err)
	}
	return nil
}

func (t postValidator) ValidatePostAdd(ctx context.Context, post entity.PostAdd) error {
	err := validatordeps.Validate.StructCtx(ctx, post)
	if err != nil {
		return domain.NewError(domain.ValidationError, err)
	}
	return nil
}

func (t postValidator) ValidatePostUpdate(ctx context.Context, post entity.PostUpdate) error {
	err := validatordeps.Validate.StructCtx(ctx, post)
	if err != nil {
		return domain.NewError(domain.ValidationError, err)
	}
	return nil
}
