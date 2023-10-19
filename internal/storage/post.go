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

type postStorage struct {
	db         *sql.DB
	q          scan.Queryer
	postMapper scan.Mapper[entity.Post]
}

func NewPostStorage(db *sql.DB) PostStorage {
	postsMapper := scan.StructMapper[entity.Post]()

	return &postStorage{
		db:         db,
		q:          bob.NewQueryer(db),
		postMapper: postsMapper,
	}
}

func getDefaultPostColumns() []any {
	return []any{"post_id", "title", "text", "user_id", "created_at", "updated_at"}
}

func (t postStorage) GetPostList(ctx context.Context, filter entity.PostFilter) ([]entity.Post, error) {
	query := psql.Select(
		sm.Columns(getColumnsWithAlias(getDefaultPostColumns(), "d")...),
		sm.From("posts").As("d"),
		sm.OrderBy(`d.title`),
		sm.Limit(filter.Limit),
		sm.Offset(filter.Offset),
	)

	queryStr, args, err := query.Build()
	if err != nil {
		return nil, err
	}

	list, err := scan.All(ctx, t.q, t.postMapper, queryStr, args...)
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (t postStorage) GetPostByID(ctx context.Context, postID uuid.UUID) (entity.Post, error) {
	query := psql.Select(
		sm.Columns(getColumnsWithAlias(getDefaultPostColumns(), "d")...),
		sm.From("posts").As("d"),
		sm.Where(psql.Raw(`d.post_id = ?`, psql.Arg(postID))),
	)

	queryStr, args, err := query.Build()
	if err != nil {
		return entity.Post{}, err
	}

	item, err := scan.One(ctx, t.q, t.postMapper, queryStr, args...)
	if err != nil {
		return entity.Post{}, err
	}

	return item, nil
}

func (t postStorage) AddPost(ctx context.Context, post entity.Post) (entity.Post, error) {
	query := psql.Insert(
		im.Into("posts", "title", "text", "user_id"),
		im.Values(psql.Arg(post.Title, post.Text, post.UserID)),
		im.Returning(getDefaultPostColumns()...),
	)

	queryStr, args, err := query.Build()
	if err != nil {
		return entity.Post{}, err
	}

	item, err := scan.One(ctx, t.q, t.postMapper, queryStr, args...)
	if err != nil {
		return entity.Post{}, err
	}

	return item, nil
}

func (t postStorage) UpdatePostByID(ctx context.Context, postID uuid.UUID, post entity.Post) (entity.Post, error) {
	query := psql.Update(
		um.Table("posts"),
		um.Set("title").To(psql.Arg(post.Title)),
		um.Set("text").To(psql.Arg(post.Text)),
		um.Set("user_id").To(psql.Arg(post.UserID)),
		um.Where(psql.Raw(`post_id = ?`, psql.Arg(postID))),
		um.Returning(getDefaultPostColumns()...),
	)

	queryStr, args, err := query.Build()
	if err != nil {
		return entity.Post{}, err
	}

	item, err := scan.One(ctx, t.q, t.postMapper, queryStr, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.Post{}, util.ErrResourceNotFound
		}
		return entity.Post{}, err
	}

	return item, nil
}

func (t postStorage) DeletePostByID(ctx context.Context, postID uuid.UUID) error {
	query := psql.Delete(
		dm.From("posts"),
		dm.Where(psql.Raw(`post_id = ?`, psql.Arg(postID))),
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
