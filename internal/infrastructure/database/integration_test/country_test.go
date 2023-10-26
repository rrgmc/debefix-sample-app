//go:build dbtest

package integration_test

import (
	"context"
	"testing"
	"time"

	"github.com/rrgmc/debefix"
	"github.com/rrgmc/debefix-sample-app/internal/domain/entity"
	"github.com/rrgmc/debefix-sample-app/internal/domain/repository"
	"github.com/rrgmc/debefix-sample-app/internal/infrastructure/database"
	"github.com/rrgmc/debefix-sample-app/internal/infrastructure/database/testutils/dbtest"
	"github.com/rrgmc/debefix-sample-app/internal/infrastructure/database/testutils/testdata"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
	"gotest.tools/v3/assert/opt"
)

func testDBCountryRepository(t *testing.T, testFn func(repository.Context, *debefix.Data, repository.CountryRepository),
	options ...dbtest.DBForTestOption) {
	db, resolvedData, dbCloseFunc, err := dbtest.DBForTest("debefix-sample-app", options...)
	assert.NilError(t, err)
	defer dbCloseFunc()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info),
	})

	rctx := database.NewContext(gormDB)

	ts := database.NewCountryRepository()

	testFn(rctx, resolvedData, ts)
}

func TestDBCountryRepositoryGetCountrys(t *testing.T) {
	testDBCountryRepository(t, func(rctx repository.Context, resolvedData *debefix.Data, ts repository.CountryRepository) {
		filter := entity.CountryFilter{
			Offset: 1,
			Limit:  2,
		}

		expectedCountrys, err := testdata.GetCountryList(
			testdata.WithFilterRefIDs([]string{"uk", "usa"}),
			testdata.WithSort("refid"),
			testdata.WithResolvedData(resolvedData),
		)
		assert.NilError(t, err)
		assert.Assert(t, is.Len(expectedCountrys.Data, 2))

		returnedCountrys, err := ts.GetCountryList(context.Background(), rctx, filter)
		assert.NilError(t, err)

		assert.Assert(t, is.Len(returnedCountrys, 2))
		assert.DeepEqual(t, expectedCountrys.Data, returnedCountrys,
			opt.TimeWithThreshold(time.Hour))
	})
}

func TestDBCountryRepositoryGetCountryByID(t *testing.T) {
	testDBCountryRepository(t, func(rctx repository.Context, resolvedData *debefix.Data, ts repository.CountryRepository) {
		expectedCountry, err := testdata.GetCountry(
			testdata.WithFilterRefIDs([]string{"usa"}),
			testdata.WithResolvedData(resolvedData),
		)
		assert.NilError(t, err)

		returnedCountry, err := ts.GetCountryByID(context.Background(), rctx, expectedCountry.CountryID)
		assert.NilError(t, err)

		assert.DeepEqual(t, expectedCountry, returnedCountry,
			opt.TimeWithThreshold(time.Hour))
	})
}
