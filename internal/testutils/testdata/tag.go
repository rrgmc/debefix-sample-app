package testdata

import (
	"fmt"
	"strings"

	"github.com/rrgmc/debefix"
	"github.com/rrgmc/debefix-sample-app/internal/entity"
)

func GetTags(options ...TestDataOption) ([]entity.Tag, error) {
	ret, err := filterData[entity.Tag]("tags", func(row debefix.Row) (entity.Tag, error) {
		return mapToStruct[entity.Tag](row.Fields)
	}, func(sort string, a, b entity.Tag) int {
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

func GetTag(options ...TestDataOption) (entity.Tag, error) {
	ret, err := GetTags(options...)
	if err != nil {
		return entity.Tag{}, err
	}
	if len(ret) != 1 {
		return entity.Tag{}, fmt.Errorf("incorrect amount of data returned for 'tag': expected %d got %d", 1, len(ret))
	}
	return ret[0], nil
}
