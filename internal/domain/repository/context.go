package repository

import "context"

// Context is a repository context. It could contain the db connection or a current transaction if one was started.
type Context interface {
	// StartUnitOfWork asks to start a new unit of work, usually a database transaction.
	// If no such thing is supported by the repository, return a noop.
	// Parent cannot be nil, the instance itself must be used if no parent is available.
	StartUnitOfWork(ctx context.Context, parent Context) (UnitOfWork, error)
}

// UnitOfWork signals that a series of repository actions should be saved together if supported by the repository.
// Usually this means database transactions.
type UnitOfWork interface {
	Context
	Commit(ctx context.Context) error
	Cancel(ctx context.Context) error
}
