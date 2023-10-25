package dbmodel

import (
	"time"

	"github.com/google/uuid"
	"github.com/rrgmc/debefix-sample-app/internal/domain/entity"
)

// data

type Post struct {
	PostID    uuid.UUID `gorm:"primaryKey"`
	Title     string
	Text      string
	UserID    uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time

	User *User `gorm:"references:UserID"`
	Tags []Tag `gorm:"many2many:posts_tags;foreignKey:PostID;joinForeignKey:PostID;References:TagID;joinReferences:TagID"`
}

func (Post) TableName() string {
	return "posts"
}

func (m Post) ToEntity() entity.Post {
	var tags []entity.Tag
	for _, tag := range m.Tags {
		tags = append(tags, tag.ToEntity())
	}

	return entity.Post{
		PostID:    m.PostID,
		Title:     m.Title,
		Text:      m.Text,
		UserID:    m.UserID,
		Tags:      tags,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func PostListToEntity(list []Post) []entity.Post {
	var ret []entity.Post
	for _, item := range list {
		ret = append(ret, item.ToEntity())
	}
	return ret
}

func PostFromEntity(m entity.Post) Post {
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
