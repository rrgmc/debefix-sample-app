package serviceimpl

import (
	"context"

	"github.com/google/uuid"
	"github.com/rrgmc/debefix-sample-app/internal/domain/entity"
	"github.com/rrgmc/debefix-sample-app/internal/domain/repository"
	"github.com/rrgmc/debefix-sample-app/internal/domain/service"
	"github.com/rrgmc/debefix-sample-app/internal/domain/validator"
)

type commentService struct {
	rctx              repository.Context
	commentRepository repository.CommentRepository
	commentValidator  validator.CommentValidator
}

func NewCommentService(rctx repository.Context, commentRepository repository.CommentRepository,
	postService service.PostService, userService service.UserService) service.CommentService {
	return &commentService{
		rctx:              rctx,
		commentRepository: commentRepository,
		commentValidator:  validator.NewCommentValidator(postService, userService),
	}
}

func (d commentService) GetCommentList(ctx context.Context, filter entity.CommentFilter) ([]entity.Comment, error) {
	err := d.commentValidator.ValidateCommentFilter(ctx, filter)
	if err != nil {
		return nil, err
	}
	return d.commentRepository.GetCommentList(ctx, d.rctx, filter)
}

func (d commentService) GetCommentByID(ctx context.Context, commentID uuid.UUID) (entity.Comment, error) {
	return d.commentRepository.GetCommentByID(ctx, d.rctx, commentID)
}

func (d commentService) ExistsCommentByID(ctx context.Context, commentID uuid.UUID) (bool, error) {
	return d.commentRepository.ExistsCommentByID(ctx, d.rctx, commentID)
}

func (d commentService) AddComment(ctx context.Context, comment entity.CommentAdd) (entity.Comment, error) {
	err := d.commentValidator.ValidateCommentAdd(ctx, comment)
	if err != nil {
		return entity.Comment{}, err
	}
	return d.commentRepository.AddComment(ctx, d.rctx, comment)
}

func (d commentService) UpdateCommentByID(ctx context.Context, commentID uuid.UUID, comment entity.CommentUpdate) (entity.Comment, error) {
	err := d.commentValidator.ValidateCommentUpdate(ctx, comment)
	if err != nil {
		return entity.Comment{}, err
	}
	return d.commentRepository.UpdateCommentByID(ctx, d.rctx, commentID, comment)
}

func (d commentService) DeleteCommentByID(ctx context.Context, commentID uuid.UUID) error {
	return d.commentRepository.DeleteCommentByID(ctx, d.rctx, commentID)
}
