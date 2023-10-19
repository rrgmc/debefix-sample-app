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

func testDBUserStorage(t *testing.T, testFn func(*sql.DB, *debefix.Data, storage.UserStorage),
	options ...dbtest.DBForTestOption) {
	db, resolvedData, dbCloseFunc, err := dbtest.DBForTest("debefix-sample-app", options...)
	assert.NilError(t, err)
	defer dbCloseFunc()

	ts := storage.NewUserStorage(db)

	testFn(db, resolvedData, ts)
}

func TestDBUserStorageGetUsers(t *testing.T) {
	testDBUserStorage(t, func(db *sql.DB, resolvedData *debefix.Data, ts storage.UserStorage) {
		filter := entity.UserFilter{
			Offset: 1,
			Limit:  1,
		}

		expectedUsers, err := testdata.GetUserList(
			testdata.WithFilterAll(true),
			testdata.WithSort("name"),
			testdata.WithOffsetLimit(filter.Offset, filter.Limit),
			testdata.WithResolvedData(resolvedData),
		)
		assert.NilError(t, err)
		assert.Assert(t, is.Len(expectedUsers, 1))

		returnedUsers, err := ts.GetUserList(context.Background(), filter)
		assert.NilError(t, err)

		assert.Assert(t, is.Len(returnedUsers, 1))
		assert.DeepEqual(t, expectedUsers, returnedUsers,
			opt.TimeWithThreshold(time.Hour))
	})
}

func TestDBUserStorageGetUserByID(t *testing.T) {
	testDBUserStorage(t, func(db *sql.DB, resolvedData *debefix.Data, ts storage.UserStorage) {
		expectedUser, err := testdata.GetUser(
			testdata.WithFilterRefIDs([]string{"test.DBUserStorageTestMergeData"}),
			testdata.WithResolvedData(resolvedData),
		)
		assert.NilError(t, err)

		returnedUser, err := ts.GetUserByID(context.Background(), expectedUser.UserID)
		assert.NilError(t, err)

		assert.DeepEqual(t, expectedUser, returnedUser,
			opt.TimeWithThreshold(time.Hour))
	}, dbtest.WithDBForTestMergeData(dbUserStorageTestMergeData()))
}

func TestDBUserStorageAddUser(t *testing.T) {
	testDBUserStorage(t, func(db *sql.DB, resolvedData *debefix.Data, ts storage.UserStorage) {
		findCountry, err := testdata.GetCountry(
			testdata.WithFilterRefIDs([]string{"usa"}),
			testdata.WithResolvedData(resolvedData),
		)
		assert.NilError(t, err)

		newUser := entity.User{
			Name:      "new user",
			Email:     "new email",
			CountryID: findCountry.CountryID,
		}

		returnedUser, err := ts.AddUser(context.Background(), newUser)
		assert.NilError(t, err)
		assert.Equal(t, "new user", returnedUser.Name)
	})
}

func TestDBUserStorageUpdateUserByID(t *testing.T) {
	testDBUserStorage(t, func(db *sql.DB, resolvedData *debefix.Data, ts storage.UserStorage) {
		findCountry, err := testdata.GetCountry(
			testdata.WithFilterRefIDs([]string{"usa"}),
			testdata.WithResolvedData(resolvedData),
		)
		assert.NilError(t, err)

		expectedUser, err := testdata.GetUser(
			testdata.WithFilterRefIDs([]string{"test.DBUserStorageTestMergeData"}),
			testdata.WithResolvedData(resolvedData),
		)
		assert.NilError(t, err)

		updatedUser := entity.User{
			Name:      "updated user",
			Email:     "updated email",
			CountryID: findCountry.CountryID,
		}

		returnedUser, err := ts.UpdateUserByID(context.Background(), expectedUser.UserID, updatedUser)
		assert.NilError(t, err)
		assert.Equal(t, "updated user", returnedUser.Name)
	}, dbtest.WithDBForTestMergeData(dbUserStorageTestMergeData()))
}

func TestDBUserStorageUpdateUserByIDNotFound(t *testing.T) {
	testDBUserStorage(t, func(db *sql.DB, resolvedData *debefix.Data, ts storage.UserStorage) {
		findCountry, err := testdata.GetCountry(
			testdata.WithFilterRefIDs([]string{"usa"}),
			testdata.WithResolvedData(resolvedData),
		)
		assert.NilError(t, err)

		updatedUser := entity.User{
			Name:      "updated user",
			Email:     "updated email",
			CountryID: findCountry.CountryID,
		}

		_, err = ts.UpdateUserByID(context.Background(), uuid.MustParse("0379ca21-7ed0-45e7-8812-4a6944f2c198"), updatedUser)
		assert.ErrorIs(t, err, util.ErrResourceNotFound)
	})
}

func TestDBUserStorageDeleteUserByID(t *testing.T) {
	testDBUserStorage(t, func(db *sql.DB, resolvedData *debefix.Data, ts storage.UserStorage) {
		expectedUser, err := testdata.GetUser(
			testdata.WithFilterRefIDs([]string{"test.DBUserStorageTestMergeData"}),
			testdata.WithResolvedData(resolvedData),
		)
		assert.NilError(t, err)

		err = ts.DeleteUserByID(context.Background(), expectedUser.UserID)
		assert.NilError(t, err)
	}, dbtest.WithDBForTestMergeData(dbUserStorageTestMergeData()))
}

func TestDBUserStorageDeleteUserByIDNotFound(t *testing.T) {
	testDBUserStorage(t, func(db *sql.DB, resolvedData *debefix.Data, ts storage.UserStorage) {
		err := ts.DeleteUserByID(context.Background(), uuid.MustParse("0379ca21-7ed0-45e7-8812-4a6944f2c198"))
		assert.ErrorIs(t, err, util.ErrResourceNotFound)
	})
}

func dbUserStorageTestMergeData() *debefix.Data {
	return &debefix.Data{
		Tables: map[string]*debefix.Table{
			"users": {
				Rows: debefix.Rows{
					{
						Config: debefix.RowConfig{
							RefID:      "test.DBUserStorageTestMergeData",
							IgnoreTags: true,
						},
						Fields: map[string]any{
							"user_id": &debefix.ValueGenerated{},
							"name":    "Test User",
							"email":   "Test Email",
							"country_id": &debefix.ValueRefID{
								TableID:   "countries",
								RefID:     "usa",
								FieldName: "country_id",
							},
							"created_at": time.Now(),
							"updated_at": time.Now(),
						},
					},
				},
			},
		},
	}
}
