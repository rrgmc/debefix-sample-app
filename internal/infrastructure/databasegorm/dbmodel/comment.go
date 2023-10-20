package dbmodel

import (
	"time"

	"github.com/google/uuid"
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
