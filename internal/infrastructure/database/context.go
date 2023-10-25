package database

import (
	gocontext "context"
	"errors"

	"github.com/rrgmc/debefix-sample-app/internal/domain"
	"github.com/rrgmc/debefix-sample-app/internal/domain/repository"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type contextGoRM struct {
	db *gorm.DB
}

func NewContext(db *gorm.DB) repository.Context {
	return &contextGoRM{
		db: db,
	}
}

func (r contextGoRM) StartUnitOfWork(ctx gocontext.Context, parent repository.Context) (repository.UnitOfWork, error) {
	var tx *gorm.DB
	initial := false

	parentUow, ok := parent.(*unitOfWork)
	if parent == nil || !ok {
		tx = r.db.Begin(nil)
		initial = true
	} else {
		if !ok {
			return nil, domain.NewError(errors.Join(domain.RepositoryError, errors.New("incompatible repository context type")))
		}
		tx = parentUow.db
	}
	return &unitOfWork{contextGoRM{tx}, initial}, nil
}

type unitOfWork struct {
	contextGoRM
	initial bool
}

func (u unitOfWork) Commit(ctx gocontext.Context) error {
	if !u.initial {
		return nil
	}
	return u.db.Commit().Error
}

func (u unitOfWork) Cancel(ctx gocontext.Context) error {
	if !u.initial {
		return nil
	}
	return u.db.Rollback().Error
}

func getDB(rctx repository.Context) (*gorm.DB, error) {
	switch t := rctx.(type) {
	case unitOfWork:
		return t.db, nil
	case contextGoRM:
		return t.db, nil
	default:
		return nil, domain.NewError(errors.Join(domain.RepositoryError, errors.New("incompatible repository context type")))
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

	err = f(db)
	if err != nil {
		_ = uow.Cancel(ctx)
	} else {
		_ = uow.Commit(ctx)
	}

	return err
}
