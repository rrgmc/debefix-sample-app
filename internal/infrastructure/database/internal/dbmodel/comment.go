package dbmodel

import (
	"time"

	"github.com/google/uuid"
	"github.com/rrgmc/debefix-sample-app/internal/domain/model"
)

// data

type Comment struct {
	CommentID uuid.UUID `gorm:"primaryKey"`
	PostID    uuid.UUID
	UserID    uuid.UUID
	Text      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Comment) TableName() string {
	return "comments"
}

func (m Comment) ToEntity() model.Comment {
	return model.Comment{
		CommentID: m.CommentID,
		PostID:    m.PostID,
		UserID:    m.UserID,
		Text:      m.Text,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func CommentListToEntity(list []Comment) []model.Comment {
	var ret []model.Comment
	for _, item := range list {
		ret = append(ret, item.ToEntity())
	}
	return ret
}

func CommentFromEntity(m model.Comment) Comment {
	return Comment{
		CommentID: m.CommentID,
		PostID:    m.PostID,
		UserID:    m.UserID,
		Text:      m.Text,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}
