//go:build dbtest

package integration_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp/cmpopts"
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
	"gorm.io/gorm/logger"
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
	"gotest.tools/v3/assert/opt"
)

func testDBPostRepository(t *testing.T, testFn func(repository.Context, *debefix.Data, repository.PostRepository),
	options ...dbtest.DBForTestOption) {
	db, resolvedData, dbCloseFunc, err := dbtest.DBForTest("debefix-sample-app", options...)
	assert.NilError(t, err)
	defer dbCloseFunc()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	rctx := database.NewContext(gormDB)

	ts := database.NewPostRepository()

	testFn(rctx, resolvedData, ts)
}

func TestDBPostRepositoryGetPosts(t *testing.T) {
	testDBPostRepository(t, func(rctx repository.Context, resolvedData *debefix.Data, ts repository.PostRepository) {
		filter := entity.PostFilter{
			Offset: 1,
			Limit:  2,
		}

		expectedPosts, err := testdata.GetPostList(
			testdata.WithFilterRefIDs([]string{"post_2", "post_3"}),
			testdata.WithSort("refid"),
			testdata.WithResolvedData(resolvedData),
		)
		assert.NilError(t, err)
		assert.Assert(t, is.Len(expectedPosts.Data, 2))

		returnedPosts, err := ts.GetPostList(context.Background(), rctx, filter)
		assert.NilError(t, err)

		assert.Assert(t, is.Len(returnedPosts, 2))
		assert.DeepEqual(t, expectedPosts.Data, returnedPosts,
			opt.TimeWithThreshold(time.Hour))
	}, dbtest.WithDBForTestFixturesTags([]string{"tests.crud"}))
}

func TestDBPostRepositoryGetPostByID(t *testing.T) {
	testDBPostRepository(t, func(rctx repository.Context, resolvedData *debefix.Data, ts repository.PostRepository) {
		expectedPost, err := testdata.GetPost(
			testdata.WithFilterRefIDs([]string{"test.DBPostRepositoryTestMergeData"}),
			testdata.WithResolvedData(resolvedData),
		)
		assert.NilError(t, err)

		returnedPost, err := ts.GetPostByID(context.Background(), rctx, expectedPost.PostID)
		assert.NilError(t, err)

		assert.DeepEqual(t, expectedPost, returnedPost,
			opt.TimeWithThreshold(time.Hour))
	},
		dbtest.WithDBForTestFixturesTags([]string{"tests.crud"}),
		dbtest.WithDBForTestMergeData(dbPostRepositoryTestMergeData()))
}

func TestDBPostRepositoryAddPost(t *testing.T) {
	testDBPostRepository(t, func(rctx repository.Context, resolvedData *debefix.Data, ts repository.PostRepository) {
		findUser, err := testdata.GetUser(
			testdata.WithFilterRefIDs([]string{"janedoe"}),
			testdata.WithResolvedData(resolvedData),
		)
		assert.NilError(t, err)

		findTags, err := testdata.GetTagList(
			testdata.WithFilterRefIDs([]string{"go", "javascript"}),
			testdata.WithSort("refid"),
			testdata.WithResolvedData(resolvedData),
		)
		assert.NilError(t, err)

		expectedPost := entity.Post{
			Title:  "new title",
			Text:   "new text",
			UserID: findUser.UserID,
			Tags:   findTags.Data,
		}

		newPost := entity.PostAdd{
			Title:  expectedPost.Title,
			Text:   expectedPost.Text,
			UserID: expectedPost.UserID,
		}
		for _, tag := range expectedPost.Tags {
			newPost.Tags = append(newPost.Tags, tag.TagID)
		}

		returnedPost, err := ts.AddPost(context.Background(), rctx, newPost)
		assert.NilError(t, err)
		assert.DeepEqual(t, expectedPost, returnedPost,
			cmpopts.IgnoreFields(entity.Post{}, "PostID"),
			cmpopts.IgnoreTypes(time.Time{}))
	}, dbtest.WithDBForTestFixturesTags([]string{"tests.crud"}))
}

func TestDBPostRepositoryUpdatePostByID(t *testing.T) {
	testDBPostRepository(t, func(rctx repository.Context, resolvedData *debefix.Data, ts repository.PostRepository) {
		findUser, err := testdata.GetUser(
			testdata.WithFilterRefIDs([]string{"johndoe"}),
			testdata.WithResolvedData(resolvedData),
		)
		assert.NilError(t, err)

		findTags, err := testdata.GetTagList(
			testdata.WithFilterRefIDs([]string{"go"}),
			testdata.WithSort("refid"),
			testdata.WithResolvedData(resolvedData),
		)
		assert.NilError(t, err)

		currentPost, err := testdata.GetPost(
			testdata.WithFilterRefIDs([]string{"test.DBPostRepositoryTestMergeData"}),
			testdata.WithResolvedData(resolvedData),
		)
		assert.NilError(t, err)

		expectedPost := entity.Post{
			Title:  "updated title",
			Text:   "updated text",
			UserID: findUser.UserID,
			Tags:   findTags.Data,
		}

		updatedPost := entity.PostUpdate{
			Title:  expectedPost.Title,
			Text:   expectedPost.Text,
			UserID: expectedPost.UserID,
		}
		for _, tag := range expectedPost.Tags {
			updatedPost.Tags = append(updatedPost.Tags, tag.TagID)
		}

		returnedPost, err := ts.UpdatePostByID(context.Background(), rctx, currentPost.PostID, updatedPost)
		assert.NilError(t, err)
		assert.DeepEqual(t, expectedPost, returnedPost,
			cmpopts.IgnoreFields(entity.Post{}, "PostID"),
			cmpopts.IgnoreTypes(time.Time{}))
	},
		dbtest.WithDBForTestFixturesTags([]string{"tests.crud"}),
		dbtest.WithDBForTestMergeData(dbPostRepositoryTestMergeData()))
}

func TestDBPostRepositoryUpdatePostByIDNotFound(t *testing.T) {
	testDBPostRepository(t, func(rctx repository.Context, resolvedData *debefix.Data, ts repository.PostRepository) {
		findUser, err := testdata.GetUser(
			testdata.WithFilterRefIDs([]string{"janedoe"}),
			testdata.WithResolvedData(resolvedData),
		)
		assert.NilError(t, err)

		updatedPost := entity.PostUpdate{
			Title:  "updated title",
			Text:   "updated text",
			UserID: findUser.UserID,
		}

		_, err = ts.UpdatePostByID(context.Background(), rctx, uuid.MustParse("0379ca21-7ed0-45e7-8812-4a6944f2c198"), updatedPost)
		assert.ErrorIs(t, err, domain.NotFound)
	}, dbtest.WithDBForTestFixturesTags([]string{"tests.crud"}))
}

func TestDBPostRepositoryDeletePostByID(t *testing.T) {
	testDBPostRepository(t, func(rctx repository.Context, resolvedData *debefix.Data, ts repository.PostRepository) {
		expectedPost, err := testdata.GetPost(
			testdata.WithFilterRefIDs([]string{"test.DBPostRepositoryTestMergeData"}),
			testdata.WithResolvedData(resolvedData),
		)
		assert.NilError(t, err)

		err = ts.DeletePostByID(context.Background(), rctx, expectedPost.PostID)
		assert.NilError(t, err)
	},
		dbtest.WithDBForTestFixturesTags([]string{"tests.crud"}),
		dbtest.WithDBForTestMergeData(dbPostRepositoryTestMergeData()))
}

func TestDBPostRepositoryDeletePostByIDNotFound(t *testing.T) {
	testDBPostRepository(t, func(rctx repository.Context, resolvedData *debefix.Data, ts repository.PostRepository) {
		err := ts.DeletePostByID(context.Background(), rctx, uuid.MustParse("0379ca21-7ed0-45e7-8812-4a6944f2c198"))
		assert.ErrorIs(t, err, domain.NotFound)
	}, dbtest.WithDBForTestFixturesTags([]string{"tests.crud"}))
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
      config:
        !dbfconfig
        refid: "test.DBPostRepositoryTestMergeData"
      deps:
        !dbfdeps
        posts_tags:
          rows:
            - post_id: !dbfexpr "parent:post_id"
              tag_id: !dbfexpr "refid:tags:javascript:tag_id"
`}
}
