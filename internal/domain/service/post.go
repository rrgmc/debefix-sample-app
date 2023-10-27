package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/rrgmc/debefix-sample-app/internal/domain/entity"
	"github.com/rrgmc/debefix-sample-app/internal/domain/repository"
	"github.com/rrgmc/debefix-sample-app/internal/domain/validator"
)

type PostService interface {
	GetPostList(ctx context.Context, filter entity.PostFilter) ([]entity.Post, error)
	GetPostByID(ctx context.Context, postID uuid.UUID) (entity.Post, error)
	AddPost(ctx context.Context, post entity.PostAdd) (entity.Post, error)
	UpdatePostByID(ctx context.Context, postID uuid.UUID, post entity.PostUpdate) (entity.Post, error)
	DeletePostByID(ctx context.Context, postID uuid.UUID) error
}

type postService struct {
	rctx           repository.Context
	postRepository repository.PostRepository
	postValidator  validator.PostValidator
}

func NewPostService(rctx repository.Context, postRepository repository.PostRepository) PostService {
	return &postService{
		rctx:           rctx,
		postRepository: postRepository,
		postValidator:  validator.NewPostValidator(),
	}
}

func (d postService) GetPostList(ctx context.Context, filter entity.PostFilter) ([]entity.Post, error) {
	err := d.postValidator.ValidatePostFilter(ctx, filter)
	if err != nil {
		return nil, err
	}
	return d.postRepository.GetPostList(ctx, d.rctx, filter)
}

func (d postService) GetPostByID(ctx context.Context, postID uuid.UUID) (entity.Post, error) {
	return d.postRepository.GetPostByID(ctx, d.rctx, postID)
}

func (d postService) AddPost(ctx context.Context, post entity.PostAdd) (entity.Post, error) {
	err := d.postValidator.ValidatePostAdd(ctx, post)
	if err != nil {
		return entity.Post{}, err
	}
	return d.postRepository.AddPost(ctx, d.rctx, post)
}

func (d postService) UpdatePostByID(ctx context.Context, postID uuid.UUID, post entity.PostUpdate) (entity.Post, error) {
	err := d.postValidator.ValidatePostUpdate(ctx, post)
	if err != nil {
		return entity.Post{}, err
	}
	return d.postRepository.UpdatePostByID(ctx, d.rctx, postID, post)
}

func (d postService) DeletePostByID(ctx context.Context, postID uuid.UUID) error {
	return d.postRepository.DeletePostByID(ctx, d.rctx, postID)
}
