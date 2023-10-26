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
		request := payload.NewTagGetListRequest()

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

		output := mapper.TagListFromEntity(list)

		c.JSON(http.StatusOK, output)
	})
}
