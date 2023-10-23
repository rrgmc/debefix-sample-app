package testdata

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/rrgmc/debefix"
	"github.com/rrgmc/debefix-sample-app/internal/domain/model"
)

func GetPostList(options ...TestDataOption) ([]model.Post, error) {
	ret, err := filterData[model.Post]("posts", func(row debefix.Row) (model.Post, error) {
		return mapToStruct[model.Post](row.Fields)
	}, func(sort string, a, b model.Post) int {
		switch sort {
		case "title":
			return strings.Compare(a.Title, b.Title)
		default:
			panic(fmt.Sprintf("unknown sort '%s' for 'post' testdata", sort))
		}
	}, options...)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(ret); i++ {
		ret[i].Tags, err = GetPostTagList(ret[i].PostID, options...)
		if err != nil {
			return nil, err
		}
	}
	return ret, nil
}

func GetPost(options ...TestDataOption) (model.Post, error) {
	ret, err := GetPostList(options...)
	if err != nil {
		return model.Post{}, err
	}
	if len(ret) != 1 {
		return model.Post{}, fmt.Errorf("incorrect amount of data returned for 'post': expected %d got %d", 1, len(ret))
	}
	return ret[0], nil
}

func GetPostTagList(postID uuid.UUID, options ...TestDataOption) ([]model.Tag, error) {
	optns := parseOptions(options...)

	postsTags, err := filterData[map[string]any]("posts_tags", func(row debefix.Row) (map[string]any, error) {
		return row.Fields, nil
	}, nil, WithFilterFields(map[string]any{
		"post_id": postID,
	}), WithResolvedData(optns.resolvedData))
	if err != nil {
		return nil, err
	}

	tags, err := GetTagList(WithFilterRow(func(row debefix.Row) (bool, error) {
		for _, pt := range postsTags {
			if row.Fields["tag_id"].(uuid.UUID) == pt["tag_id"].(uuid.UUID) {
				return true, nil
			}
		}
		return false, nil
	}), WithSort("name"), WithResolvedData(optns.resolvedData))
	if err != nil {
		return nil, err
	}
	if len(tags) != len(postsTags) {
		return nil, fmt.Errorf("unexpected amount of records returned, expected %d got %d", len(postsTags), len(tags))
	}
	return tags, nil
}
