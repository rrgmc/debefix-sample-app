package service

import (
	"context"

	"github.com/rrgmc/debefix-sample-app/internal/domain/model"
	"github.com/rrgmc/debefix-sample-app/internal/domain/repository"
)

type TagService interface {
	AddTag(ctx context.Context, tag model.Tag) (model.Tag, error)
}

type DefaultTagService struct {
	Repository repository.TagRepository
}

func (d DefaultTagService) AddTag(ctx context.Context, tag model.Tag) (model.Tag, error) {
	return d.Repository.AddTag(ctx, tag)
}
