package database

import (
	"context"
	"errors"

	"github.com/rrgmc/debefix-sample-app/internal/domain"
	"github.com/rrgmc/debefix-sample-app/internal/domain/repository"
	"gorm.io/gorm"
)

type contextDB struct {
	db *gorm.DB
}

func NewContext(db *gorm.DB) repository.Context {
	return &contextDB{
		db: db,
	}
}

func (r contextDB) StartUnitOfWork(ctx context.Context, parent repository.Context) (repository.UnitOfWork, error) {
	var tx *gorm.DB
	initial := false

	switch pt := parent.(type) {
	case nil, *contextDB:
		tx = r.db.Begin(nil)
		if tx.Error != nil {
			return nil, domain.NewError(errors.Join(domain.RepositoryError, tx.Error))
		}
		initial = true
	case *unitOfWork:
		tx = pt.db
	default:
		return nil, domain.NewError(domain.RepositoryError, errors.New("incompatible repository context type"))
	}
	return &unitOfWork{contextDB{tx}, initial}, nil
}

type unitOfWork struct {
	contextDB
	initial bool
}

func (u unitOfWork) Commit(ctx context.Context) error {
	if !u.initial {
		return nil
	}
	return u.db.Commit().Error
}

func (u unitOfWork) Cancel(ctx context.Context) error {
	if !u.initial {
		return nil
	}
	return u.db.Rollback().Error
}

func getDB(rctx repository.Context) (*gorm.DB, error) {
	switch t := rctx.(type) {
	case *unitOfWork:
		return t.db, nil
	case *contextDB:
		return t.db, nil
	default:
		return nil, domain.NewError(domain.RepositoryError, errors.New("incompatible repository context type"))
	}
}

func startUnitOfWork(ctx context.Context, rctx repository.Context) (*gorm.DB, repository.UnitOfWork, error) {
	uow, err := rctx.StartUnitOfWork(ctx, rctx)
	if err != nil {
		return nil, nil, err
	}

	db, err := getDB(uow)
	if err != nil {
		return nil, nil, err
	}

	return db, uow, err
}

func doInUnitOfWork(ctx context.Context, rctx repository.Context, f func(db *gorm.DB) error) error {
	db, uow, err := startUnitOfWork(ctx, rctx)
	if err != nil {
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			_ = uow.Cancel(ctx)
			panic(r) // rethrow
		}
	}()

	err = f(db)
	if err != nil {
		_ = uow.Cancel(ctx)
	} else {
		_ = uow.Commit(ctx)
	}

	return err
}
