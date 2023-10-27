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
	CreatedAt time.Time `gorm:"autoCreateTime:false"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:false"`

	User   *User     `gorm:"references:UserID"`
	Tags   []Tag     `gorm:"many2many:posts_tags;foreignKey:PostID;joinForeignKey:PostID;References:TagID;joinReferences:TagID"`
	TagIDs []PostTag `gorm:"foreignKey:PostID;references:PostID"`
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

func PostAddFromEntity(m entity.PostAdd) Post {
	var tags []PostTag
	for _, tag := range m.Tags {
		tags = append(tags, PostTag{
			TagID: tag,
		})
	}

	return Post{
		Title:  m.Title,
		Text:   m.Text,
		UserID: m.UserID,
		TagIDs: tags,
	}
}

func PostUpdateFromEntity(postID uuid.UUID, m entity.PostUpdate) Post {
	var tags []PostTag
	for _, tag := range m.Tags {
		tags = append(tags, PostTag{
			PostID: postID,
			TagID:  tag,
		})
	}

	return Post{
		Title:  m.Title,
		Text:   m.Text,
		UserID: m.UserID,
		TagIDs: tags,
	}
}

// PostTag

type PostTag struct {
	PostID uuid.UUID `gorm:"primaryKey"`
	TagID  uuid.UUID `gorm:"primaryKey"`
}

func (PostTag) TableName() string {
	return "posts_tags"
}
