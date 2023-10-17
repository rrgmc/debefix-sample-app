//go:build dbtest

package integration_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/rrgmc/debefix-sample-app/internal/entity"
	"github.com/rrgmc/debefix-sample-app/internal/storage"
	"github.com/rrgmc/debefix-sample-app/internal/testutils/dbtest"
	"github.com/rrgmc/debefix-sample-app/internal/testutils/testdata"
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
	"gotest.tools/v3/assert/opt"
)

func testDBTagStorage(t *testing.T, testFn func(*sql.DB, storage.TagStorage)) {
	db, dbCloseFunc, err := dbtest.DBForTest("debefix-sample-app")
	assert.NilError(t, err)
	defer dbCloseFunc()

	ts := storage.NewTagStorage(db)

	testFn(db, ts)
}

func TestDBTagStorageGetTags(t *testing.T) {
	testDBTagStorage(t, func(db *sql.DB, ts storage.TagStorage) {
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
	testDBTagStorage(t, func(db *sql.DB, ts storage.TagStorage) {
		expectedTag := testdata.GetTag(
			testdata.WithFilterRefIDs([]string{"javascript"}),
		)

		returnedTag, err := ts.GetTagByID(context.Background(), expectedTag.TagID)
		assert.NilError(t, err)

		assert.DeepEqual(t, expectedTag, returnedTag,
			opt.TimeWithThreshold(time.Hour))
	})
}
