package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rrgmc/debefix-sample-app/internal/domain/service"
)

func TagsRoute(group *gin.RouterGroup, tagService service.TagService) {
	group.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "")
	})
}
