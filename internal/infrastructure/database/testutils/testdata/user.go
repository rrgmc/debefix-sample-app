package testdata

import (
	"fmt"
	"strings"

	"github.com/rrgmc/debefix"
	"github.com/rrgmc/debefix-sample-app/internal/domain/model"
)

func GetUserList(options ...TestDataOption) ([]model.User, error) {
	ret, err := filterData[model.User]("users", func(row debefix.Row) (model.User, error) {
		return mapToStruct[model.User](row.Fields)
	}, func(sort string, a, b model.User) int {
		switch sort {
		case "name":
			return strings.Compare(a.Name, b.Name)
		default:
			panic(fmt.Sprintf("unknown sort '%s' for 'user' testdata", sort))
		}
	}, options...)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func GetUser(options ...TestDataOption) (model.User, error) {
	ret, err := GetUserList(options...)
	if err != nil {
		return model.User{}, err
	}
	if len(ret) != 1 {
		return model.User{}, fmt.Errorf("incorrect amount of data returned for 'user': expected %d got %d", 1, len(ret))
	}
	return ret[0], nil
}
