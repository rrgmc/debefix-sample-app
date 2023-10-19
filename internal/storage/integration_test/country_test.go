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

func testDBCountryStorage(t *testing.T, testFn func(*sql.DB, *debefix.Data, storage.CountryStorage),
	options ...dbtest.DBForTestOption) {
	db, resolvedData, dbCloseFunc, err := dbtest.DBForTest("debefix-sample-app", options...)
	assert.NilError(t, err)
	defer dbCloseFunc()

	ts := storage.NewCountryStorage(db)

	testFn(db, resolvedData, ts)
}

func TestDBCountryStorageGetCountrys(t *testing.T) {
	testDBCountryStorage(t, func(db *sql.DB, resolvedData *debefix.Data, ts storage.CountryStorage) {
		filter := entity.CountryFilter{
			Offset: 1,
			Limit:  2,
		}

		expectedCountrys, err := testdata.GetCountryList(
			testdata.WithFilterAll(true),
			testdata.WithSort("name"),
			testdata.WithOffsetLimit(filter.Offset, filter.Limit),
		)
		assert.NilError(t, err)
		assert.Assert(t, is.Len(expectedCountrys, 2))

		returnedCountrys, err := ts.GetCountryList(context.Background(), filter)
		assert.NilError(t, err)

		assert.Assert(t, is.Len(returnedCountrys, 2))
		assert.DeepEqual(t, expectedCountrys, returnedCountrys,
			opt.TimeWithThreshold(time.Hour))
	})
}

func TestDBCountryStorageGetCountryByID(t *testing.T) {
	testDBCountryStorage(t, func(db *sql.DB, resolvedData *debefix.Data, ts storage.CountryStorage) {
		expectedCountry, err := testdata.GetCountry(
			testdata.WithFilterRefIDs([]string{"usa"}),
			testdata.WithResolvedData(resolvedData),
		)
		assert.NilError(t, err)

		returnedCountry, err := ts.GetCountryByID(context.Background(), expectedCountry.CountryID)
		assert.NilError(t, err)

		assert.DeepEqual(t, expectedCountry, returnedCountry,
			opt.TimeWithThreshold(time.Hour))
	})
}
