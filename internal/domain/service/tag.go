package service

import (
	"context"

	"github.com/rrgmc/debefix-sample-app/internal/domain/entity"
	"github.com/rrgmc/debefix-sample-app/internal/domain/repository"
	"github.com/rrgmc/debefix-sample-app/internal/domain/validator"
)

type TagService interface {
	GetTagList(ctx context.Context, filter entity.TagFilter) ([]entity.Tag, error)
	AddTag(ctx context.Context, tag entity.TagAdd) (entity.Tag, error)
}

type tagService struct {
	rctx          repository.Context
	tagRepository repository.TagRepository
	tagValidator  validator.TagValidator
}

func NewTagService(rctx repository.Context, tagRepository repository.TagRepository) TagService {
	return &tagService{
		rctx:          rctx,
		tagRepository: tagRepository,
		tagValidator:  validator.NewTagValidator(),
	}
}

func (d tagService) GetTagList(ctx context.Context, filter entity.TagFilter) ([]entity.Tag, error) {
	err := d.tagValidator.ValidateTagFilter(ctx, filter)
	if err != nil {
		return nil, err
	}
	return d.tagRepository.GetTagList(ctx, d.rctx, filter)
}

func (d tagService) AddTag(ctx context.Context, tag entity.TagAdd) (entity.Tag, error) {
	err := d.tagValidator.ValidateTagAdd(ctx, tag)
	if err != nil {
		return entity.Tag{}, err
	}
	return d.tagRepository.AddTag(ctx, d.rctx, tag)
}
