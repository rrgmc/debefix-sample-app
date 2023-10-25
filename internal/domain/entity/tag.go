package entity

import (
	"time"

	"github.com/google/uuid"
)

// data

type Tag struct {
	TagID     uuid.UUID
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// helpers

type TagFilter struct {
	Offset int
	Limit  int
}

type TagChange struct {
	Name string
}
