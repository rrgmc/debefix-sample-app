package database

import (
	"context"
	"database/sql"
	"errors"

	"github.com/georgysavva/scany/v2/sqlscan"
	"github.com/google/uuid"
	"github.com/rrgmc/debefix-sample-app/internal/domain/model"
	"github.com/rrgmc/debefix-sample-app/internal/domain/repository"
	"github.com/rrgmc/debefix-sample-app/internal/util"
	"github.com/stephenafamo/bob/dialect/psql"
	"github.com/stephenafamo/bob/dialect/psql/dm"
	"github.com/stephenafamo/bob/dialect/psql/im"
	"github.com/stephenafamo/bob/dialect/psql/sm"
	"github.com/stephenafamo/bob/dialect/psql/um"
)

type postRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) repository.PostRepository {
	return &postRepository{
		db: db,
	}
}

func getDefaultPostColumns() []any {
	return []any{"post_id", "title", "text", "user_id", "created_at", "updated_at"}
}

func (t postRepository) GetPostList(ctx context.Context, filter model.PostFilter) ([]model.Post, error) {
	query := psql.Select(
		sm.Columns(getColumnsWithAlias(getDefaultPostColumns(), "d")...),
		sm.From("posts d"),
		// sm.InnerJoin("posts_tags pt").On(psql.Raw("d.post_id = pt.post_id")),
		// sm.InnerJoin("tags t").On(psql.Raw("pt.tag_id = t.tag_id")),
		sm.OrderBy(`d.title`),
		sm.Limit(filter.Limit),
		sm.Offset(filter.Offset),
	)

	if filter.UserID != nil {
		query.Apply(
			sm.Where(psql.Raw(`d.user_id = ?`, psql.Arg(*filter.UserID))),
		)
	}

	queryStr, args, err := query.Build()
	if err != nil {
		return nil, err
	}

	var list []model.Post
	err = sqlscan.Select(ctx, t.db, &list, queryStr, args...)
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (t postRepository) GetPostByID(ctx context.Context, postID uuid.UUID) (model.Post, error) {
	query := psql.Select(
		sm.Columns(getColumnsWithAlias(getDefaultPostColumns(), "d")...),
		sm.From("posts").As("d"),
		sm.Where(psql.Raw(`d.post_id = ?`, psql.Arg(postID))),
	)

	queryStr, args, err := query.Build()
	if err != nil {
		return model.Post{}, err
	}

	var item model.Post
	err = sqlscan.Get(ctx, t.db, &item, queryStr, args...)
	if err != nil {
		return model.Post{}, err
	}

	return item, nil
}

func (t postRepository) AddPost(ctx context.Context, post model.Post) (model.Post, error) {
	query := psql.Insert(
		im.Into("posts", "title", "text", "user_id"),
		im.Values(psql.Arg(post.Title, post.Text, post.UserID)),
		im.Returning(getDefaultPostColumns()...),
	)

	queryStr, args, err := query.Build()
	if err != nil {
		return model.Post{}, err
	}

	var item model.Post
	err = sqlscan.Get(ctx, t.db, &item, queryStr, args...)
	if err != nil {
		return model.Post{}, err
	}

	return item, nil
}

func (t postRepository) UpdatePostByID(ctx context.Context, postID uuid.UUID, post model.Post) (model.Post, error) {
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
		return model.Post{}, err
	}

	var item model.Post
	err = sqlscan.Get(ctx, t.db, &item, queryStr, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Post{}, util.ErrResourceNotFound
		}
		return model.Post{}, err
	}

	return item, nil
}

func (t postRepository) DeletePostByID(ctx context.Context, postID uuid.UUID) error {
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
