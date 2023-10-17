package storage

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/rrgmc/debefix-sample-app/internal/entity"
	"github.com/stephenafamo/bob"
	"github.com/stephenafamo/bob/dialect/psql"
	"github.com/stephenafamo/bob/dialect/psql/dm"
	"github.com/stephenafamo/bob/dialect/psql/im"
	"github.com/stephenafamo/bob/dialect/psql/sm"
	"github.com/stephenafamo/bob/dialect/psql/um"
	"github.com/stephenafamo/scan"
)

type tagStorage struct {
	db         *sql.DB
	q          scan.Queryer
	tagsMapper scan.Mapper[entity.Tag]
}

func NewTagStorage(db *sql.DB) TagStorage {
	tagsMapper := scan.StructMapper[entity.Tag]()

	return &tagStorage{
		db:         db,
		q:          bob.NewQueryer(db),
		tagsMapper: tagsMapper,
	}
}

func getDefaultTagsColumns() []any {
	return []any{"tag_id", "name", "created_at", "updated_at"}
}

func (t tagStorage) GetTags(ctx context.Context, filter entity.TagsFilter) ([]entity.Tag, error) {
	query := psql.Select(
		sm.Columns(getColumnsWithAlias(getDefaultTagsColumns(), "d")...),
		sm.From("tags").As("d"),
		sm.OrderBy(`d.name`),
		sm.Limit(filter.Limit),
		sm.Offset(filter.Offset),
	)

	queryStr, args, err := query.Build()
	if err != nil {
		return nil, err
	}

	list, err := scan.All(ctx, t.q, t.tagsMapper, queryStr, args...)
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (t tagStorage) GetTagByID(ctx context.Context, tagID uuid.UUID) (entity.Tag, error) {
	query := psql.Select(
		sm.Columns(getColumnsWithAlias(getDefaultTagsColumns(), "d")...),
		sm.From("tags").As("d"),
		sm.Where(psql.Raw(`d.tag_id = ?`, psql.Arg(tagID))),
	)

	queryStr, args, err := query.Build()
	if err != nil {
		return entity.Tag{}, err
	}

	item, err := scan.One(ctx, t.q, t.tagsMapper, queryStr, args...)
	if err != nil {
		return entity.Tag{}, err
	}

	return item, nil
}

func (t tagStorage) AddTag(ctx context.Context, tag entity.Tag) (entity.Tag, error) {
	query := psql.Insert(
		im.Into("tags", "name"),
		im.Values(psql.Arg(tag.Name)),
		im.Returning(getDefaultTagsColumns()...),
	)

	queryStr, args, err := query.Build()
	if err != nil {
		return entity.Tag{}, err
	}

	item, err := scan.One(ctx, t.q, t.tagsMapper, queryStr, args...)
	if err != nil {
		return entity.Tag{}, err
	}

	return item, nil
}

func (t tagStorage) UpdateTagByID(ctx context.Context, tagID uuid.UUID, tag entity.Tag) (entity.Tag, error) {
	query := psql.Update(
		um.Table("tags"),
		um.Set("name").To(psql.Arg(tag.Name)),
		um.Where(psql.Raw(`tag_id = ?`, psql.Arg(tagID))),
		um.Returning(getDefaultTagsColumns()...),
	)

	queryStr, args, err := query.Build()
	if err != nil {
		return entity.Tag{}, err
	}

	item, err := scan.One(ctx, t.q, t.tagsMapper, queryStr, args...)
	if err != nil {
		return entity.Tag{}, err
	}

	return item, nil
}

func (t tagStorage) DeleteTagByID(ctx context.Context, tagID uuid.UUID) error {
	query := psql.Delete(
		dm.From("tags"),
		dm.Where(psql.Raw(`tag_id = ?`, psql.Arg(tagID))),
	)

	queryStr, args, err := query.Build()
	if err != nil {
		return err
	}

	_, err = t.db.ExecContext(ctx, queryStr, args...)
	return err
}
