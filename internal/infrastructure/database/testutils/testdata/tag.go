package testdata

import (
	"fmt"
	"strings"

	"github.com/rrgmc/debefix"
	"github.com/rrgmc/debefix-sample-app/internal/domain/model"
)

func GetTagList(options ...TestDataOption) ([]model.Tag, error) {
	ret, err := filterData[model.Tag]("tags", func(row debefix.Row) (model.Tag, error) {
		return mapToStruct[model.Tag](row.Fields)
	}, func(sort string, a, b model.Tag) int {
		switch sort {
		case "name":
			return strings.Compare(a.Name, b.Name)
		default:
			panic(fmt.Sprintf("unknown sort '%s' for 'tag' testdata", sort))
		}
	}, options...)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func GetTag(options ...TestDataOption) (model.Tag, error) {
	ret, err := GetTagList(options...)
	if err != nil {
		return model.Tag{}, err
	}
	if len(ret) != 1 {
		return model.Tag{}, fmt.Errorf("incorrect amount of data returned for 'tag': expected %d got %d", 1, len(ret))
	}
	return ret[0], nil
}