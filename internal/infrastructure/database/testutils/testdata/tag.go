package testdata

import (
	"fmt"
	"strings"

	"github.com/rrgmc/debefix"
	"github.com/rrgmc/debefix-sample-app/internal/domain/entity"
	"github.com/rrgmc/debefix/filter"
)

func GetTagList(options ...TestDataOption) (filter.FilterDataRefIDResult[entity.Tag], error) {
	ret, err := filterDataRefID[entity.Tag]("tags", func(row debefix.Row) (entity.Tag, error) {
		return mapToStruct[entity.Tag](row.Fields)
	}, func(sort string, a, b filter.FilterItem[entity.Tag]) int {
		switch sort {
		case "name":
			return strings.Compare(a.Item.Name, b.Item.Name)
		default:
			panic(fmt.Sprintf("unknown sort '%s' for 'tag' testdata", sort))
		}
	}, options...)
	if err != nil {
		return filter.FilterDataRefIDResult[entity.Tag]{}, err
	}
	return ret, nil
}

func GetTag(options ...TestDataOption) (entity.Tag, error) {
	ret, err := GetTagList(options...)
	if err != nil {
		return entity.Tag{}, err
	}
	if len(ret.Data) != 1 {
		return entity.Tag{}, fmt.Errorf("incorrect amount of data returned for 'tag': expected %d got %d", 1, len(ret.Data))
	}
	return ret.Data[0], nil
}
