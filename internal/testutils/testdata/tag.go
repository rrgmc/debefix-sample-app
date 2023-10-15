package testdata

import (
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/RangelReale/debefix"
	"github.com/RangelReale/debefix-sample-app/internal/entity"
	"github.com/google/uuid"
)

func GetTags(refIDs []string, options ...TestDataOption) []entity.Tag {
	ret := resolveData[entity.Tag](refIDs, "tags", func(row debefix.Row) entity.Tag {
		return entity.Tag{
			TagID:     row.Fields["tag_id"].(uuid.UUID),
			Name:      row.Fields["name"].(string),
			CreatedAt: row.Fields["created_at"].(time.Time),
			UpdatedAt: row.Fields["updated_at"].(time.Time),
		}
	}, options...)

	optns := parseOptions(options...)
	if optns.sort != "" {
		slices.SortFunc(ret, func(a, b entity.Tag) int {
			switch optns.sort {
			case "name":
				return strings.Compare(a.Name, b.Name)
			default:
				panic(fmt.Sprintf("unknown sort '%s' for tag testdata", optns.sort))
			}
		})
	}
	return ret
}

func GetTag(refID string, options ...TestDataOption) entity.Tag {
	return GetTags([]string{refID}, options...)[0]
}
