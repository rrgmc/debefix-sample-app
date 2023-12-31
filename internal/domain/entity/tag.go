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

// data change

type TagAdd struct {
	Name string
}

type TagUpdate = TagAdd

// helpers

type TagFilter struct {
	Offset int
	Limit  int
}
