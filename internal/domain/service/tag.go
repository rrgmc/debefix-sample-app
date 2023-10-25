package service

import (
	"context"

	"github.com/rrgmc/debefix-sample-app/internal/domain/entity"
	"github.com/rrgmc/debefix-sample-app/internal/domain/repository"
)

type TagService interface {
	AddTag(ctx context.Context, tag entity.Tag) (entity.Tag, error)
}

type DefaultTagService struct {
	Repository repository.TagRepository
}

func (d DefaultTagService) AddTag(ctx context.Context, tag entity.Tag) (entity.Tag, error) {
	return d.Repository.AddTag(ctx, tag)
}
