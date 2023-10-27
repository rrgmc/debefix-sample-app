package mapper

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/rrgmc/debefix-sample-app/internal/domain"
	"github.com/rrgmc/debefix-sample-app/internal/domain/entity"
	"github.com/rrgmc/debefix-sample-app/internal/interfaces/http/payload"
	"github.com/rrgmc/debefix-sample-app/internal/utils"
)

// From entity

func PostFromEntity(post entity.Post) payload.Post {
	return payload.Post{
		PostID:    post.PostID.String(),
		Title:     post.Title,
		Text:      post.Text,
		UserID:    post.UserID.String(),
		Tags:      TagListFromEntity(post.Tags),
		CreatedAt: post.CreatedAt.Format(time.RFC3339),
		UpdatedAt: post.UpdatedAt.Format(time.RFC3339),
	}
}

func PostListFromEntity(postList []entity.Post) []payload.Post {
	var list []payload.Post
	for _, item := range postList {
		list = append(list, PostFromEntity(item))
	}
	return list
}

// To entity

func PostAddToEntity(post payload.PostAdd) (entity.PostAdd, error) {
	userID, err := uuid.Parse(post.UserID)
	if err != nil {
		return entity.PostAdd{}, domain.NewError(domain.ValidationError, errors.New("user_id: invalid uuid value"))
	}

	tags, err := utils.UUIDListParse(post.Tags)
	if err != nil {
		return entity.PostAdd{}, domain.NewError(domain.ValidationError, fmt.Errorf("tags: %w", err))
	}

	return entity.PostAdd{
		Title:  post.Title,
		Text:   post.Text,
		UserID: userID,
		Tags:   tags,
	}, nil
}

func PostUpdateToEntity(post payload.PostUpdate) (entity.PostUpdate, error) {
	return PostAddToEntity(post)
}

func PostFilterToEntity(postFilter payload.PostFilter) (entity.PostFilter, error) {
	ret := entity.PostFilter{
		Offset: postFilter.Offset,
		Limit:  postFilter.Limit,
	}

	if postFilter.UserID != nil {
		userID, err := uuid.Parse(*postFilter.UserID)
		if err != nil {
			return entity.PostFilter{}, domain.NewError(domain.ValidationError, errors.New("user_id: invalid uuid value"))
		}
		ret.UserID = &userID
	}

	return ret, nil
}
