package validator

import (
	"context"

	"github.com/rrgmc/debefix-sample-app/internal/domain/entity"
)

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
	return nil
}

func (t tagValidator) ValidateTagUpdate(ctx context.Context, tag entity.TagUpdate) error {
	return nil
}
