package mapper

import (
	"time"

	"github.com/rrgmc/debefix-sample-app/internal/domain/entity"
	"github.com/rrgmc/debefix-sample-app/internal/interfaces/http/payload"
)

// From entity

func TagFromEntity(tag entity.Tag) payload.Tag {
	return payload.Tag{
		TagID:     tag.TagID.String(),
		Name:      tag.Name,
		CreatedAt: tag.CreatedAt.Format(time.RFC3339),
		UpdatedAt: tag.UpdatedAt.Format(time.RFC3339),
	}
}

// To entity

func TagAddToEntity(tag payload.TagAdd) entity.TagAdd {
	return entity.TagAdd{
		Name: tag.Name,
	}
}

func TagUpdateToEntity(tag payload.TagUpdate) entity.TagUpdate {
	return TagAddToEntity(tag)
}
