package dbmodel

import (
	"time"

	"github.com/google/uuid"
)

// data

type Post struct {
	PostID    uuid.UUID `gorm:"primaryKey"`
	Title     string
	Text      string
	UserID    uuid.UUID
	Tags      []Tag
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Post) TableName() string {
	return "posts"
}
