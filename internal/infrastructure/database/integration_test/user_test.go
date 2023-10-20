//go:build dbtest

package integration_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/rrgmc/debefix"
	"github.com/rrgmc/debefix-sample-app/internal/domain/model"
	"github.com/rrgmc/debefix-sample-app/internal/domain/repository"
	"github.com/rrgmc/debefix-sample-app/internal/infrastructure/database"
	"github.com/rrgmc/debefix-sample-app/internal/infrastructure/database/testutils/dbtest"
	"github.com/rrgmc/debefix-sample-app/internal/infrastructure/database/testutils/testdata"
	"github.com/rrgmc/debefix-sample-app/internal/utils"
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
	"gotest.tools/v3/assert/opt"
)

func testDBUserRepository(t *testing.T, testFn func(*sql.DB, *debefix.Data, repository.UserRepository),
	options ...dbtest.DBForTestOption) {
	db, resolvedData, dbCloseFunc, err := dbtest.DBForTest("debefix-sample-app", options...)
	assert.NilError(t, err)
	defer dbCloseFunc()

	ts := database.NewUserRepository(db)

	testFn(db, resolvedData, ts)
}

func TestDBUserRepositoryGetUsers(t *testing.T) {
	testDBUserRepository(t, func(db *sql.DB, resolvedData *debefix.Data, ts repository.UserRepository) {
		filter := model.UserFilter{
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

func TestDBUserRepositoryGetUserByID(t *testing.T) {
	testDBUserRepository(t, func(db *sql.DB, resolvedData *debefix.Data, ts repository.UserRepository) {
		expectedUser, err := testdata.GetUser(
			testdata.WithFilterRefIDs([]string{"test.DBUserRepositoryTestMergeData"}),
			testdata.WithResolvedData(resolvedData),
		)
		assert.NilError(t, err)

		returnedUser, err := ts.GetUserByID(context.Background(), expectedUser.UserID)
		assert.NilError(t, err)

		assert.DeepEqual(t, expectedUser, returnedUser,
			opt.TimeWithThreshold(time.Hour))
	}, dbtest.WithDBForTestMergeData(dbUserRepositoryTestMergeData()))
}

func TestDBUserRepositoryAddUser(t *testing.T) {
	testDBUserRepository(t, func(db *sql.DB, resolvedData *debefix.Data, ts repository.UserRepository) {
		findCountry, err := testdata.GetCountry(
			testdata.WithFilterRefIDs([]string{"usa"}),
			testdata.WithResolvedData(resolvedData),
		)
		assert.NilError(t, err)

		newUser := model.User{
			Name:      "new user",
			Email:     "new email",
			CountryID: findCountry.CountryID,
		}

		returnedUser, err := ts.AddUser(context.Background(), newUser)
		assert.NilError(t, err)
		assert.Equal(t, "new user", returnedUser.Name)
	})
}

func TestDBUserRepositoryUpdateUserByID(t *testing.T) {
	testDBUserRepository(t, func(db *sql.DB, resolvedData *debefix.Data, ts repository.UserRepository) {
		findCountry, err := testdata.GetCountry(
			testdata.WithFilterRefIDs([]string{"usa"}),
			testdata.WithResolvedData(resolvedData),
		)
		assert.NilError(t, err)

		expectedUser, err := testdata.GetUser(
			testdata.WithFilterRefIDs([]string{"test.DBUserRepositoryTestMergeData"}),
			testdata.WithResolvedData(resolvedData),
		)
		assert.NilError(t, err)

		updatedUser := model.User{
			Name:      "updated user",
			Email:     "updated email",
			CountryID: findCountry.CountryID,
		}

		returnedUser, err := ts.UpdateUserByID(context.Background(), expectedUser.UserID, updatedUser)
		assert.NilError(t, err)
		assert.Equal(t, "updated user", returnedUser.Name)
	}, dbtest.WithDBForTestMergeData(dbUserRepositoryTestMergeData()))
}

func TestDBUserRepositoryUpdateUserByIDNotFound(t *testing.T) {
	testDBUserRepository(t, func(db *sql.DB, resolvedData *debefix.Data, ts repository.UserRepository) {
		findCountry, err := testdata.GetCountry(
			testdata.WithFilterRefIDs([]string{"usa"}),
			testdata.WithResolvedData(resolvedData),
		)
		assert.NilError(t, err)

		updatedUser := model.User{
			Name:      "updated user",
			Email:     "updated email",
			CountryID: findCountry.CountryID,
		}

		_, err = ts.UpdateUserByID(context.Background(), uuid.MustParse("0379ca21-7ed0-45e7-8812-4a6944f2c198"), updatedUser)
		assert.ErrorIs(t, err, utils.ErrResourceNotFound)
	})
}

func TestDBUserRepositoryDeleteUserByID(t *testing.T) {
	testDBUserRepository(t, func(db *sql.DB, resolvedData *debefix.Data, ts repository.UserRepository) {
		expectedUser, err := testdata.GetUser(
			testdata.WithFilterRefIDs([]string{"test.DBUserRepositoryTestMergeData"}),
			testdata.WithResolvedData(resolvedData),
		)
		assert.NilError(t, err)

		err = ts.DeleteUserByID(context.Background(), expectedUser.UserID)
		assert.NilError(t, err)
	}, dbtest.WithDBForTestMergeData(dbUserRepositoryTestMergeData()))
}

func TestDBUserRepositoryDeleteUserByIDNotFound(t *testing.T) {
	testDBUserRepository(t, func(db *sql.DB, resolvedData *debefix.Data, ts repository.UserRepository) {
		err := ts.DeleteUserByID(context.Background(), uuid.MustParse("0379ca21-7ed0-45e7-8812-4a6944f2c198"))
		assert.ErrorIs(t, err, utils.ErrResourceNotFound)
	})
}

func dbUserRepositoryTestMergeData() *debefix.Data {
	return &debefix.Data{
		Tables: map[string]*debefix.Table{
			"users": {
				Rows: debefix.Rows{
					{
						Config: debefix.RowConfig{
							RefID:      "test.DBUserRepositoryTestMergeData",
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
