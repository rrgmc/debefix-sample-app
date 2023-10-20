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

func (t Tag) ToEntity() model.Tag {
	return model.Tag{
		TagID:     t.TagID,
		Name:      t.Name,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}
}

func TagListToEntity(list []Tag) []model.Tag {
	var ret []model.Tag
	for _, item := range list {
		ret = append(ret, item.ToEntity())
	}
	return ret
}
