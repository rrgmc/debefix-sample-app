package database

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/rrgmc/debefix-sample-app/internal/domain/model"
	"github.com/rrgmc/debefix-sample-app/internal/domain/repository"
	"github.com/rrgmc/debefix-sample-app/internal/util"
	"github.com/stephenafamo/bob"
	"github.com/stephenafamo/bob/dialect/psql"
	"github.com/stephenafamo/bob/dialect/psql/dm"
	"github.com/stephenafamo/bob/dialect/psql/im"
	"github.com/stephenafamo/bob/dialect/psql/sm"
	"github.com/stephenafamo/bob/dialect/psql/um"
	"github.com/stephenafamo/scan"
)

type tagRepository struct {
	db        *sql.DB
	q         scan.Queryer
	tagMapper scan.Mapper[model.Tag]
}

func NewTagRepository(db *sql.DB) repository.TagRepository {
	tagsMapper := scan.StructMapper[model.Tag]()

	return &tagRepository{
		db:        db,
		q:         bob.NewQueryer(db),
		tagMapper: tagsMapper,
	}
}

func getDefaultTagColumns() []any {
	return []any{"tag_id", "name", "created_at", "updated_at"}
}

func (t tagRepository) GetTagList(ctx context.Context, filter model.TagFilter) ([]model.Tag, error) {
	query := psql.Select(
		sm.Columns(getColumnsWithAlias(getDefaultTagColumns(), "d")...),
		sm.From("tags").As("d"),
		sm.OrderBy(`d.name`),
		sm.Limit(filter.Limit),
		sm.Offset(filter.Offset),
	)

	queryStr, args, err := query.Build()
	if err != nil {
		return nil, err
	}

	list, err := scan.All(ctx, t.q, t.tagMapper, queryStr, args...)
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (t tagRepository) GetTagByID(ctx context.Context, tagID uuid.UUID) (model.Tag, error) {
	query := psql.Select(
		sm.Columns(getColumnsWithAlias(getDefaultTagColumns(), "d")...),
		sm.From("tags").As("d"),
		sm.Where(psql.Raw(`d.tag_id = ?`, psql.Arg(tagID))),
	)

	queryStr, args, err := query.Build()
	if err != nil {
		return model.Tag{}, err
	}

	item, err := scan.One(ctx, t.q, t.tagMapper, queryStr, args...)
	if err != nil {
		return model.Tag{}, err
	}

	return item, nil
}

func (t tagRepository) AddTag(ctx context.Context, tag model.Tag) (model.Tag, error) {
	query := psql.Insert(
		im.Into("tags", "name"),
		im.Values(psql.Arg(tag.Name)),
		im.Returning(getDefaultTagColumns()...),
	)

	queryStr, args, err := query.Build()
	if err != nil {
		return model.Tag{}, err
	}

	item, err := scan.One(ctx, t.q, t.tagMapper, queryStr, args...)
	if err != nil {
		return model.Tag{}, err
	}

	return item, nil
}

func (t tagRepository) UpdateTagByID(ctx context.Context, tagID uuid.UUID, tag model.Tag) (model.Tag, error) {
	query := psql.Update(
		um.Table("tags"),
		um.Set("name").To(psql.Arg(tag.Name)),
		um.Where(psql.Raw(`tag_id = ?`, psql.Arg(tagID))),
		um.Returning(getDefaultTagColumns()...),
	)

	queryStr, args, err := query.Build()
	if err != nil {
		return model.Tag{}, err
	}

	item, err := scan.One(ctx, t.q, t.tagMapper, queryStr, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Tag{}, util.ErrResourceNotFound
		}
		return model.Tag{}, err
	}

	return item, nil
}

func (t tagRepository) DeleteTagByID(ctx context.Context, tagID uuid.UUID) error {
	query := psql.Delete(
		dm.From("tags"),
		dm.Where(psql.Raw(`tag_id = ?`, psql.Arg(tagID))),
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
