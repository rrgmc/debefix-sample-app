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

func testDBPostStorage(t *testing.T, testFn func(*sql.DB, *debefix.Data, storage.PostStorage),
	options ...dbtest.DBForTestOption) {
	db, resolvedData, dbCloseFunc, err := dbtest.DBForTest("debefix-sample-app", options...)
	assert.NilError(t, err)
	defer dbCloseFunc()

	ts := storage.NewPostStorage(db)

	testFn(db, resolvedData, ts)
}

func TestDBPostStorageGetPosts(t *testing.T) {
	testDBPostStorage(t, func(db *sql.DB, resolvedData *debefix.Data, ts storage.PostStorage) {
		filter := entity.PostFilter{
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

func TestDBPostStorageGetPostByID(t *testing.T) {
	testDBPostStorage(t, func(db *sql.DB, resolvedData *debefix.Data, ts storage.PostStorage) {
		expectedPost, err := testdata.GetPost(
			testdata.WithFilterRefIDs([]string{"test.DBPostStorageTestMergeData"}),
			testdata.WithResolvedData(resolvedData),
		)
		assert.NilError(t, err)

		returnedPost, err := ts.GetPostByID(context.Background(), expectedPost.PostID)
		assert.NilError(t, err)

		assert.DeepEqual(t, expectedPost, returnedPost,
			opt.TimeWithThreshold(time.Hour))
	}, dbtest.WithDBForTestMergeData(dbPostStorageTestMergeData()))
}

func TestDBPostStorageAddPost(t *testing.T) {
	testDBPostStorage(t, func(db *sql.DB, resolvedData *debefix.Data, ts storage.PostStorage) {
		findUser, err := testdata.GetUser(
			testdata.WithFilterRefIDs([]string{"janedoe"}),
			testdata.WithResolvedData(resolvedData),
		)
		assert.NilError(t, err)

		newPost := entity.Post{
			Title:  "new title",
			Text:   "new text",
			UserID: findUser.UserID,
		}

		returnedPost, err := ts.AddPost(context.Background(), newPost)
		assert.NilError(t, err)
		assert.Equal(t, "new title", returnedPost.Title)
		assert.Equal(t, "new text", returnedPost.Text)
	})
}

func TestDBPostStorageUpdatePostByID(t *testing.T) {
	testDBPostStorage(t, func(db *sql.DB, resolvedData *debefix.Data, ts storage.PostStorage) {
		findUser, err := testdata.GetUser(
			testdata.WithFilterRefIDs([]string{"janedoe"}),
			testdata.WithResolvedData(resolvedData),
		)
		assert.NilError(t, err)

		expectedPost, err := testdata.GetPost(
			testdata.WithFilterRefIDs([]string{"test.DBPostStorageTestMergeData"}),
			testdata.WithResolvedData(resolvedData),
		)
		assert.NilError(t, err)

		updatedPost := entity.Post{
			Title:  "updated title",
			Text:   "updated text",
			UserID: findUser.UserID,
		}

		returnedPost, err := ts.UpdatePostByID(context.Background(), expectedPost.PostID, updatedPost)
		assert.NilError(t, err)
		assert.Equal(t, "updated title", returnedPost.Title)
		assert.Equal(t, "updated text", returnedPost.Text)
	}, dbtest.WithDBForTestMergeData(dbPostStorageTestMergeData()))
}

func TestDBPostStorageUpdatePostByIDNotFound(t *testing.T) {
	testDBPostStorage(t, func(db *sql.DB, resolvedData *debefix.Data, ts storage.PostStorage) {
		findUser, err := testdata.GetUser(
			testdata.WithFilterRefIDs([]string{"janedoe"}),
			testdata.WithResolvedData(resolvedData),
		)
		assert.NilError(t, err)

		updatedPost := entity.Post{
			Title:  "updated title",
			Text:   "updated text",
			UserID: findUser.UserID,
		}

		_, err = ts.UpdatePostByID(context.Background(), uuid.MustParse("0379ca21-7ed0-45e7-8812-4a6944f2c198"), updatedPost)
		assert.ErrorIs(t, err, util.ErrResourceNotFound)
	})
}

func TestDBPostStorageDeletePostByID(t *testing.T) {
	testDBPostStorage(t, func(db *sql.DB, resolvedData *debefix.Data, ts storage.PostStorage) {
		expectedPost, err := testdata.GetPost(
			testdata.WithFilterRefIDs([]string{"test.DBPostStorageTestMergeData"}),
			testdata.WithResolvedData(resolvedData),
		)
		assert.NilError(t, err)

		err = ts.DeletePostByID(context.Background(), expectedPost.PostID)
		assert.NilError(t, err)
	}, dbtest.WithDBForTestMergeData(dbPostStorageTestMergeData()))
}

func TestDBPostStorageDeletePostByIDNotFound(t *testing.T) {
	testDBPostStorage(t, func(db *sql.DB, resolvedData *debefix.Data, ts storage.PostStorage) {
		err := ts.DeletePostByID(context.Background(), uuid.MustParse("0379ca21-7ed0-45e7-8812-4a6944f2c198"))
		assert.ErrorIs(t, err, util.ErrResourceNotFound)
	})
}

func dbPostStorageTestMergeData() *debefix.Data {
	return &debefix.Data{
		Tables: map[string]*debefix.Table{
			"posts": {
				Rows: debefix.Rows{
					{
						Config: debefix.RowConfig{
							RefID:      "test.DBPostStorageTestMergeData",
							IgnoreTags: true,
						},
						Fields: map[string]any{
							"post_id": &debefix.ValueGenerated{},
							"title":   "Test Title",
							"text":    "Test Text",
							"user_id": &debefix.ValueRefID{
								TableID:   "users",
								RefID:     "janedoe",
								FieldName: "user_id",
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
