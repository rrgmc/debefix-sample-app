package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/rrgmc/debefix-sample-app/internal/domain/entity"
)

type PostRepository interface {
	GetPostList(ctx context.Context, rctx Context, filter entity.PostFilter) ([]entity.Post, error)
	GetPostByID(ctx context.Context, rctx Context, postID uuid.UUID) (entity.Post, error)
	AddPost(ctx context.Context, rctx Context, post entity.Post) (entity.Post, error)
	UpdatePostByID(ctx context.Context, rctx Context, postID uuid.UUID, post entity.Post) (entity.Post, error)
	DeletePostByID(ctx context.Context, rctx Context, postID uuid.UUID) error
}
