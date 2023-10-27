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

func JoinErrors(errs ...error) error {
	n := 0
	for _, err := range errs {
		if err != nil {
			n++
		}
	}
	if n == 0 {
		return nil
	}
	e := &joinError{
		errs: make([]error, 0, n),
	}
	for _, err := range errs {
		if err != nil {
			e.errs = append(e.errs, err)
		}
	}
	return e
}

type joinError struct {
	errs []error
}

func (e *joinError) Error() string {
	var b []byte
	for i, err := range e.errs {
		if i > 0 {
			b = append(b, ':', ' ')
		}
		b = append(b, err.Error()...)
	}
	return string(b)
}

func (e *joinError) Unwrap() []error {
	return e.errs
}
