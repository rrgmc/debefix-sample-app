package testdata

import (
	"fmt"
	"slices"
	"strings"

	"github.com/rrgmc/debefix"
	"github.com/rrgmc/debefix-sample-app/internal/entity"
)

func GetTags(options ...TestDataOption) []entity.Tag {
	ret := resolveData[entity.Tag]("tags", func(row debefix.Row) (entity.Tag, error) {
		return mapToStruct[entity.Tag](row.Fields)
	}, options...)

	optns := parseOptions(options...)
	if optns.sort != "" {
		slices.SortFunc(ret, func(a, b entity.Tag) int {
			switch optns.sort {
			case "name":
				return strings.Compare(a.Name, b.Name)
			default:
				panic(fmt.Sprintf("unknown sort '%s' for 'tag' testdata", optns.sort))
			}
		})
	}
	return ret
}

func GetTag(options ...TestDataOption) entity.Tag {
	ret := GetTags(options...)
	if len(ret) != 1 {
		panic(fmt.Sprintf("incorrect amount of data returned for 'tag': expected %d got %d", 1, len(ret)))
	}
	return ret[0]
}
