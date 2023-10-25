package validator

import (
	"context"

	"github.com/rrgmc/debefix-sample-app/internal/domain/entity"
	"github.com/rrgmc/debefix-sample-app/internal/domain/validator/validatordeps"
)

func init() {
	validatordeps.Validate.RegisterStructValidationMapRules(map[string]string{
		"name": "required,min=1,max=100",
	}, entity.TagAdd{}, entity.TagUpdate{})
}

type TagValidator interface {
	ValidateTagAdd(ctx context.Context, tag entity.TagAdd) error
	ValidateTagUpdate(ctx context.Context, tag entity.TagUpdate) error
}

type tagValidator struct {
}

func NewTagValidator() TagValidator {
	return &tagValidator{}
}

func (t tagValidator) ValidateTagAdd(ctx context.Context, tag entity.TagAdd) error {
	return validatordeps.Validate.Struct(tag)
}

func (t tagValidator) ValidateTagUpdate(ctx context.Context, tag entity.TagUpdate) error {
	return validatordeps.Validate.Struct(tag)
}
