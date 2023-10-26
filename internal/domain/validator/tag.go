package validator

import (
	"context"

	"github.com/rrgmc/debefix-sample-app/internal/domain"
	"github.com/rrgmc/debefix-sample-app/internal/domain/entity"
	"github.com/rrgmc/debefix-sample-app/internal/domain/validator/validatordeps"
)

func init() {
	validatordeps.Validate.RegisterStructValidationMapRules(map[string]string{
		"Offset": "required,min=0",
		"Limit":  "required,min=10,max=500",
	}, entity.TagFilter{})
	validatordeps.Validate.RegisterStructValidationMapRules(map[string]string{
		"Name": "required,min=1,max=100",
	}, entity.TagAdd{}, entity.TagUpdate{})
}

type TagValidator interface {
	ValidateTagFilter(ctx context.Context, tagFilter entity.TagFilter) error
	ValidateTagAdd(ctx context.Context, tag entity.TagAdd) error
	ValidateTagUpdate(ctx context.Context, tag entity.TagUpdate) error
}

type tagValidator struct {
}

func NewTagValidator() TagValidator {
	return &tagValidator{}
}

func (t tagValidator) ValidateTagFilter(ctx context.Context, tagFilter entity.TagFilter) error {
	err := validatordeps.Validate.StructCtx(ctx, tagFilter)
	if err != nil {
		return domain.NewError(domain.ValidationError, err)
	}
	return nil
}

func (t tagValidator) ValidateTagAdd(ctx context.Context, tag entity.TagAdd) error {
	err := validatordeps.Validate.StructCtx(ctx, tag)
	if err != nil {
		return domain.NewError(domain.ValidationError, err)
	}
	return nil
}

func (t tagValidator) ValidateTagUpdate(ctx context.Context, tag entity.TagUpdate) error {
	err := validatordeps.Validate.StructCtx(ctx, tag)
	if err != nil {
		return domain.NewError(domain.ValidationError, err)
	}
	return nil
}
