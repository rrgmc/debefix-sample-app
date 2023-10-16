package storage

import (
	"context"

	"github.com/google/uuid"
	"github.com/rrgmc/debefix-sample-app/internal/entity"
)

type TagStorage interface {
	GetTags(ctx context.Context, filter entity.TagsFilter) ([]entity.Tag, error)
	GetTagByID(ctx context.Context, tagID uuid.UUID) (entity.Tag, error)
	AddTag(ctx context.Context, tag entity.Tag) (entity.Tag, error)
	UpdateTagByID(ctx context.Context, tagID uuid.UUID, tag entity.Tag) (entity.Tag, error)
	DeleteTagByID(ctx context.Context, tagID uuid.UUID) error
}
