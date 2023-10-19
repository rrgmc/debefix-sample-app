package storage

import (
	"context"

	"github.com/google/uuid"
	"github.com/rrgmc/debefix-sample-app/internal/entity"
)

type CommentStorage interface {
	GetCommentList(ctx context.Context, filter entity.CommentFilter) ([]entity.Comment, error)
	GetCommentByID(ctx context.Context, commentID uuid.UUID) (entity.Comment, error)
	AddComment(ctx context.Context, comment entity.Comment) (entity.Comment, error)
	UpdateCommentByID(ctx context.Context, commentID uuid.UUID, comment entity.Comment) (entity.Comment, error)
	DeleteCommentByID(ctx context.Context, commentID uuid.UUID) error
}

type CountryStorage interface {
	GetCountryList(ctx context.Context, filter entity.CountryFilter) ([]entity.Country, error)
	GetCountryByID(ctx context.Context, CountryID uuid.UUID) (entity.Country, error)
}

type PostStorage interface {
	GetPostList(ctx context.Context, filter entity.PostFilter) ([]entity.Post, error)
	GetPostByID(ctx context.Context, postID uuid.UUID) (entity.Post, error)
	AddPost(ctx context.Context, post entity.Post) (entity.Post, error)
	UpdatePostByID(ctx context.Context, postID uuid.UUID, post entity.Post) (entity.Post, error)
	DeletePostByID(ctx context.Context, postID uuid.UUID) error
}

type TagStorage interface {
	GetTagList(ctx context.Context, filter entity.TagFilter) ([]entity.Tag, error)
	GetTagByID(ctx context.Context, tagID uuid.UUID) (entity.Tag, error)
	AddTag(ctx context.Context, tag entity.Tag) (entity.Tag, error)
	UpdateTagByID(ctx context.Context, tagID uuid.UUID, tag entity.Tag) (entity.Tag, error)
	DeleteTagByID(ctx context.Context, tagID uuid.UUID) error
}

type UserStorage interface {
	GetUserList(ctx context.Context, filter entity.UserFilter) ([]entity.User, error)
	GetUserByID(ctx context.Context, userID uuid.UUID) (entity.User, error)
	AddUser(ctx context.Context, user entity.User) (entity.User, error)
	UpdateUserByID(ctx context.Context, userID uuid.UUID, user entity.User) (entity.User, error)
	DeleteUserByID(ctx context.Context, userID uuid.UUID) error
}
