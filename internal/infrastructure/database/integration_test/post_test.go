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
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
	"gotest.tools/v3/assert/opt"
)

func testDBPostRepository(t *testing.T, testFn func(*sql.DB, *debefix.Data, repository.PostRepository),
	options ...dbtest.DBForTestOption) {
	db, resolvedData, dbCloseFunc, err := dbtest.DBForTest("debefix-sample-app", options...)
	assert.NilError(t, err)
	defer dbCloseFunc()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info),
	})

	ts := database.NewPostRepository(gormDB)

	testFn(db, resolvedData, ts)
}

func TestDBPostRepositoryGetPosts(t *testing.T) {
	testDBPostRepository(t, func(db *sql.DB, resolvedData *debefix.Data, ts repository.PostRepository) {
		filter := model.PostFilter{
			Offset: 1,
			Limit:  1,
		}

		expectedPosts, err := testdata.GetPostList(
			testdata.WithFilterAll(true),
			testdata.WithSort("title"),
			testdata.WithOffsetLimit(filter.Offset, filter.Limit),
			testdata.WithResolvedData(resolvedData),
		)
		assert.NilError(t, err)
		assert.Assert(t, is.Len(expectedPosts, 1))

		returnedPosts, err := ts.GetPostList(context.Background(), filter)
		assert.NilError(t, err)

		assert.Assert(t, is.Len(returnedPosts, 1))
		assert.DeepEqual(t, expectedPosts, returnedPosts,
			opt.TimeWithThreshold(time.Hour))
	})
}

func TestDBPostRepositoryGetPostByID(t *testing.T) {
	testDBPostRepository(t, func(db *sql.DB, resolvedData *debefix.Data, ts repository.PostRepository) {
		expectedPost, err := testdata.GetPost(
			testdata.WithFilterRefIDs([]string{"test.DBPostRepositoryTestMergeData"}),
			testdata.WithResolvedData(resolvedData),
		)
		assert.NilError(t, err)

		returnedPost, err := ts.GetPostByID(context.Background(), expectedPost.PostID)
		assert.NilError(t, err)

		assert.DeepEqual(t, expectedPost, returnedPost,
			opt.TimeWithThreshold(time.Hour))
	}, dbtest.WithDBForTestMergeData(dbPostRepositoryTestMergeData()))
}

func TestDBPostRepositoryAddPost(t *testing.T) {
	testDBPostRepository(t, func(db *sql.DB, resolvedData *debefix.Data, ts repository.PostRepository) {
		findUser, err := testdata.GetUser(
			testdata.WithFilterRefIDs([]string{"janedoe"}),
			testdata.WithResolvedData(resolvedData),
		)
		assert.NilError(t, err)

		findTags, err := testdata.GetTagList(
			testdata.WithFilterRefIDs([]string{"javascript", "go"}),
			testdata.WithResolvedData(resolvedData),
		)
		assert.NilError(t, err)

		newPost := model.Post{
			Title:  "new title",
			Text:   "new text",
			UserID: findUser.UserID,
			Tags:   findTags,
		}

		returnedPost, err := ts.AddPost(context.Background(), newPost)
		assert.NilError(t, err)
		assert.Equal(t, "new title", returnedPost.Title)
		assert.Equal(t, "new text", returnedPost.Text)
	})
}

func TestDBPostRepositoryUpdatePostByID(t *testing.T) {
	testDBPostRepository(t, func(db *sql.DB, resolvedData *debefix.Data, ts repository.PostRepository) {
		findUser, err := testdata.GetUser(
			testdata.WithFilterRefIDs([]string{"janedoe"}),
			testdata.WithResolvedData(resolvedData),
		)
		assert.NilError(t, err)

		expectedPost, err := testdata.GetPost(
			testdata.WithFilterRefIDs([]string{"test.DBPostRepositoryTestMergeData"}),
			testdata.WithResolvedData(resolvedData),
		)
		assert.NilError(t, err)

		updatedPost := model.Post{
			Title:  "updated title",
			Text:   "updated text",
			UserID: findUser.UserID,
		}

		returnedPost, err := ts.UpdatePostByID(context.Background(), expectedPost.PostID, updatedPost)
		assert.NilError(t, err)
		assert.Equal(t, "updated title", returnedPost.Title)
		assert.Equal(t, "updated text", returnedPost.Text)
	}, dbtest.WithDBForTestMergeData(dbPostRepositoryTestMergeData()))
}

func TestDBPostRepositoryUpdatePostByIDNotFound(t *testing.T) {
	testDBPostRepository(t, func(db *sql.DB, resolvedData *debefix.Data, ts repository.PostRepository) {
		findUser, err := testdata.GetUser(
			testdata.WithFilterRefIDs([]string{"janedoe"}),
			testdata.WithResolvedData(resolvedData),
		)
		assert.NilError(t, err)

		updatedPost := model.Post{
			Title:  "updated title",
			Text:   "updated text",
			UserID: findUser.UserID,
		}

		_, err = ts.UpdatePostByID(context.Background(), uuid.MustParse("0379ca21-7ed0-45e7-8812-4a6944f2c198"), updatedPost)
		assert.ErrorIs(t, err, utils.ErrResourceNotFound)
	})
}

func TestDBPostRepositoryDeletePostByID(t *testing.T) {
	testDBPostRepository(t, func(db *sql.DB, resolvedData *debefix.Data, ts repository.PostRepository) {
		expectedPost, err := testdata.GetPost(
			testdata.WithFilterRefIDs([]string{"test.DBPostRepositoryTestMergeData"}),
			testdata.WithResolvedData(resolvedData),
		)
		assert.NilError(t, err)

		err = ts.DeletePostByID(context.Background(), expectedPost.PostID)
		assert.NilError(t, err)
	}, dbtest.WithDBForTestMergeData(dbPostRepositoryTestMergeData()))
}

func TestDBPostRepositoryDeletePostByIDNotFound(t *testing.T) {
	testDBPostRepository(t, func(db *sql.DB, resolvedData *debefix.Data, ts repository.PostRepository) {
		err := ts.DeletePostByID(context.Background(), uuid.MustParse("0379ca21-7ed0-45e7-8812-4a6944f2c198"))
		assert.ErrorIs(t, err, utils.ErrResourceNotFound)
	})
}

func dbPostRepositoryTestMergeData() []string {
	return []string{`
posts:
  rows:
    - post_id: !dbfexpr generated:uuid
      title: "Test Title"
      text: "Test Text"
      user_id: !dbfexpr "refid:users:janedoe:user_id"
      created_at: !!timestamp 2023-03-01T12:30:12Z
      updated_at: !!timestamp 2023-03-01T12:30:12Z
      _dbfconfig:
        refid: "test.DBPostRepositoryTestMergeData"
      _dbfdeps:
        posts_tags:
          rows:
            - post_id: !dbfexpr "parent:post_id"
              tag_id: !dbfexpr "refid:tags:javascript:tag_id"
`}
}
