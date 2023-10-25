package http

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rrgmc/debefix-sample-app/internal/domain/service"
	"github.com/rrgmc/debefix-sample-app/internal/interfaces/http/middlewares"
	"github.com/rrgmc/debefix-sample-app/internal/interfaces/http/routes"
)

func NewHTTPHandler(tagService service.TagService) http.Handler {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AddAllowHeaders("Authorization")
	router.Use(cors.New(config))

	router.Use(middlewares.Error)

	api := router.Group("/api/v1")

	routes.TagsRoute(api.Group("/tags"),
		tagService)

	return router
}
