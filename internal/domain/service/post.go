package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/rrgmc/debefix-sample-app/internal/domain/entity"
)

type PostService interface {
	GetPostList(ctx context.Context, filter entity.PostFilter) ([]entity.Post, error)
	GetPostByID(ctx context.Context, postID uuid.UUID) (entity.Post, error)
	ExistsPostByID(ctx context.Context, postID uuid.UUID) (bool, error)
	AddPost(ctx context.Context, post entity.PostAdd) (entity.Post, error)
	UpdatePostByID(ctx context.Context, postID uuid.UUID, post entity.PostUpdate) (entity.Post, error)
	DeletePostByID(ctx context.Context, postID uuid.UUID) error
}
