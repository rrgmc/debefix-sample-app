package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/rrgmc/debefix-sample-app/internal/domain/model"
)

type CommentRepository interface {
	GetCommentList(ctx context.Context, filter model.CommentFilter) ([]model.Comment, error)
	GetCommentByID(ctx context.Context, commentID uuid.UUID) (model.Comment, error)
	AddComment(ctx context.Context, comment model.Comment) (model.Comment, error)
	UpdateCommentByID(ctx context.Context, commentID uuid.UUID, comment model.Comment) (model.Comment, error)
	DeleteCommentByID(ctx context.Context, commentID uuid.UUID) error
}
