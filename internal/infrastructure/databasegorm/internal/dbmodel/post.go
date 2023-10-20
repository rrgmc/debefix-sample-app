package dbmodel

import (
	"time"

	"github.com/google/uuid"
	"github.com/rrgmc/debefix-sample-app/internal/domain/model"
)

// data

type Post struct {
	PostID    uuid.UUID `gorm:"primaryKey"`
	Title     string
	Text      string
	UserID    uuid.UUID
	Tags      []Tag `gorm:"many2many:posts_tags;joinForeignKey:post_id"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Post) TableName() string {
	return "posts"
}

func (m Post) ToEntity() model.Post {
	var tags []model.Tag
	for _, tag := range m.Tags {
		tags = append(tags, tag.ToEntity())
	}

	return model.Post{
		PostID:    m.PostID,
		Title:     m.Title,
		Text:      m.Text,
		UserID:    m.UserID,
		Tags:      tags,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func PostListToEntity(list []Post) []model.Post {
	var ret []model.Post
	for _, item := range list {
		ret = append(ret, item.ToEntity())
	}
	return ret
}

func PostFromEntity(m model.Post) Post {
	var tags []Tag
	for _, tag := range m.Tags {
		tags = append(tags, TagFromEntity(tag))
	}

	return Post{
		PostID:    m.PostID,
		Title:     m.Title,
		Text:      m.Text,
		UserID:    m.UserID,
		Tags:      tags,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}
