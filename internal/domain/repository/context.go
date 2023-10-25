package repository

import "context"

type Context interface {
	StartUnitOfWork(ctx context.Context, parent Context) (UnitOfWork, error)
}

type UnitOfWork interface {
	Context
	Commit(ctx context.Context) error
	Cancel(ctx context.Context) error
}
