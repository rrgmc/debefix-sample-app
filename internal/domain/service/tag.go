package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/rrgmc/debefix-sample-app/internal/domain/entity"
)

type TagService interface {
	GetTagList(ctx context.Context, filter entity.TagFilter) ([]entity.Tag, error)
	GetTagByID(ctx context.Context, tagID uuid.UUID) (entity.Tag, error)
	AddTag(ctx context.Context, tag entity.TagAdd) (entity.Tag, error)
	UpdateTagByID(ctx context.Context, tagID uuid.UUID, tag entity.TagUpdate) (entity.Tag, error)
	DeleteTagByID(ctx context.Context, tagID uuid.UUID) error
}
