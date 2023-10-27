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
	countryService service.CountryService,
	tagService service.TagService,
	userService service.UserService,
	postService service.PostService,
	commentService service.CommentService,
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

	routes.CountriesRoute(
		api.Group("/countries"),
		countryService)
	routes.TagsRoute(
		api.Group("/tags"),
		tagService)
	routes.UsersRoute(
		api.Group("/users"),
		userService)
	routes.PostsRoute(
		api.Group("/posts"),
		postService)
	routes.CommentsRoute(
		api.Group("/comments"),
		commentService)

	return router
}
