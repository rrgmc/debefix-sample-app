package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/rrgmc/debefix-sample-app/internal/domain/entity"
	"github.com/rrgmc/debefix-sample-app/internal/domain/repository"
	"github.com/rrgmc/debefix-sample-app/internal/domain/validator"
)

type TagService interface {
	GetTagList(ctx context.Context, filter entity.TagFilter) ([]entity.Tag, error)
	GetTagByID(ctx context.Context, tagID uuid.UUID) (entity.Tag, error)
	AddTag(ctx context.Context, tag entity.TagAdd) (entity.Tag, error)
	UpdateTagByID(ctx context.Context, tagID uuid.UUID, tag entity.TagUpdate) (entity.Tag, error)
	DeleteTagByID(ctx context.Context, tagID uuid.UUID) error
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

func (d tagService) GetTagByID(ctx context.Context, tagID uuid.UUID) (entity.Tag, error) {
	return d.tagRepository.GetTagByID(ctx, d.rctx, tagID)
}

func (d tagService) AddTag(ctx context.Context, tag entity.TagAdd) (entity.Tag, error) {
	err := d.tagValidator.ValidateTagAdd(ctx, tag)
	if err != nil {
		return entity.Tag{}, err
	}
	return d.tagRepository.AddTag(ctx, d.rctx, tag)
}

func (d tagService) UpdateTagByID(ctx context.Context, tagID uuid.UUID, tag entity.TagUpdate) (entity.Tag, error) {
	err := d.tagValidator.ValidateTagUpdate(ctx, tag)
	if err != nil {
		return entity.Tag{}, err
	}
	return d.tagRepository.UpdateTagByID(ctx, d.rctx, tagID, tag)
}

func (d tagService) DeleteTagByID(ctx context.Context, tagID uuid.UUID) error {
	return d.tagRepository.DeleteTagByID(ctx, d.rctx, tagID)
}
