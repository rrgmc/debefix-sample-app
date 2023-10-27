package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/rrgmc/debefix-sample-app/internal/domain/entity"
)

type CommentService interface {
	GetCommentList(ctx context.Context, filter entity.CommentFilter) ([]entity.Comment, error)
	GetCommentByID(ctx context.Context, commentID uuid.UUID) (entity.Comment, error)
	ExistsCommentByID(ctx context.Context, commentID uuid.UUID) (bool, error)
	AddComment(ctx context.Context, comment entity.CommentAdd) (entity.Comment, error)
	UpdateCommentByID(ctx context.Context, commentID uuid.UUID, comment entity.CommentUpdate) (entity.Comment, error)
	DeleteCommentByID(ctx context.Context, commentID uuid.UUID) error
}
