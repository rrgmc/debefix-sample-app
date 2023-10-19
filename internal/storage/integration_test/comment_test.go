//go:build dbtest

package integration_test

import (
	"context"
	"database/sql"
	"fmt"
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

func testDBCommentStorage(t *testing.T, testFn func(*sql.DB, *debefix.Data, storage.CommentStorage),
	options ...dbtest.DBForTestOption) {
	db, resolvedData, dbCloseFunc, err := dbtest.DBForTest("debefix-sample-app", options...)
	assert.NilError(t, err)
	defer dbCloseFunc()

	ts := storage.NewCommentStorage(db)

	testFn(db, resolvedData, ts)
}

func TestDBCommentStorageGetComments(t *testing.T) {
	testDBCommentStorage(t, func(db *sql.DB, resolvedData *debefix.Data, ts storage.CommentStorage) {
		filter := entity.CommentFilter{
			Offset: 1,
			Limit:  1,
		}

		expectedComments, err := testdata.GetCommentList(
			testdata.WithFilterAll(true),
			testdata.WithSort("created_at"),
			testdata.WithOffsetLimit(filter.Offset, filter.Limit),
			testdata.WithResolvedData(resolvedData),
		)
		assert.NilError(t, err)
		assert.Assert(t, is.Len(expectedComments, 1))

		returnedComments, err := ts.GetCommentList(context.Background(), filter)
		assert.NilError(t, err)

		assert.Assert(t, is.Len(returnedComments, 1))
		assert.DeepEqual(t, expectedComments, returnedComments,
			opt.TimeWithThreshold(time.Hour))
	})
}

func TestDBCommentStorageGetCommentByID(t *testing.T) {
	testDBCommentStorage(t, func(db *sql.DB, resolvedData *debefix.Data, ts storage.CommentStorage) {
		expectedComment, err := testdata.GetComment(
			testdata.WithFilterRefIDs([]string{"test.DBCommentStorageTestMergeData"}),
			testdata.WithResolvedData(resolvedData),
		)
		assert.NilError(t, err)

		returnedComment, err := ts.GetCommentByID(context.Background(), expectedComment.CommentID)
		assert.NilError(t, err)

		assert.DeepEqual(t, expectedComment, returnedComment,
			opt.TimeWithThreshold(time.Hour))
	}, dbtest.WithDBForTestMergeData(dbCommentStorageTestMergeData()))
}

func TestDBCommentStorageAddComment(t *testing.T) {
	testDBCommentStorage(t, func(db *sql.DB, resolvedData *debefix.Data, ts storage.CommentStorage) {
		findPost, err := testdata.GetPost(
			testdata.WithFilterRefIDs([]string{"test.DBPostStorageTestMergeData"}),
			testdata.WithResolvedData(resolvedData),
		)
		assert.NilError(t, err)
		findUser, err := testdata.GetUser(
			testdata.WithFilterRefIDs([]string{"janedoe"}),
			testdata.WithResolvedData(resolvedData),
		)
		assert.NilError(t, err)

		newComment := entity.Comment{
			PostID: findPost.PostID,
			UserID: findUser.UserID,
			Text:   "new text",
		}

		returnedComment, err := ts.AddComment(context.Background(), newComment)
		assert.NilError(t, err)
		assert.Equal(t, "new text", returnedComment.Text)
	}, dbtest.WithDBForTestMergeData(dbPostStorageTestMergeData()))
}

func TestDBCommentStorageUpdateCommentByID(t *testing.T) {
	testDBCommentStorage(t, func(db *sql.DB, resolvedData *debefix.Data, ts storage.CommentStorage) {
		findPost, err := testdata.GetPost(
			testdata.WithFilterRefIDs([]string{"test.DBPostStorageTestMergeData"}),
			testdata.WithResolvedData(resolvedData),
		)
		assert.NilError(t, err)
		findUser, err := testdata.GetUser(
			testdata.WithFilterRefIDs([]string{"janedoe"}),
			testdata.WithResolvedData(resolvedData),
		)
		assert.NilError(t, err)

		expectedComment, err := testdata.GetComment(
			testdata.WithFilterRefIDs([]string{"test.DBCommentStorageTestMergeData"}),
			testdata.WithResolvedData(resolvedData),
		)
		assert.NilError(t, err)

		updatedComment := entity.Comment{
			PostID: findPost.PostID,
			UserID: findUser.UserID,
			Text:   "updated text",
		}

		returnedComment, err := ts.UpdateCommentByID(context.Background(), expectedComment.CommentID, updatedComment)
		assert.NilError(t, err)
		assert.Equal(t, "updated text", returnedComment.Text)
	}, dbtest.WithDBForTestMergeData(dbCommentStorageTestMergeData()))
}

func TestDBCommentStorageUpdateCommentByIDNotFound(t *testing.T) {
	testDBCommentStorage(t, func(db *sql.DB, resolvedData *debefix.Data, ts storage.CommentStorage) {
		findPost, err := testdata.GetPost(
			testdata.WithFilterRefIDs([]string{"test.DBPostStorageTestMergeData"}),
			testdata.WithResolvedData(resolvedData),
		)
		assert.NilError(t, err)
		findUser, err := testdata.GetUser(
			testdata.WithFilterRefIDs([]string{"janedoe"}),
			testdata.WithResolvedData(resolvedData),
		)
		assert.NilError(t, err)

		updatedComment := entity.Comment{
			PostID: findPost.PostID,
			UserID: findUser.UserID,
			Text:   "updated text",
		}

		_, err = ts.UpdateCommentByID(context.Background(), uuid.MustParse("0379ca21-7ed0-45e7-8812-4a6944f2c198"), updatedComment)
		assert.ErrorIs(t, err, util.ErrResourceNotFound)
	}, dbtest.WithDBForTestMergeData(dbCommentStorageTestMergeData()))
}

func TestDBCommentStorageDeleteCommentByID(t *testing.T) {
	testDBCommentStorage(t, func(db *sql.DB, resolvedData *debefix.Data, ts storage.CommentStorage) {
		expectedComment, err := testdata.GetComment(
			testdata.WithFilterRefIDs([]string{"test.DBCommentStorageTestMergeData"}),
			testdata.WithResolvedData(resolvedData),
		)
		assert.NilError(t, err)

		err = ts.DeleteCommentByID(context.Background(), expectedComment.CommentID)
		assert.NilError(t, err)
	}, dbtest.WithDBForTestMergeData(dbCommentStorageTestMergeData()))
}

func TestDBCommentStorageDeleteCommentByIDNotFound(t *testing.T) {
	testDBCommentStorage(t, func(db *sql.DB, resolvedData *debefix.Data, ts storage.CommentStorage) {
		err := ts.DeleteCommentByID(context.Background(), uuid.MustParse("0379ca21-7ed0-45e7-8812-4a6944f2c198"))
		assert.ErrorIs(t, err, util.ErrResourceNotFound)
	})
}

func dbCommentStorageTestMergeData() *debefix.Data {
	ret, err := debefix.MergeData(
		dbPostStorageTestMergeData(),
		&debefix.Data{
			Tables: map[string]*debefix.Table{
				"comments": {
					Rows: debefix.Rows{
						{
							Config: debefix.RowConfig{
								RefID:      "test.DBCommentStorageTestMergeData",
								IgnoreTags: true,
							},
							Fields: map[string]any{
								"comment_id": &debefix.ValueGenerated{},
								"post_id": &debefix.ValueRefID{
									TableID:   "posts",
									RefID:     "test.DBPostStorageTestMergeData",
									FieldName: "post_id",
								},
								"user_id": &debefix.ValueRefID{
									TableID:   "users",
									RefID:     "janedoe",
									FieldName: "user_id",
								},
								"text":       "Test Text",
								"created_at": time.Now(),
								"updated_at": time.Now(),
							},
						},
					},
				},
			},
		})
	if err != nil {
		panic(fmt.Sprintf("error merging test data: %s", err))
	}
	return ret
}
