package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/rrgmc/debefix-sample-app/internal/domain/model"
)

type PostRepository interface {
	GetPostList(ctx context.Context, filter model.PostFilter) ([]model.Post, error)
	GetPostByID(ctx context.Context, postID uuid.UUID) (model.Post, error)
	AddPost(ctx context.Context, post model.Post) (model.Post, error)
	UpdatePostByID(ctx context.Context, postID uuid.UUID, post model.Post) (model.Post, error)
	DeletePostByID(ctx context.Context, postID uuid.UUID) error
}
