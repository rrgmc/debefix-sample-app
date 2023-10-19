package testdata

import (
	"fmt"
	"strings"

	"github.com/rrgmc/debefix"
	"github.com/rrgmc/debefix-sample-app/internal/entity"
)

func GetPostList(options ...TestDataOption) ([]entity.Post, error) {
	ret, err := filterData[entity.Post]("posts", func(row debefix.Row) (entity.Post, error) {
		return mapToStruct[entity.Post](row.Fields)
	}, func(sort string, a, b entity.Post) int {
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

func GetPost(options ...TestDataOption) (entity.Post, error) {
	ret, err := GetPostList(options...)
	if err != nil {
		return entity.Post{}, err
	}
	if len(ret) != 1 {
		return entity.Post{}, fmt.Errorf("incorrect amount of data returned for 'post': expected %d got %d", 1, len(ret))
	}
	return ret[0], nil
}
