package model

import (
	"time"

	"github.com/google/uuid"
)

// data

type Comment struct {
	CommentID uuid.UUID
	PostID    uuid.UUID
	UserID    uuid.UUID
	Text      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// helpers

type CommentFilter struct {
	Offset int
	Limit  int
	PostID *uuid.UUID
	UserID *uuid.UUID
}
