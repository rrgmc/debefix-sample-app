package testdata

import (
	"fmt"
	"strings"

	"github.com/rrgmc/debefix"
	"github.com/rrgmc/debefix-sample-app/internal/domain/model"
	"github.com/rrgmc/debefix/filter"
)

func GetTagList(options ...TestDataOption) (filter.FilterDataRefIDResult[model.Tag], error) {
	ret, err := filterDataRefID[model.Tag]("tags", func(row debefix.Row) (model.Tag, error) {
		return mapToStruct[model.Tag](row.Fields)
	}, func(sort string, a, b filter.FilterItem[model.Tag]) int {
		switch sort {
		case "name":
			return strings.Compare(a.Item.Name, b.Item.Name)
		default:
			panic(fmt.Sprintf("unknown sort '%s' for 'tag' testdata", sort))
		}
	}, options...)
	if err != nil {
		return filter.FilterDataRefIDResult[model.Tag]{}, err
	}
	return ret, nil
}

func GetTag(options ...TestDataOption) (model.Tag, error) {
	ret, err := GetTagList(options...)
	if err != nil {
		return model.Tag{}, err
	}
	if len(ret.Data) != 1 {
		return model.Tag{}, fmt.Errorf("incorrect amount of data returned for 'tag': expected %d got %d", 1, len(ret.Data))
	}
	return ret.Data[0], nil
}
