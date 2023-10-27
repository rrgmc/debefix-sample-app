package mapper

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/rrgmc/debefix-sample-app/internal/domain"
	"github.com/rrgmc/debefix-sample-app/internal/domain/entity"
	"github.com/rrgmc/debefix-sample-app/internal/interfaces/http/payload"
)

// From entity

func CommentFromEntity(comment entity.Comment) payload.Comment {
	return payload.Comment{
		CommentID: comment.CommentID.String(),
		PostID:    comment.PostID.String(),
		UserID:    comment.UserID.String(),
		Text:      comment.Text,
		CreatedAt: comment.CreatedAt.Format(time.RFC3339),
		UpdatedAt: comment.UpdatedAt.Format(time.RFC3339),
	}
}

func CommentListFromEntity(commentList []entity.Comment) []payload.Comment {
	var list []payload.Comment
	for _, item := range commentList {
		list = append(list, CommentFromEntity(item))
	}
	return list
}

// To entity

func CommentAddToEntity(comment payload.CommentAdd) (entity.CommentAdd, error) {
	postID, err := uuid.Parse(comment.PostID)
	if err != nil {
		return entity.CommentAdd{}, domain.NewError(domain.ValidationError, errors.New("post_id: invalid uuid value"))
	}

	userID, err := uuid.Parse(comment.UserID)
	if err != nil {
		return entity.CommentAdd{}, domain.NewError(domain.ValidationError, errors.New("user_id: invalid uuid value"))
	}

	return entity.CommentAdd{
		PostID: postID,
		UserID: userID,
		Text:   comment.Text,
	}, nil
}

func CommentUpdateToEntity(comment payload.CommentUpdate) (entity.CommentUpdate, error) {
	return CommentAddToEntity(comment)
}

func CommentFilterToEntity(commentFilter payload.CommentFilter) (entity.CommentFilter, error) {
	ret := entity.CommentFilter{
		Offset: commentFilter.Offset,
		Limit:  commentFilter.Limit,
	}

	if commentFilter.PostID != nil {
		postID, err := uuid.Parse(*commentFilter.PostID)
		if err != nil {
			return entity.CommentFilter{}, domain.NewError(domain.ValidationError, errors.New("post_id: invalid uuid value"))
		}
		ret.PostID = &postID
	}

	if commentFilter.UserID != nil {
		userID, err := uuid.Parse(*commentFilter.UserID)
		if err != nil {
			return entity.CommentFilter{}, domain.NewError(domain.ValidationError, errors.New("user_id: invalid uuid value"))
		}
		ret.UserID = &userID
	}

	return ret, nil
}
