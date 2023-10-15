package storage

import (
	"context"

	"github.com/RangelReale/debefix-sample-app/internal/entity"
)

type TagStorage interface {
	GetTags(ctx context.Context, filter *entity.TagsFilter) ([]entity.Tag, error)
}
