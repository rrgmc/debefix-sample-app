package testdata

import (
	"fmt"
	"strings"

	"github.com/rrgmc/debefix"
	"github.com/rrgmc/debefix-sample-app/internal/entity"
)

func GetUserList(options ...TestDataOption) ([]entity.User, error) {
	ret, err := filterData[entity.User]("users", func(row debefix.Row) (entity.User, error) {
		return mapToStruct[entity.User](row.Fields)
	}, func(sort string, a, b entity.User) int {
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

func GetUser(options ...TestDataOption) (entity.User, error) {
	ret, err := GetUserList(options...)
	if err != nil {
		return entity.User{}, err
	}
	if len(ret) != 1 {
		return entity.User{}, fmt.Errorf("incorrect amount of data returned for 'user': expected %d got %d", 1, len(ret))
	}
	return ret[0], nil
}
