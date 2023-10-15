package storage

import (
	"context"

	"github.com/RangelReale/debefix-sample-app/internal/entity"
	"github.com/google/uuid"
)

type TagStorage interface {
	GetTags(ctx context.Context, filter entity.TagsFilter) ([]entity.Tag, error)
	GetTagByID(ctx context.Context, tagID uuid.UUID) (entity.Tag, error)
	AddTag(ctx context.Context, tag entity.Tag) error
	UpdateTagByID(ctx context.Context, tagID uuid.UUID, tag entity.Tag) error
	DeleteTagByID(ctx context.Context, tagID uuid.UUID) error
}
