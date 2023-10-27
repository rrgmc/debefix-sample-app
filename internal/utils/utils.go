package utils

import (
	"errors"

	"github.com/google/uuid"
)

func Ptr[T any](v T) *T {
	return &v
}

func UUIDListParse(list []string) ([]uuid.UUID, error) {
	var ret []uuid.UUID
	for _, item := range list {
		u, err := uuid.Parse(item)
		if err != nil {
			return nil, errors.New("invalid uuid value")
		}
		ret = append(ret, u)
	}
	return ret, nil
}
