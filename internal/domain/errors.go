package domain

import "errors"

var (
	// NotFound error indicates a missing / not found record
	NotFound = errors.New("not found")

	// ValidationError indicates an error in input validation
	ValidationError = errors.New("validation error")

	// ResourceAlreadyExists indicates a duplicate / already existing record
	ResourceAlreadyExists = errors.New("resource already exists")

	// RepositoryError indicates a repository (e.g database) error
	RepositoryError = errors.New("repository error")

	// NotAuthenticated indicates an authentication error
	NotAuthenticated = errors.New("not Authenticated")

	// NotAuthorized indicates an authorization error
	NotAuthorized = errors.New("not authorized")

	// UnknownError indicates an error that the app cannot find the cause for
	UnknownError = errors.New("unknown error")
)
