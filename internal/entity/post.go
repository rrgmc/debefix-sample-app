package entity

import (
	"time"

	"github.com/google/uuid"
)

// data

type Post struct {
	PostID    uuid.UUID
	Title     string
	Text      string
	UserID    uuid.UUID
	Tags      []uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}

// helpers

type PostFilter struct {
	Offset int
	Limit  int
	UserID *uuid.UUID
}
