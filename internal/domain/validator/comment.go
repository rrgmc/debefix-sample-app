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
	}, entity.CommentFilter{})
	validatordeps.Validate.RegisterStructValidationMapRules(map[string]string{
		"PostID": "required",
		"UserID": "required",
		"Text":   "required,min=3,max=1000",
	}, entity.CommentAdd{}, entity.CommentUpdate{})
}

type CommentValidator interface {
	ValidateCommentFilter(ctx context.Context, commentFilter entity.CommentFilter) error
	ValidateCommentAdd(ctx context.Context, comment entity.CommentAdd) error
	ValidateCommentUpdate(ctx context.Context, comment entity.CommentUpdate) error
}

type commentValidator struct {
}

func NewCommentValidator() CommentValidator {
	return &commentValidator{}
}

func (t commentValidator) ValidateCommentFilter(ctx context.Context, commentFilter entity.CommentFilter) error {
	err := validatordeps.Validate.StructCtx(ctx, commentFilter)
	if err != nil {
		return domain.NewError(domain.ValidationError, err)
	}
	return nil
}

func (t commentValidator) ValidateCommentAdd(ctx context.Context, comment entity.CommentAdd) error {
	err := validatordeps.Validate.StructCtx(ctx, comment)
	if err != nil {
		return domain.NewError(domain.ValidationError, err)
	}
	return nil
}

func (t commentValidator) ValidateCommentUpdate(ctx context.Context, comment entity.CommentUpdate) error {
	err := validatordeps.Validate.StructCtx(ctx, comment)
	if err != nil {
		return domain.NewError(domain.ValidationError, err)
	}
	return nil
}
