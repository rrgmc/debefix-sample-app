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
	postService service.PostService
	userService service.UserService
}

func NewCommentValidator(postService service.PostService,
	userService service.UserService) CommentValidator {
	return &commentValidator{
		postService: postService,
		userService: userService,
	}
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

	found, err := t.postService.ExistsPostByID(ctx, comment.PostID)
	if err != nil {
		return domain.NewError(domain.ValidationError, fmt.Errorf("post_id: %w", err))
	} else if !found {
		return domain.NewError(domain.ValidationError, errors.New("post_id: not found"))
	}

	found, err = t.userService.ExistsUserByID(ctx, comment.UserID)
	if err != nil {
		return domain.NewError(domain.ValidationError, fmt.Errorf("user_id: %w", err))
	} else if !found {
		return domain.NewError(domain.ValidationError, errors.New("user_id: not found"))
	}

	return nil
}

func (t commentValidator) ValidateCommentUpdate(ctx context.Context, comment entity.CommentUpdate) error {
	return t.ValidateCommentUpdate(ctx, comment)
}
