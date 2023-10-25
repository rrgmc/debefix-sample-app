package testdata

import (
	"fmt"
	"strings"

	"github.com/rrgmc/debefix"
	"github.com/rrgmc/debefix-sample-app/internal/domain/entity"
	"github.com/rrgmc/debefix/filter"
)

func GetUserList(options ...TestDataOption) (filter.FilterDataRefIDResult[entity.User], error) {
	ret, err := filterDataRefID[entity.User]("users", func(row debefix.Row) (entity.User, error) {
		return mapToStruct[entity.User](row.Fields)
	}, func(sort string, a, b filter.FilterItem[entity.User]) int {
		switch sort {
		case "name":
			return strings.Compare(a.Item.Name, b.Item.Name)
		default:
			panic(fmt.Sprintf("unknown sort '%s' for 'user' testdata", sort))
		}
	}, options...)
	if err != nil {
		return filter.FilterDataRefIDResult[entity.User]{}, err
	}
	return ret, nil
}

func GetUser(options ...TestDataOption) (entity.User, error) {
	ret, err := GetUserList(options...)
	if err != nil {
		return entity.User{}, err
	}
	if len(ret.Data) != 1 {
		return entity.User{}, fmt.Errorf("incorrect amount of data returned for 'user': expected %d got %d", 1, len(ret.Data))
	}
	return ret.Data[0], nil
}
