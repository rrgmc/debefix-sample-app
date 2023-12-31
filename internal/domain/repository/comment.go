package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/rrgmc/debefix-sample-app/internal/domain/entity"
)

type CommentRepository interface {
	GetCommentList(ctx context.Context, rctx Context, filter entity.CommentFilter) ([]entity.Comment, error)
	GetCommentByID(ctx context.Context, rctx Context, commentID uuid.UUID) (entity.Comment, error)
	ExistsCommentByID(ctx context.Context, rctx Context, commentID uuid.UUID) (bool, error)
	AddComment(ctx context.Context, rctx Context, comment entity.CommentAdd) (entity.Comment, error)
	UpdateCommentByID(ctx context.Context, rctx Context, commentID uuid.UUID, comment entity.CommentUpdate) (entity.Comment, error)
	DeleteCommentByID(ctx context.Context, rctx Context, commentID uuid.UUID) error
}
