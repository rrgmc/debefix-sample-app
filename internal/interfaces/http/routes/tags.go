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

func TagsRoute(group *gin.RouterGroup, tagService service.TagService) {
	// List
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

	// One
	group.GET("/:id", func(c *gin.Context) {
		tagID := c.Param("id")
		tagIDUUID, err := uuid.Parse(tagID)
		if err != nil {
			_ = c.Error(domain.NewError(domain.ValidationError, errors.New("invalid uuid value")))
			return
		}

		item, err := tagService.GetTagByID(c, tagIDUUID)
		if err != nil {
			_ = c.Error(err)
			return
		}

		output := mapper.TagFromEntity(item)

		c.JSON(http.StatusOK, output)
	})

	// Add
	group.POST("/", func(c *gin.Context) {
		request := payload.TagAddRequest{}

		if err := c.Bind(&request.Tag); err != nil {
			appError := domain.NewError(domain.ValidationError, err)
			_ = c.Error(appError)
			return
		}

		tag := mapper.TagAddToEntity(request.Tag)

		item, err := tagService.AddTag(c, tag)
		if err != nil {
			_ = c.Error(err)
			return
		}

		output := mapper.TagFromEntity(item)

		c.JSON(http.StatusOK, output)
	})

	// Update
	group.PUT("/:id", func(c *gin.Context) {
		tagID := c.Param("id")
		tagIDUUID, err := uuid.Parse(tagID)
		if err != nil {
			_ = c.Error(domain.NewError(domain.ValidationError, errors.New("invalid uuid value")))
			return
		}

		request := payload.TagUpdateRequest{}

		if err := c.Bind(&request.Tag); err != nil {
			appError := domain.NewError(domain.ValidationError, err)
			_ = c.Error(appError)
			return
		}

		tag := mapper.TagUpdateToEntity(request.Tag)

		item, err := tagService.UpdateTagByID(c, tagIDUUID, tag)
		if err != nil {
			_ = c.Error(err)
			return
		}

		output := mapper.TagFromEntity(item)

		c.JSON(http.StatusOK, output)
	})

	// Delete
	group.DELETE("/:id", func(c *gin.Context) {
		tagID := c.Param("id")
		tagIDUUID, err := uuid.Parse(tagID)
		if err != nil {
			_ = c.Error(domain.NewError(domain.ValidationError, errors.New("invalid uuid value")))
			return
		}

		err = tagService.DeleteTagByID(c, tagIDUUID)
		if err != nil {
			_ = c.Error(err)
			return
		}

		c.JSON(http.StatusOK, payload.Success{Message: "delete successful"})
	})
}
