package routes

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rrgmc/debefix-sample-app/internal/domain"
	"github.com/rrgmc/debefix-sample-app/internal/domain/service"
	"github.com/rrgmc/debefix-sample-app/internal/interfaces/http/mapper"
	"github.com/rrgmc/debefix-sample-app/internal/interfaces/http/payload"
)

func PostsRoute(group *gin.RouterGroup, postService service.PostService) {
	// List
	group.GET("/", func(c *gin.Context) {
		request := payload.NewPostGetListRequest()

		if err := c.Bind(&request); err != nil {
			appError := domain.NewError(domain.ValidationError, err)
			_ = c.Error(appError)
			return
		}

		filter, err := mapper.PostFilterToEntity(request.Filter)
		if err != nil {
			_ = c.Error(err)
			return
		}

		list, err := postService.GetPostList(c, filter)
		if err != nil {
			_ = c.Error(err)
			return
		}

		output := mapper.PostListFromEntity(list)

		c.JSON(http.StatusOK, output)
	})

	// One
	group.GET("/:id", func(c *gin.Context) {
		postID := c.Param("id")
		postIDUUID, err := uuid.Parse(postID)
		if err != nil {
			_ = c.Error(domain.NewError(domain.ValidationError, errors.New("invalid uuid value")))
			return
		}

		item, err := postService.GetPostByID(c, postIDUUID)
		if err != nil {
			_ = c.Error(err)
			return
		}

		output := mapper.PostFromEntity(item)

		c.JSON(http.StatusOK, output)
	})

	// Add
	group.POST("/", func(c *gin.Context) {
		request := payload.PostAddRequest{}

		if err := c.Bind(&request.Post); err != nil {
			appError := domain.NewError(domain.ValidationError, err)
			_ = c.Error(appError)
			return
		}

		post, err := mapper.PostAddToEntity(request.Post)
		if err != nil {
			_ = c.Error(err)
			return
		}

		item, err := postService.AddPost(c, post)
		if err != nil {
			_ = c.Error(err)
			return
		}

		output := mapper.PostFromEntity(item)

		c.JSON(http.StatusOK, output)
	})

	// Update
	group.PUT("/:id", func(c *gin.Context) {
		postID := c.Param("id")
		postIDUUID, err := uuid.Parse(postID)
		if err != nil {
			_ = c.Error(domain.NewError(domain.ValidationError, errors.New("invalid uuid value")))
			return
		}

		request := payload.PostUpdateRequest{}

		if err := c.Bind(&request.Post); err != nil {
			appError := domain.NewError(domain.ValidationError, err)
			_ = c.Error(appError)
			return
		}

		post, err := mapper.PostUpdateToEntity(request.Post)
		if err != nil {
			_ = c.Error(err)
			return
		}

		item, err := postService.UpdatePostByID(c, postIDUUID, post)
		if err != nil {
			_ = c.Error(err)
			return
		}

		output := mapper.PostFromEntity(item)

		c.JSON(http.StatusOK, output)
	})

	// Delete
	group.DELETE("/:id", func(c *gin.Context) {
		postID := c.Param("id")
		postIDUUID, err := uuid.Parse(postID)
		if err != nil {
			_ = c.Error(domain.NewError(domain.ValidationError, errors.New("invalid uuid value")))
			return
		}

		err = postService.DeletePostByID(c, postIDUUID)
		if err != nil {
			_ = c.Error(err)
			return
		}

		c.JSON(http.StatusOK, payload.Success{Message: "delete successful"})
	})
}
