//go:build dbtest

package integration_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/rrgmc/debefix"
	"github.com/rrgmc/debefix-sample-app/internal/domain/model"
	"github.com/rrgmc/debefix-sample-app/internal/domain/repository"
	"github.com/rrgmc/debefix-sample-app/internal/infrastructure/database/testutils/dbtest"
	"github.com/rrgmc/debefix-sample-app/internal/infrastructure/database/testutils/testdata"
	"github.com/rrgmc/debefix-sample-app/internal/infrastructure/databasegorm"
	"github.com/rrgmc/debefix-sample-app/internal/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
	"gotest.tools/v3/assert/opt"
)

func testDBTagRepository(t *testing.T, testFn func(*gorm.DB, *debefix.Data, repository.TagRepository),
	options ...dbtest.DBForTestOption) {
	db, resolvedData, dbCloseFunc, err := dbtest.DBForTest("debefix-sample-app", options...)
	assert.NilError(t, err)
	defer dbCloseFunc()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})

	ts := databasegorm.NewTagRepository(gormDB)

	testFn(gormDB, resolvedData, ts)
}

func TestDBTagRepositoryGetTags(t *testing.T) {
	testDBTagRepository(t, func(db *gorm.DB, resolvedData *debefix.Data, ts repository.TagRepository) {
		filter := model.TagFilter{
			Offset: 1,
			Limit:  2,
		}

		expectedTags, err := testdata.GetTagList(
			testdata.WithFilterAll(true),
			testdata.WithSort("name"),
			testdata.WithOffsetLimit(filter.Offset, filter.Limit),
			testdata.WithResolvedData(resolvedData),
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

func TestDBTagRepositoryGetTagByID(t *testing.T) {
	testDBTagRepository(t, func(db *gorm.DB, resolvedData *debefix.Data, ts repository.TagRepository) {
		expectedTag, err := testdata.GetTag(
			testdata.WithFilterRefIDs([]string{"test.DBTagRepositoryTestMergeData"}),
			testdata.WithResolvedData(resolvedData),
		)
		assert.NilError(t, err)

		returnedTag, err := ts.GetTagByID(context.Background(), expectedTag.TagID)
		assert.NilError(t, err)

		assert.DeepEqual(t, expectedTag, returnedTag,
			opt.TimeWithThreshold(time.Hour))
	}, dbtest.WithDBForTestMergeData(dbTagRepositoryTestMergeData()))
}

func TestDBTagRepositoryAddTag(t *testing.T) {
	testDBTagRepository(t, func(db *gorm.DB, resolvedData *debefix.Data, ts repository.TagRepository) {
		newTag := model.Tag{
			Name: "new tag",
		}

		returnedTag, err := ts.AddTag(context.Background(), newTag)
		assert.NilError(t, err)
		assert.Equal(t, "new tag", returnedTag.Name)
	})
}

func TestDBTagRepositoryUpdateTagByID(t *testing.T) {
	testDBTagRepository(t, func(db *gorm.DB, resolvedData *debefix.Data, ts repository.TagRepository) {
		expectedTag, err := testdata.GetTag(
			testdata.WithFilterRefIDs([]string{"test.DBTagRepositoryTestMergeData"}),
			testdata.WithResolvedData(resolvedData),
		)
		assert.NilError(t, err)

		updatedTag := model.Tag{
			Name: "updated tag",
		}

		returnedTag, err := ts.UpdateTagByID(context.Background(), expectedTag.TagID, updatedTag)
		assert.NilError(t, err)
		assert.Equal(t, "updated tag", returnedTag.Name)
	}, dbtest.WithDBForTestMergeData(dbTagRepositoryTestMergeData()))
}

func TestDBTagRepositoryUpdateTagByIDNotFound(t *testing.T) {
	testDBTagRepository(t, func(db *gorm.DB, resolvedData *debefix.Data, ts repository.TagRepository) {
		updatedTag := model.Tag{
			Name: "updated tag",
		}

		_, err := ts.UpdateTagByID(context.Background(), uuid.MustParse("0379ca21-7ed0-45e7-8812-4a6944f2c198"), updatedTag)
		assert.ErrorIs(t, err, utils.ErrResourceNotFound)
	})
}

func TestDBTagRepositoryDeleteTagByID(t *testing.T) {
	testDBTagRepository(t, func(db *gorm.DB, resolvedData *debefix.Data, ts repository.TagRepository) {
		expectedTag, err := testdata.GetTag(
			testdata.WithFilterRefIDs([]string{"test.DBTagRepositoryTestMergeData"}),
			testdata.WithResolvedData(resolvedData),
		)
		assert.NilError(t, err)

		err = ts.DeleteTagByID(context.Background(), expectedTag.TagID)
		assert.NilError(t, err)
	}, dbtest.WithDBForTestMergeData(dbTagRepositoryTestMergeData()))
}

func TestDBTagRepositoryDeleteTagByIDNotFound(t *testing.T) {
	testDBTagRepository(t, func(db *gorm.DB, resolvedData *debefix.Data, ts repository.TagRepository) {
		err := ts.DeleteTagByID(context.Background(), uuid.MustParse("0379ca21-7ed0-45e7-8812-4a6944f2c198"))
		assert.ErrorIs(t, err, utils.ErrResourceNotFound)
	})
}

func dbTagRepositoryTestMergeData() *debefix.Data {
	return &debefix.Data{
		Tables: map[string]*debefix.Table{
			"tags": {
				Rows: debefix.Rows{
					{
						Config: debefix.RowConfig{
							RefID:      "test.DBTagRepositoryTestMergeData",
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
