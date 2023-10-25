package dbmodel

import (
	"time"

	"github.com/google/uuid"
	"github.com/rrgmc/debefix-sample-app/internal/domain/entity"
)

type Tag struct {
	TagID     uuid.UUID `gorm:"primaryKey"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Tag) TableName() string {
	return "tags"
}

func (m Tag) ToEntity() entity.Tag {
	return entity.Tag{
		TagID:     m.TagID,
		Name:      m.Name,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func TagListToEntity(list []Tag) []entity.Tag {
	var ret []entity.Tag
	for _, item := range list {
		ret = append(ret, item.ToEntity())
	}
	return ret
}

func TagFromEntity(m entity.Tag) Tag {
	return Tag{
		TagID:     m.TagID,
		Name:      m.Name,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func TagChangeFromEntity(m entity.TagChange) Tag {
	return Tag{
		Name: m.Name,
	}
}
