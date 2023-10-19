//go:build dbtest

package integration_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/rrgmc/debefix"
	"github.com/rrgmc/debefix-sample-app/internal/entity"
	"github.com/rrgmc/debefix-sample-app/internal/storage"
	"github.com/rrgmc/debefix-sample-app/internal/testutils/dbtest"
	"github.com/rrgmc/debefix-sample-app/internal/testutils/testdata"
	"github.com/rrgmc/debefix-sample-app/internal/util"
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
	"gotest.tools/v3/assert/opt"
)

func testDBTagStorage(t *testing.T, testFn func(*sql.DB, *debefix.Data, storage.TagStorage),
	options ...dbtest.DBForTestOption) {
	db, resolvedData, dbCloseFunc, err := dbtest.DBForTest("debefix-sample-app", options...)
	assert.NilError(t, err)
	defer dbCloseFunc()

	ts := storage.NewTagStorage(db)

	testFn(db, resolvedData, ts)
}

func TestDBTagStorageGetTags(t *testing.T) {
	testDBTagStorage(t, func(db *sql.DB, resolvedData *debefix.Data, ts storage.TagStorage) {
		filter := entity.TagFilter{
			Offset: 1,
			Limit:  2,
		}

		expectedTags, err := testdata.GetTagList(
			testdata.WithFilterAll(true),
			testdata.WithSort("name"),
			testdata.WithOffsetLimit(filter.Offset, filter.Limit),
		)
		assert.NilError(t, err)
		assert.Assert(t, is.Len(expectedTags, 2))

		returnedTags, err := ts.GetTagList(context.Background(), filter)
		assert.NilError(t, err)

		assert.Assert(t, is.Len(returnedTags, 2))
		assert.DeepEqual(t, expectedTags, returnedTags,
			opt.TimeWithThreshold(time.Hour))
	})
}

func TestDBTagStorageGetTagByID(t *testing.T) {
	testDBTagStorage(t, func(db *sql.DB, resolvedData *debefix.Data, ts storage.TagStorage) {
		expectedTag, err := testdata.GetTag(
			testdata.WithFilterRefIDs([]string{"test.DBTagStorageTestMergeData"}),
			testdata.WithResolvedData(resolvedData),
		)
		assert.NilError(t, err)

		returnedTag, err := ts.GetTagByID(context.Background(), expectedTag.TagID)
		assert.NilError(t, err)

		assert.DeepEqual(t, expectedTag, returnedTag,
			opt.TimeWithThreshold(time.Hour))
	}, dbtest.WithDBForTestMergeData(dbTagStorageTestMergeData()))
}

func TestDBTagStorageAddTag(t *testing.T) {
	testDBTagStorage(t, func(db *sql.DB, resolvedData *debefix.Data, ts storage.TagStorage) {
		newTag := entity.Tag{
			Name: "new tag",
		}

		returnedTag, err := ts.AddTag(context.Background(), newTag)
		assert.NilError(t, err)
		assert.Equal(t, "new tag", returnedTag.Name)
	})
}

func TestDBTagStorageUpdateTagByID(t *testing.T) {
	testDBTagStorage(t, func(db *sql.DB, resolvedData *debefix.Data, ts storage.TagStorage) {
		expectedTag, err := testdata.GetTag(
			testdata.WithFilterRefIDs([]string{"test.DBTagStorageTestMergeData"}),
			testdata.WithResolvedData(resolvedData),
		)
		assert.NilError(t, err)

		updatedTag := entity.Tag{
			Name: "updated tag",
		}

		returnedTag, err := ts.UpdateTagByID(context.Background(), expectedTag.TagID, updatedTag)
		assert.NilError(t, err)
		assert.Equal(t, "updated tag", returnedTag.Name)
	}, dbtest.WithDBForTestMergeData(dbTagStorageTestMergeData()))
}

func TestDBTagStorageUpdateTagByIDNotFound(t *testing.T) {
	testDBTagStorage(t, func(db *sql.DB, resolvedData *debefix.Data, ts storage.TagStorage) {
		updatedTag := entity.Tag{
			Name: "updated tag",
		}

		_, err := ts.UpdateTagByID(context.Background(), uuid.MustParse("0379ca21-7ed0-45e7-8812-4a6944f2c198"), updatedTag)
		assert.ErrorIs(t, err, util.ErrResourceNotFound)
	})
}

func TestDBTagStorageDeleteTagByID(t *testing.T) {
	testDBTagStorage(t, func(db *sql.DB, resolvedData *debefix.Data, ts storage.TagStorage) {
		expectedTag, err := testdata.GetTag(
			testdata.WithFilterRefIDs([]string{"test.DBTagStorageTestMergeData"}),
			testdata.WithResolvedData(resolvedData),
		)
		assert.NilError(t, err)

		err = ts.DeleteTagByID(context.Background(), expectedTag.TagID)
		assert.NilError(t, err)
	}, dbtest.WithDBForTestMergeData(dbTagStorageTestMergeData()))
}

func TestDBTagStorageDeleteTagByIDNotFound(t *testing.T) {
	testDBTagStorage(t, func(db *sql.DB, resolvedData *debefix.Data, ts storage.TagStorage) {
		err := ts.DeleteTagByID(context.Background(), uuid.MustParse("0379ca21-7ed0-45e7-8812-4a6944f2c198"))
		assert.ErrorIs(t, err, util.ErrResourceNotFound)
	})
}

func dbTagStorageTestMergeData() *debefix.Data {
	return &debefix.Data{
		Tables: map[string]*debefix.Table{
			"tags": {
				Rows: debefix.Rows{
					{
						Config: debefix.RowConfig{
							RefID:      "test.DBTagStorageTestMergeData",
							IgnoreTags: true,
						},
						Fields: map[string]any{
							"tag_id":     &debefix.ValueGenerated{},
							"name":       "Test Tag",
							"created_at": time.Now(),
							"updated_at": time.Now(),
						},
					},
				},
			},
		},
	}
}
