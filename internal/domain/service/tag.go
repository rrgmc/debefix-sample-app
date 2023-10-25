package service

import (
	"context"

	"github.com/rrgmc/debefix-sample-app/internal/domain/entity"
	"github.com/rrgmc/debefix-sample-app/internal/domain/repository"
)

type TagService interface {
	AddTag(ctx context.Context, tag entity.Tag) (entity.Tag, error)
}

type tagService struct {
	tagRepository repository.TagRepository
}

func NewTagService(tagRepository repository.TagRepository) TagService {
	return &tagService{
		tagRepository: tagRepository,
	}
}

func (d tagService) AddTag(ctx context.Context, tag entity.Tag) (entity.Tag, error) {
	return d.tagRepository.AddTag(ctx, tag)
}
