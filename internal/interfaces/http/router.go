package http

import (
	"log/slog"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rrgmc/debefix-sample-app/internal/domain/service"
	"github.com/rrgmc/debefix-sample-app/internal/interfaces/http/middlewares"
	"github.com/rrgmc/debefix-sample-app/internal/interfaces/http/routes"
	sloggin "github.com/samber/slog-gin"
)

func NewHTTPHandler(logger *slog.Logger,
	tagService service.TagService,
	userService service.UserService,
	postService service.PostService,
) http.Handler {
	router := gin.New()

	router.Use(sloggin.New(logger))
	router.Use(gin.Recovery())

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
