//go:build dbtest

package integration_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/rrgmc/debefix"
	"github.com/rrgmc/debefix-sample-app/internal/domain"
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

func testDBTagRepository(t *testing.T, testFn func(repository.Context, *debefix.Data, repository.TagRepository),
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

	ts := database.NewTagRepository()

	testFn(rctx, resolvedData, ts)
}

func TestDBTagRepositoryGetTagList(t *testing.T) {
	testDBTagRepository(t, func(rctx repository.Context, resolvedData *debefix.Data, ts repository.TagRepository) {
		filter := entity.TagFilter{
			Offset: 1,
			Limit:  2,
		}

		expectedTags, err := testdata.GetTagList(
			testdata.WithFilterRefIDs([]string{"go", "javascript"}),
			testdata.WithSort("refid"),
			testdata.WithResolvedData(resolvedData),
		)
		assert.NilError(t, err)
		assert.Assert(t, is.Len(expectedTags.Data, 2))

		returnedTags, err := ts.GetTagList(context.Background(), rctx, filter)
		assert.NilError(t, err)

		assert.Assert(t, is.Len(returnedTags, 2))
		assert.DeepEqual(t, expectedTags.Data, returnedTags,
			opt.TimeWithThreshold(time.Hour))
	}, dbtest.WithDBForTestFixturesTags([]string{"tests.crud"}))
}

func TestDBTagRepositoryGetTagByID(t *testing.T) {
	testDBTagRepository(t, func(rctx repository.Context, resolvedData *debefix.Data, ts repository.TagRepository) {
		expectedTag, err := testdata.GetTag(
			testdata.WithFilterRefIDs([]string{"test.DBTagRepositoryTestMergeData"}),
			testdata.WithResolvedData(resolvedData),
		)
		assert.NilError(t, err)

		returnedTag, err := ts.GetTagByID(context.Background(), rctx, expectedTag.TagID)
		assert.NilError(t, err)

		assert.DeepEqual(t, expectedTag, returnedTag,
			opt.TimeWithThreshold(time.Hour))
	},
		dbtest.WithDBForTestFixturesTags([]string{"tests.crud"}),
		dbtest.WithDBForTestMergeData(dbTagRepositoryTestMergeData()))
}

func TestDBTagRepositoryAddTag(t *testing.T) {
	testDBTagRepository(t, func(rctx repository.Context, resolvedData *debefix.Data, ts repository.TagRepository) {
		newTag := entity.TagAdd{
			Name: "new tag",
		}

		returnedTag, err := ts.AddTag(context.Background(), rctx, newTag)
		assert.NilError(t, err)
		assert.Equal(t, "new tag", returnedTag.Name)
	}, dbtest.WithDBForTestFixturesTags([]string{"tests.crud"}))
}

func TestDBTagRepositoryUpdateTagByID(t *testing.T) {
	testDBTagRepository(t, func(rctx repository.Context, resolvedData *debefix.Data, ts repository.TagRepository) {
		expectedTag, err := testdata.GetTag(
			testdata.WithFilterRefIDs([]string{"test.DBTagRepositoryTestMergeData"}),
			testdata.WithResolvedData(resolvedData),
		)
		assert.NilError(t, err)

		updatedTag := entity.TagAdd{
			Name: "updated tag",
		}

		returnedTag, err := ts.UpdateTagByID(context.Background(), rctx, expectedTag.TagID, updatedTag)
		assert.NilError(t, err)
		assert.Equal(t, "updated tag", returnedTag.Name)
	},
		dbtest.WithDBForTestFixturesTags([]string{"tests.crud"}),
		dbtest.WithDBForTestMergeData(dbTagRepositoryTestMergeData()))
}

func TestDBTagRepositoryUpdateTagByIDNotFound(t *testing.T) {
	testDBTagRepository(t, func(rctx repository.Context, resolvedData *debefix.Data, ts repository.TagRepository) {
		updatedTag := entity.TagUpdate{
			Name: "updated tag",
		}

		_, err := ts.UpdateTagByID(context.Background(), rctx, uuid.MustParse("0379ca21-7ed0-45e7-8812-4a6944f2c198"), updatedTag)
		assert.ErrorIs(t, err, domain.NotFound)
	}, dbtest.WithDBForTestFixturesTags([]string{"tests.crud"}))
}

func TestDBTagRepositoryDeleteTagByID(t *testing.T) {
	testDBTagRepository(t, func(rctx repository.Context, resolvedData *debefix.Data, ts repository.TagRepository) {
		expectedTag, err := testdata.GetTag(
			testdata.WithFilterRefIDs([]string{"test.DBTagRepositoryTestMergeData"}),
			testdata.WithResolvedData(resolvedData),
		)
		assert.NilError(t, err)

		err = ts.DeleteTagByID(context.Background(), rctx, expectedTag.TagID)
		assert.NilError(t, err)
	},
		dbtest.WithDBForTestFixturesTags([]string{"tests.crud"}),
		dbtest.WithDBForTestMergeData(dbTagRepositoryTestMergeData()))
}

func TestDBTagRepositoryDeleteTagByIDNotFound(t *testing.T) {
	testDBTagRepository(t, func(rctx repository.Context, resolvedData *debefix.Data, ts repository.TagRepository) {
		err := ts.DeleteTagByID(context.Background(), rctx, uuid.MustParse("0379ca21-7ed0-45e7-8812-4a6944f2c198"))
		assert.ErrorIs(t, err, domain.NotFound)
	}, dbtest.WithDBForTestFixturesTags([]string{"tests.crud"}))
}

func dbTagRepositoryTestMergeData() []string {
	return []string{`
tags:
  rows:
    - tag_id: !dbfexpr generated:uuid
      name: "Test Tag"
      created_at: !!timestamp 2023-03-01T12:30:12Z
      updated_at: !!timestamp 2023-03-01T12:30:12Z
      _dbfconfig:
        refid: "test.DBTagRepositoryTestMergeData"
`}
}
