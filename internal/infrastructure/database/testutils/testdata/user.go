package testdata

import (
	"fmt"
	"strings"

	"github.com/rrgmc/debefix"
	"github.com/rrgmc/debefix-sample-app/internal/domain/model"
	"github.com/rrgmc/debefix/filter"
)

func GetUserList(options ...TestDataOption) (filter.FilterDataRefIDResult[model.User], error) {
	ret, err := filterDataRefID[model.User]("users", func(row debefix.Row) (model.User, error) {
		return mapToStruct[model.User](row.Fields)
	}, func(sort string, a, b filter.FilterItem[model.User]) int {
		switch sort {
		case "name":
			return strings.Compare(a.Item.Name, b.Item.Name)
		default:
			panic(fmt.Sprintf("unknown sort '%s' for 'user' testdata", sort))
		}
	}, options...)
	if err != nil {
		return filter.FilterDataRefIDResult[model.User]{}, err
	}
	return ret, nil
}

func GetUser(options ...TestDataOption) (model.User, error) {
	ret, err := GetUserList(options...)
	if err != nil {
		return model.User{}, err
	}
	if len(ret.Data) != 1 {
		return model.User{}, fmt.Errorf("incorrect amount of data returned for 'user': expected %d got %d", 1, len(ret.Data))
	}
	return ret.Data[0], nil
}
