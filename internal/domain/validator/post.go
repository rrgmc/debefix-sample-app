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
	tagService  service.TagService
	userService service.UserService
}

func NewPostValidator(tagService service.TagService, userService service.UserService) PostValidator {
	return &postValidator{
		tagService:  tagService,
		userService: userService,
	}
}

func (t postValidator) ValidatePostFilter(ctx context.Context, postFilter entity.PostFilter) error {
	err := validatordeps.Validate.StructCtx(ctx, postFilter)
	if err != nil {
		return domain.NewError(domain.ValidationError, err)
	}

	if postFilter.UserID != nil {
		found, err := t.userService.ExistsUserByID(ctx, *postFilter.UserID)
		if err != nil {
			return domain.NewError(domain.ValidationError, fmt.Errorf("user_id: %w", err))
		} else if !found {
			return domain.NewError(domain.ValidationError, errors.New("user_id: not found"))
		}
	}

	return nil
}

func (t postValidator) ValidatePostAdd(ctx context.Context, post entity.PostAdd) error {
	err := validatordeps.Validate.StructCtx(ctx, post)
	if err != nil {
		return domain.NewError(domain.ValidationError, err)
	}

	found, err := t.userService.ExistsUserByID(ctx, post.UserID)
	if err != nil {
		return domain.NewError(domain.ValidationError, fmt.Errorf("user_id: %w", err))
	} else if !found {
		return domain.NewError(domain.ValidationError, errors.New("user_id: not found"))
	}

	for _, tag := range post.Tags {
		_, err = t.tagService.GetTagByID(ctx, tag)
		if err != nil {
			return domain.NewError(domain.ValidationError, fmt.Errorf("tag_id: %w", err))
		}
	}

	return nil
}

func (t postValidator) ValidatePostUpdate(ctx context.Context, post entity.PostUpdate) error {
	return t.ValidatePostAdd(ctx, post)
}
