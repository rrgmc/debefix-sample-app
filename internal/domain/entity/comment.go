package entity

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

// data change

type CommentAdd struct {
	PostID uuid.UUID
	UserID uuid.UUID
	Text   string
}

type CommentUpdate = CommentAdd

// helpers

type CommentFilter struct {
	Offset int
	Limit  int
	PostID *uuid.UUID
	UserID *uuid.UUID
}
