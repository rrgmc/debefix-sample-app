package domain

import "errors"

type Error struct {
	Err error
}

func NewError(err error, extraErr ...error) *Error {
	if len(extraErr) == 0 {
		return &Error{err}
	}
	return &Error{errors.Join(append([]error{err}, extraErr...)...)}
}

func (e Error) Error() string {
	return e.Err.Error()
}

func (e Error) Unwrap() error {
	return e.Err
}
