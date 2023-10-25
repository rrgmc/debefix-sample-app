package testdata

import (
	"fmt"

	"github.com/rrgmc/debefix"
	"github.com/rrgmc/debefix-sample-app/internal/domain/entity"
	"github.com/rrgmc/debefix/filter"
)

func GetCommentList(options ...TestDataOption) (filter.FilterDataRefIDResult[entity.Comment], error) {
	ret, err := filterDataRefID[entity.Comment]("comments", func(row debefix.Row) (entity.Comment, error) {
		return mapToStruct[entity.Comment](row.Fields)
	}, func(sort string, a, b filter.FilterItem[entity.Comment]) int {
		switch sort {
		case "created_at":
			if a.Item.CreatedAt.Before(b.Item.CreatedAt) {
				return -1
			} else if a.Item.CreatedAt.After(b.Item.CreatedAt) {
				return 1
			}
			return 0
		default:
			panic(fmt.Sprintf("unknown sort '%s' for 'comment' testdata", sort))
		}
	}, options...)
	if err != nil {
		return filter.FilterDataRefIDResult[entity.Comment]{}, err
	}
	return ret, nil
}

func GetComment(options ...TestDataOption) (entity.Comment, error) {
	ret, err := GetCommentList(options...)
	if err != nil {
		return entity.Comment{}, err
	}
	if len(ret.Data) != 1 {
		return entity.Comment{}, fmt.Errorf("incorrect amount of data returned for 'comment': expected %d got %d", 1, len(ret.Data))
	}
	return ret.Data[0], nil
}
