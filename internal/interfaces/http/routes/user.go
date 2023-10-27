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

func UsersRoute(group *gin.RouterGroup, userService service.UserService) {
	// List
	group.GET("/", func(c *gin.Context) {
		request := payload.NewUserGetListRequest()

		if err := c.Bind(&request); err != nil {
			appError := domain.NewError(domain.ValidationError, err)
			_ = c.Error(appError)
			return
		}

		filter := mapper.UserFilterToEntity(request.Filter)

		list, err := userService.GetUserList(c, filter)
		if err != nil {
			_ = c.Error(err)
			return
		}

		output := mapper.UserListFromEntity(list)

		c.JSON(http.StatusOK, output)
	})

	// One
	group.GET("/:id", func(c *gin.Context) {
		userID := c.Param("id")
		userIDUUID, err := uuid.Parse(userID)
		if err != nil {
			_ = c.Error(domain.NewError(domain.ValidationError, errors.New("invalid uuid value")))
			return
		}

		item, err := userService.GetUserByID(c, userIDUUID)
		if err != nil {
			_ = c.Error(err)
			return
		}

		output := mapper.UserFromEntity(item)

		c.JSON(http.StatusOK, output)
	})

	// Add
	group.POST("/", func(c *gin.Context) {
		request := payload.UserAddRequest{}

		if err := c.Bind(&request.User); err != nil {
			appError := domain.NewError(domain.ValidationError, err)
			_ = c.Error(appError)
			return
		}

		user, err := mapper.UserAddToEntity(request.User)
		if err != nil {
			_ = c.Error(err)
			return
		}

		item, err := userService.AddUser(c, user)
		if err != nil {
			_ = c.Error(err)
			return
		}

		output := mapper.UserFromEntity(item)

		c.JSON(http.StatusOK, output)
	})

	// Update
	group.PUT("/:id", func(c *gin.Context) {
		userID := c.Param("id")
		userIDUUID, err := uuid.Parse(userID)
		if err != nil {
			_ = c.Error(domain.NewError(domain.ValidationError, errors.New("invalid uuid value")))
			return
		}

		request := payload.UserUpdateRequest{}

		if err := c.Bind(&request.User); err != nil {
			appError := domain.NewError(domain.ValidationError, err)
			_ = c.Error(appError)
			return
		}

		user, err := mapper.UserUpdateToEntity(request.User)
		if err != nil {
			_ = c.Error(err)
			return
		}

		item, err := userService.UpdateUserByID(c, userIDUUID, user)
		if err != nil {
			_ = c.Error(err)
			return
		}

		output := mapper.UserFromEntity(item)

		c.JSON(http.StatusOK, output)
	})

	// Delete
	group.DELETE("/:id", func(c *gin.Context) {
		userID := c.Param("id")
		userIDUUID, err := uuid.Parse(userID)
		if err != nil {
			_ = c.Error(domain.NewError(domain.ValidationError, errors.New("invalid uuid value")))
			return
		}

		err = userService.DeleteUserByID(c, userIDUUID)
		if err != nil {
			_ = c.Error(err)
			return
		}

		c.JSON(http.StatusOK, payload.Success{Message: "delete successful"})
	})
}
