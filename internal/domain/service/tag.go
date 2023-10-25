package service

import (
	"context"

	"github.com/rrgmc/debefix-sample-app/internal/domain/entity"
	"github.com/rrgmc/debefix-sample-app/internal/domain/repository"
	"github.com/rrgmc/debefix-sample-app/internal/domain/validator"
)

type TagService interface {
	AddTag(ctx context.Context, tag entity.TagAdd) (entity.Tag, error)
}

type tagService struct {
	tagRepository repository.TagRepository
	tagValidator  validator.TagValidator
}

func NewTagService(tagRepository repository.TagRepository) TagService {
	return &tagService{
		tagRepository: tagRepository,
		tagValidator:  validator.NewTagValidator(),
	}
}

func (d tagService) AddTag(ctx context.Context, tag entity.TagAdd) (entity.Tag, error) {
	err := d.tagValidator.ValidateTagAdd(ctx, tag)
	if err != nil {
		return entity.Tag{}, err
	}
	return d.tagRepository.AddTag(ctx, tag)
}
