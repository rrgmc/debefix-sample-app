package storage

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/rrgmc/debefix-sample-app/internal/entity"
	"github.com/rrgmc/debefix-sample-app/internal/util"
	"github.com/stephenafamo/bob"
	"github.com/stephenafamo/bob/dialect/psql"
	"github.com/stephenafamo/bob/dialect/psql/dm"
	"github.com/stephenafamo/bob/dialect/psql/im"
	"github.com/stephenafamo/bob/dialect/psql/sm"
	"github.com/stephenafamo/bob/dialect/psql/um"
	"github.com/stephenafamo/scan"
)

type userStorage struct {
	db         *sql.DB
	q          scan.Queryer
	userMapper scan.Mapper[entity.User]
}

func NewUserStorage(db *sql.DB) UserStorage {
	usersMapper := scan.StructMapper[entity.User]()

	return &userStorage{
		db:         db,
		q:          bob.NewQueryer(db),
		userMapper: usersMapper,
	}
}

func getDefaultUserColumns() []any {
	return []any{"user_id", "name", "email", "country_id", "created_at", "updated_at"}
}

func (t userStorage) GetUserList(ctx context.Context, filter entity.UserFilter) ([]entity.User, error) {
	query := psql.Select(
		sm.Columns(getColumnsWithAlias(getDefaultUserColumns(), "d")...),
		sm.From("users").As("d"),
		sm.OrderBy(`d.name`),
		sm.Limit(filter.Limit),
		sm.Offset(filter.Offset),
	)

	queryStr, args, err := query.Build()
	if err != nil {
		return nil, err
	}

	list, err := scan.All(ctx, t.q, t.userMapper, queryStr, args...)
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (t userStorage) GetUserByID(ctx context.Context, userID uuid.UUID) (entity.User, error) {
	query := psql.Select(
		sm.Columns(getColumnsWithAlias(getDefaultUserColumns(), "d")...),
		sm.From("users").As("d"),
		sm.Where(psql.Raw(`d.user_id = ?`, psql.Arg(userID))),
	)

	queryStr, args, err := query.Build()
	if err != nil {
		return entity.User{}, err
	}

	item, err := scan.One(ctx, t.q, t.userMapper, queryStr, args...)
	if err != nil {
		return entity.User{}, err
	}

	return item, nil
}

func (t userStorage) AddUser(ctx context.Context, user entity.User) (entity.User, error) {
	query := psql.Insert(
		im.Into("users", "name", "email", "country_id"),
		im.Values(psql.Arg(user.Name, user.Email, user.CountryID)),
		im.Returning(getDefaultUserColumns()...),
	)

	queryStr, args, err := query.Build()
	if err != nil {
		return entity.User{}, err
	}

	item, err := scan.One(ctx, t.q, t.userMapper, queryStr, args...)
	if err != nil {
		return entity.User{}, err
	}

	return item, nil
}

func (t userStorage) UpdateUserByID(ctx context.Context, userID uuid.UUID, user entity.User) (entity.User, error) {
	query := psql.Update(
		um.Table("users"),
		um.Set("name").To(psql.Arg(user.Name)),
		um.Set("email").To(psql.Arg(user.Email)),
		um.Set("country_id").To(psql.Arg(user.CountryID)),
		um.Where(psql.Raw(`user_id = ?`, psql.Arg(userID))),
		um.Returning(getDefaultUserColumns()...),
	)

	queryStr, args, err := query.Build()
	if err != nil {
		return entity.User{}, err
	}

	item, err := scan.One(ctx, t.q, t.userMapper, queryStr, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.User{}, util.ErrResourceNotFound
		}
		return entity.User{}, err
	}

	return item, nil
}

func (t userStorage) DeleteUserByID(ctx context.Context, userID uuid.UUID) error {
	query := psql.Delete(
		dm.From("users"),
		dm.Where(psql.Raw(`user_id = ?`, psql.Arg(userID))),
	)

	queryStr, args, err := query.Build()
	if err != nil {
		return err
	}

	result, err := t.db.ExecContext(ctx, queryStr, args...)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected != 1 {
		return util.ErrResourceNotFound
	}
	return err
}
