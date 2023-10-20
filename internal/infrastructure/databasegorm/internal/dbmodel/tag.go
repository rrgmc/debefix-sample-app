package dbmodel

import (
	"time"

	"github.com/google/uuid"
	"github.com/rrgmc/debefix-sample-app/internal/domain/model"
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

func (m Tag) ToEntity() model.Tag {
	return model.Tag{
		TagID:     m.TagID,
		Name:      m.Name,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func TagListToEntity(list []Tag) []model.Tag {
	var ret []model.Tag
	for _, item := range list {
		ret = append(ret, item.ToEntity())
	}
	return ret
}

func TagFromEntity(m model.Tag) Tag {
	return Tag{
		TagID:     m.TagID,
		Name:      m.Name,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}
