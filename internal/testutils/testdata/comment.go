package testdata

import (
	"fmt"

	"github.com/rrgmc/debefix"
	"github.com/rrgmc/debefix-sample-app/internal/entity"
)

func GetCommentList(options ...TestDataOption) ([]entity.Comment, error) {
	ret, err := filterData[entity.Comment]("comments", func(row debefix.Row) (entity.Comment, error) {
		return mapToStruct[entity.Comment](row.Fields)
	}, func(sort string, a, b entity.Comment) int {
		switch sort {
		case "created_at":
			if a.CreatedAt.Before(b.CreatedAt) {
				return -1
			} else if a.CreatedAt.After(b.CreatedAt) {
				return 1
			}
			return 0
		default:
			panic(fmt.Sprintf("unknown sort '%s' for 'comment' testdata", sort))
		}
	}, options...)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func GetComment(options ...TestDataOption) (entity.Comment, error) {
	ret, err := GetCommentList(options...)
	if err != nil {
		return entity.Comment{}, err
	}
	if len(ret) != 1 {
		return entity.Comment{}, fmt.Errorf("incorrect amount of data returned for 'comment': expected %d got %d", 1, len(ret))
	}
	return ret[0], nil
}
