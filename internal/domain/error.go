package domain

import (
	"github.com/rrgmc/debefix-sample-app/internal/utils"
)

type Error struct {
	Err error
}

func NewError(err error, extraErr ...error) *Error {
	if len(extraErr) == 0 {
		return &Error{err}
	}
	return &Error{utils.JoinErrors(append([]error{err}, extraErr...)...)}
}

func (e Error) Error() string {
	return e.Err.Error()
}

func (e Error) Unwrap() error {
	return e.Err
}
