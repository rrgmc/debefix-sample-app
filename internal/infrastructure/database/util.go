package database

import (
	"fmt"
)

//nolint:unparam
func getColumnsWithAlias(columns []any, alias string) []any {
	var ret []any
	for _, col := range columns {
		ret = append(ret, fmt.Sprintf("%s.%s", alias, col.(string)))
	}
	return ret
}
