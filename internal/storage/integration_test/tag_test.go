//go:build dbtest

package integration_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/rrgmc/debefix"
	"github.com/rrgmc/debefix-sample-app/internal/entity"
	"github.com/rrgmc/debefix-sample-app/internal/storage"
	"github.com/rrgmc/debefix-sample-app/internal/testutils/dbtest"
	"github.com/rrgmc/debefix-sample-app/internal/testutils/testdata"
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
		expectedTags := testdata.GetTags(
			testdata.WithFilterAll(true),
			testdata.WithSort("name"),
		)

		returnedTags, err := ts.GetTags(context.Background(), entity.TagsFilter{
			Offset: 0,
			Limit:  100,
		})
		assert.NilError(t, err)

		assert.Assert(t, is.Len(returnedTags, 3))
		assert.DeepEqual(t, expectedTags, returnedTags,
			opt.TimeWithThreshold(time.Hour))
	})
}

func TestDBTagStorageGetTagByID(t *testing.T) {
	mergeData := &debefix.Data{
		Tables: map[string]*debefix.Table{
			"tags": {
				Rows: debefix.Rows{
					{
						Config: debefix.RowConfig{
							RefID: "test.TestDBTagStorageGetTagByID",
							Tags:  []string{"base"},
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

	testDBTagStorage(t, func(db *sql.DB, resolvedData *debefix.Data, ts storage.TagStorage) {
		expectedTag := testdata.GetTag(
			testdata.WithFilterRefIDs([]string{"test.TestDBTagStorageGetTagByID"}),
			testdata.WithResolvedData(resolvedData),
		)

		returnedTag, err := ts.GetTagByID(context.Background(), expectedTag.TagID)
		assert.NilError(t, err)

		assert.DeepEqual(t, expectedTag, returnedTag,
			opt.TimeWithThreshold(time.Hour))
	}, dbtest.WithDBForTestMergeData(mergeData))
}
