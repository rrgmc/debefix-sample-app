package testdata

import (
	"fmt"
	"strings"

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
