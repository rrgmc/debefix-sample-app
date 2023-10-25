package testdata

import (
	"fmt"

	"github.com/rrgmc/debefix"
	"github.com/rrgmc/debefix-sample-app/internal/domain/model"
	"github.com/rrgmc/debefix/filter"
)

func GetCommentList(options ...TestDataOption) (filter.FilterDataRefIDResult[model.Comment], error) {
	ret, err := filterDataRefID[model.Comment]("comments", func(row debefix.Row) (model.Comment, error) {
		return mapToStruct[model.Comment](row.Fields)
	}, func(sort string, a, b filter.FilterItem[model.Comment]) int {
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
		return filter.FilterDataRefIDResult[model.Comment]{}, err
	}
	return ret, nil
}

func GetComment(options ...TestDataOption) (model.Comment, error) {
	ret, err := GetCommentList(options...)
	if err != nil {
		return model.Comment{}, err
	}
	if len(ret.Data) != 1 {
		return model.Comment{}, fmt.Errorf("incorrect amount of data returned for 'comment': expected %d got %d", 1, len(ret.Data))
	}
	return ret.Data[0], nil
}
