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
	"github.com/rrgmc/debefix-sample-app/internal/domain/model"
	"github.com/rrgmc/debefix-sample-app/internal/domain/repository"
	"github.com/rrgmc/debefix-sample-app/internal/infrastructure/database"
	"github.com/rrgmc/debefix-sample-app/internal/infrastructure/database/testutils/dbtest"
	"github.com/rrgmc/debefix-sample-app/internal/infrastructure/database/testutils/testdata"
	"github.com/rrgmc/debefix-sample-app/internal/util"
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
	"gotest.tools/v3/assert/opt"
)

func testDBCommentRepository(t *testing.T, testFn func(*sql.DB, *debefix.Data, repository.CommentRepository),
	options ...dbtest.DBForTestOption) {
	db, resolvedData, dbCloseFunc, err := dbtest.DBForTest("debefix-sample-app", options...)
	assert.NilError(t, err)
	defer dbCloseFunc()

	ts := database.NewCommentRepository(db)

	testFn(db, resolvedData, ts)
}

func TestDBCommentRepositoryGetComments(t *testing.T) {
	testDBCommentRepository(t, func(db *sql.DB, resolvedData *debefix.Data, ts repository.CommentRepository) {
		filter := model.CommentFilter{
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

func TestDBCommentRepositoryGetCommentByID(t *testing.T) {
	testDBCommentRepository(t, func(db *sql.DB, resolvedData *debefix.Data, ts repository.CommentRepository) {
		expectedComment, err := testdata.GetComment(
			testdata.WithFilterRefIDs([]string{"test.DBCommentRepositoryTestMergeData"}),
			testdata.WithResolvedData(resolvedData),
		)
		assert.NilError(t, err)

		returnedComment, err := ts.GetCommentByID(context.Background(), expectedComment.CommentID)
		assert.NilError(t, err)

		assert.DeepEqual(t, expectedComment, returnedComment,
			opt.TimeWithThreshold(time.Hour))
	}, dbtest.WithDBForTestMergeData(dbCommentRepositoryTestMergeData()))
}

func TestDBCommentRepositoryAddComment(t *testing.T) {
	testDBCommentRepository(t, func(db *sql.DB, resolvedData *debefix.Data, ts repository.CommentRepository) {
		findPost, err := testdata.GetPost(
			testdata.WithFilterRefIDs([]string{"test.DBPostRepositoryTestMergeData"}),
			testdata.WithResolvedData(resolvedData),
		)
		assert.NilError(t, err)
		findUser, err := testdata.GetUser(
			testdata.WithFilterRefIDs([]string{"janedoe"}),
			testdata.WithResolvedData(resolvedData),
		)
		assert.NilError(t, err)

		newComment := model.Comment{
			PostID: findPost.PostID,
			UserID: findUser.UserID,
			Text:   "new text",
		}

		returnedComment, err := ts.AddComment(context.Background(), newComment)
		assert.NilError(t, err)
		assert.Equal(t, "new text", returnedComment.Text)
	}, dbtest.WithDBForTestMergeData(dbPostRepositoryTestMergeData()))
}

func TestDBCommentRepositoryUpdateCommentByID(t *testing.T) {
	testDBCommentRepository(t, func(db *sql.DB, resolvedData *debefix.Data, ts repository.CommentRepository) {
		findPost, err := testdata.GetPost(
			testdata.WithFilterRefIDs([]string{"test.DBPostRepositoryTestMergeData"}),
			testdata.WithResolvedData(resolvedData),
		)
		assert.NilError(t, err)
		findUser, err := testdata.GetUser(
			testdata.WithFilterRefIDs([]string{"janedoe"}),
			testdata.WithResolvedData(resolvedData),
		)
		assert.NilError(t, err)

		expectedComment, err := testdata.GetComment(
			testdata.WithFilterRefIDs([]string{"test.DBCommentRepositoryTestMergeData"}),
			testdata.WithResolvedData(resolvedData),
		)
		assert.NilError(t, err)

		updatedComment := model.Comment{
			PostID: findPost.PostID,
			UserID: findUser.UserID,
			Text:   "updated text",
		}

		returnedComment, err := ts.UpdateCommentByID(context.Background(), expectedComment.CommentID, updatedComment)
		assert.NilError(t, err)
		assert.Equal(t, "updated text", returnedComment.Text)
	}, dbtest.WithDBForTestMergeData(dbCommentRepositoryTestMergeData()))
}

func TestDBCommentRepositoryUpdateCommentByIDNotFound(t *testing.T) {
	testDBCommentRepository(t, func(db *sql.DB, resolvedData *debefix.Data, ts repository.CommentRepository) {
		findPost, err := testdata.GetPost(
			testdata.WithFilterRefIDs([]string{"test.DBPostRepositoryTestMergeData"}),
			testdata.WithResolvedData(resolvedData),
		)
		assert.NilError(t, err)
		findUser, err := testdata.GetUser(
			testdata.WithFilterRefIDs([]string{"janedoe"}),
			testdata.WithResolvedData(resolvedData),
		)
		assert.NilError(t, err)

		updatedComment := model.Comment{
			PostID: findPost.PostID,
			UserID: findUser.UserID,
			Text:   "updated text",
		}

		_, err = ts.UpdateCommentByID(context.Background(), uuid.MustParse("0379ca21-7ed0-45e7-8812-4a6944f2c198"), updatedComment)
		assert.ErrorIs(t, err, util.ErrResourceNotFound)
	}, dbtest.WithDBForTestMergeData(dbCommentRepositoryTestMergeData()))
}

func TestDBCommentRepositoryDeleteCommentByID(t *testing.T) {
	testDBCommentRepository(t, func(db *sql.DB, resolvedData *debefix.Data, ts repository.CommentRepository) {
		expectedComment, err := testdata.GetComment(
			testdata.WithFilterRefIDs([]string{"test.DBCommentRepositoryTestMergeData"}),
			testdata.WithResolvedData(resolvedData),
		)
		assert.NilError(t, err)

		err = ts.DeleteCommentByID(context.Background(), expectedComment.CommentID)
		assert.NilError(t, err)
	}, dbtest.WithDBForTestMergeData(dbCommentRepositoryTestMergeData()))
}

func TestDBCommentRepositoryDeleteCommentByIDNotFound(t *testing.T) {
	testDBCommentRepository(t, func(db *sql.DB, resolvedData *debefix.Data, ts repository.CommentRepository) {
		err := ts.DeleteCommentByID(context.Background(), uuid.MustParse("0379ca21-7ed0-45e7-8812-4a6944f2c198"))
		assert.ErrorIs(t, err, util.ErrResourceNotFound)
	})
}

func dbCommentRepositoryTestMergeData() *debefix.Data {
	ret, err := debefix.MergeData(
		dbPostRepositoryTestMergeData(),
		&debefix.Data{
			Tables: map[string]*debefix.Table{
				"comments": {
					Rows: debefix.Rows{
						{
							Config: debefix.RowConfig{
								RefID:      "test.DBCommentRepositoryTestMergeData",
								IgnoreTags: true,
							},
							Fields: map[string]any{
								"comment_id": &debefix.ValueGenerated{},
								"post_id": &debefix.ValueRefID{
									TableID:   "posts",
									RefID:     "test.DBPostRepositoryTestMergeData",
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