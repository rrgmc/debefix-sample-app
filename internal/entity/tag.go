package entity

import (
	"time"

	"github.com/google/uuid"
)

// data

type Tag struct {
	ID        uuid.UUID
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// helpers

type TagsFilter struct {
	Offset int
	Limit  int
}
