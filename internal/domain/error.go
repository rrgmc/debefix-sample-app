package domain

type Error struct {
	Err error
}

func NewError(err error) *Error {
	return &Error{err}
}

func (e Error) Error() string {
	return e.Err.Error()
}

func (e Error) Unwrap() error {
	return e.Err
}
