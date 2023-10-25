package middlewares

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rrgmc/debefix-sample-app/internal/domain"
)

// Error is Gin middleware to handle errors.
func Error(c *gin.Context) {
	// Execute request handlers and then handle any errors
	c.Next()

	errs := c.Errors

	if len(errs) > 0 {
		var err *domain.Error
		if errors.As(errs[0].Err, &err) {
			if errors.Is(err, domain.NotFound) {
				c.JSON(http.StatusNotFound, err.Error())
				return
			} else if errors.Is(err, domain.ValidationError) {
				c.JSON(http.StatusBadRequest, err.Error())
				return
			} else if errors.Is(err, domain.ResourceAlreadyExists) {
				c.JSON(http.StatusConflict, err.Error())
				return
			} else if errors.Is(err, domain.NotAuthenticated) {
				c.JSON(http.StatusUnauthorized, err.Error())
				return
			} else if errors.Is(err, domain.NotAuthorized) {
				c.JSON(http.StatusForbidden, err.Error())
				return
			} else if errors.Is(err, domain.RepositoryError) {
				c.JSON(http.StatusInternalServerError, err.Error())
				return
			}
		}

		// Error is not AppError, return a generic internal server error
		c.JSON(http.StatusInternalServerError, "Internal Server Errror")
		return
	}
}
