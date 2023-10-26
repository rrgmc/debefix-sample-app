package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rrgmc/debefix-sample-app/internal/domain"
	"github.com/rrgmc/debefix-sample-app/internal/domain/service"
	"github.com/rrgmc/debefix-sample-app/internal/interfaces/http/mapper"
	"github.com/rrgmc/debefix-sample-app/internal/interfaces/http/payload"
)

func TagsRoute(group *gin.RouterGroup, tagService service.TagService) {
	group.GET("/", func(c *gin.Context) {
		var request payload.TagGetListRequest

		if err := c.Bind(&request); err != nil {
			appError := domain.NewError(domain.ValidationError, err)
			_ = c.Error(appError)
			return
		}

		filter := mapper.TagFilterToEntity(request.Filter)

		list, err := tagService.GetTagList(c, filter)
		if err != nil {
			_ = c.Error(err)
			return
		}
		c.JSON(http.StatusOK, list)
	})
}
