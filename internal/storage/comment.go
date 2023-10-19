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

type commentStorage struct {
	db            *sql.DB
	q             scan.Queryer
	commentMapper scan.Mapper[entity.Comment]
}

func NewCommentStorage(db *sql.DB) CommentStorage {
	commentsMapper := scan.StructMapper[entity.Comment]()

	return &commentStorage{
		db:            db,
		q:             bob.NewQueryer(db),
		commentMapper: commentsMapper,
	}
}

func getDefaultCommentColumns() []any {
	return []any{"comment_id", "post_id", "user_id", "text", "created_at", "updated_at"}
}

func (t commentStorage) GetCommentList(ctx context.Context, filter entity.CommentFilter) ([]entity.Comment, error) {
	query := psql.Select(
		sm.Columns(getColumnsWithAlias(getDefaultCommentColumns(), "d")...),
		sm.From("comments").As("d"),
		sm.OrderBy(`d.created_at`),
		sm.Limit(filter.Limit),
		sm.Offset(filter.Offset),
	)

	queryStr, args, err := query.Build()
	if err != nil {
		return nil, err
	}

	list, err := scan.All(ctx, t.q, t.commentMapper, queryStr, args...)
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (t commentStorage) GetCommentByID(ctx context.Context, commentID uuid.UUID) (entity.Comment, error) {
	query := psql.Select(
		sm.Columns(getColumnsWithAlias(getDefaultCommentColumns(), "d")...),
		sm.From("comments").As("d"),
		sm.Where(psql.Raw(`d.comment_id = ?`, psql.Arg(commentID))),
	)

	queryStr, args, err := query.Build()
	if err != nil {
		return entity.Comment{}, err
	}

	item, err := scan.One(ctx, t.q, t.commentMapper, queryStr, args...)
	if err != nil {
		return entity.Comment{}, err
	}

	return item, nil
}

func (t commentStorage) AddComment(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	query := psql.Insert(
		im.Into("comments", "post_id", "user_id", "text"),
		im.Values(psql.Arg(comment.PostID, comment.UserID, comment.Text)),
		im.Returning(getDefaultCommentColumns()...),
	)

	queryStr, args, err := query.Build()
	if err != nil {
		return entity.Comment{}, err
	}

	item, err := scan.One(ctx, t.q, t.commentMapper, queryStr, args...)
	if err != nil {
		return entity.Comment{}, err
	}

	return item, nil
}

func (t commentStorage) UpdateCommentByID(ctx context.Context, commentID uuid.UUID, comment entity.Comment) (entity.Comment, error) {
	query := psql.Update(
		um.Table("comments"),
		um.Set("post_id").To(psql.Arg(comment.PostID)),
		um.Set("user_id").To(psql.Arg(comment.UserID)),
		um.Set("text").To(psql.Arg(comment.Text)),
		um.Where(psql.Raw(`comment_id = ?`, psql.Arg(commentID))),
		um.Returning(getDefaultCommentColumns()...),
	)

	queryStr, args, err := query.Build()
	if err != nil {
		return entity.Comment{}, err
	}

	item, err := scan.One(ctx, t.q, t.commentMapper, queryStr, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.Comment{}, util.ErrResourceNotFound
		}
		return entity.Comment{}, err
	}

	return item, nil
}

func (t commentStorage) DeleteCommentByID(ctx context.Context, commentID uuid.UUID) error {
	query := psql.Delete(
		dm.From("comments"),
		dm.Where(psql.Raw(`comment_id = ?`, psql.Arg(commentID))),
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
