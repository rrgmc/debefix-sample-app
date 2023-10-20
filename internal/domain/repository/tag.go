package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/rrgmc/debefix-sample-app/internal/domain/model"
)

type TagRepository interface {
	GetTagList(ctx context.Context, filter model.TagFilter) ([]model.Tag, error)
	GetTagByID(ctx context.Context, tagID uuid.UUID) (model.Tag, error)
	AddTag(ctx context.Context, tag model.Tag) (model.Tag, error)
	UpdateTagByID(ctx context.Context, tagID uuid.UUID, tag model.Tag) (model.Tag, error)
	DeleteTagByID(ctx context.Context, tagID uuid.UUID) error
}
