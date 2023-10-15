//go:build dbtest

package integration_test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/RangelReale/debefix-sample-app/internal/entity"
	"github.com/RangelReale/debefix-sample-app/internal/storage"
	"github.com/RangelReale/debefix-sample-app/internal/testutils"
	"github.com/RangelReale/debefix-sample-app/internal/testutils/dbtest"
	"github.com/RangelReale/debefix-sample-app/internal/testutils/testdata"
	"github.com/stretchr/testify/require"
)

func testDBTagStorage(t *testing.T, testFn func(*sql.DB, storage.TagStorage)) {
	db, dbCloseFunc, err := dbtest.DBForTest("olympus")
	require.NoError(t, err)
	defer dbCloseFunc()

	ts := storage.NewTagStorage(db)

	testFn(db, ts)
}

func TestDBTagStorageGetTags(t *testing.T) {
	testDBTagStorage(t, func(db *sql.DB, ts storage.TagStorage) {
		expectedTags := testdata.GetTags([]string{"go", "javascript", "cpp"},
			testdata.WithSort("name"))

		returnedTags, err := ts.GetTags(context.Background(), entity.TagsFilter{
			Offset: 0,
			Limit:  100,
		})
		require.NoError(t, err)

		require.Len(t, returnedTags, 3)
		testutils.Equal(t, expectedTags, returnedTags)
	})
}
