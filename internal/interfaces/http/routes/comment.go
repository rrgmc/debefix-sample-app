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

func CommentsRoute(group *gin.RouterGroup, commentService service.CommentService) {
	// List
	group.GET("/", func(c *gin.Context) {
		request := payload.NewCommentGetListRequest()

		if err := c.Bind(&request); err != nil {
			appError := domain.NewError(domain.ValidationError, err)
			_ = c.Error(appError)
			return
		}

		filter, err := mapper.CommentFilterToEntity(request.Filter)
		if err != nil {
			_ = c.Error(err)
			return
		}

		list, err := commentService.GetCommentList(c, filter)
		if err != nil {
			_ = c.Error(err)
			return
		}

		output := mapper.CommentListFromEntity(list)

		c.JSON(http.StatusOK, output)
	})

	// One
	group.GET("/:id", func(c *gin.Context) {
		commentID := c.Param("id")
		commentIDUUID, err := uuid.Parse(commentID)
		if err != nil {
			_ = c.Error(domain.NewError(domain.ValidationError, errors.New("invalid uuid value")))
			return
		}

		item, err := commentService.GetCommentByID(c, commentIDUUID)
		if err != nil {
			_ = c.Error(err)
			return
		}

		output := mapper.CommentFromEntity(item)

		c.JSON(http.StatusOK, output)
	})

	// Add
	group.POST("/", func(c *gin.Context) {
		request := payload.CommentAddRequest{}

		if err := c.Bind(&request.Comment); err != nil {
			appError := domain.NewError(domain.ValidationError, err)
			_ = c.Error(appError)
			return
		}

		comment, err := mapper.CommentAddToEntity(request.Comment)
		if err != nil {
			_ = c.Error(err)
			return
		}

		item, err := commentService.AddComment(c, comment)
		if err != nil {
			_ = c.Error(err)
			return
		}

		output := mapper.CommentFromEntity(item)

		c.JSON(http.StatusOK, output)
	})

	// Update
	group.PUT("/:id", func(c *gin.Context) {
		commentID := c.Param("id")
		commentIDUUID, err := uuid.Parse(commentID)
		if err != nil {
			_ = c.Error(domain.NewError(domain.ValidationError, errors.New("invalid uuid value")))
			return
		}

		request := payload.CommentUpdateRequest{}

		if err := c.Bind(&request.Comment); err != nil {
			appError := domain.NewError(domain.ValidationError, err)
			_ = c.Error(appError)
			return
		}

		comment, err := mapper.CommentUpdateToEntity(request.Comment)
		if err != nil {
			_ = c.Error(err)
			return
		}

		item, err := commentService.UpdateCommentByID(c, commentIDUUID, comment)
		if err != nil {
			_ = c.Error(err)
			return
		}

		output := mapper.CommentFromEntity(item)

		c.JSON(http.StatusOK, output)
	})

	// Delete
	group.DELETE("/:id", func(c *gin.Context) {
		commentID := c.Param("id")
		commentIDUUID, err := uuid.Parse(commentID)
		if err != nil {
			_ = c.Error(domain.NewError(domain.ValidationError, errors.New("invalid uuid value")))
			return
		}

		err = commentService.DeleteCommentByID(c, commentIDUUID)
		if err != nil {
			_ = c.Error(err)
			return
		}

		c.JSON(http.StatusOK, payload.Success{Message: "delete successful"})
	})
}
