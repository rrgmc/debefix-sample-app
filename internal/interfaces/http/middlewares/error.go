package middlewares

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rrgmc/debefix-sample-app/internal/domain"
	"github.com/rrgmc/debefix-sample-app/internal/interfaces/http/payload"
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
				c.JSON(http.StatusNotFound, payload.Error{ErrorMessage: err.Error()})
				return
			} else if errors.Is(err, domain.ValidationError) {
				var verr validator.ValidationErrors
				if errors.As(err, &verr) {
					verrorMsg := map[string]string{}
					for _, field := range verr {
						verrorMsg[field.Field()] = field.Error()
					}

					c.JSON(http.StatusBadRequest, payload.ErrorValidation{
						Error: payload.Error{
							ErrorMessage: domain.ValidationError.Error(),
						},
						Validation: verrorMsg,
					})
					return
				}
				c.JSON(http.StatusBadRequest, payload.Error{ErrorMessage: err.Error()})
				return
			} else if errors.Is(err, domain.ResourceAlreadyExists) {
				c.JSON(http.StatusConflict, payload.Error{ErrorMessage: err.Error()})
				return
			} else if errors.Is(err, domain.NotAuthenticated) {
				c.JSON(http.StatusUnauthorized, payload.Error{ErrorMessage: err.Error()})
				return
			} else if errors.Is(err, domain.NotAuthorized) {
				c.JSON(http.StatusForbidden, payload.Error{ErrorMessage: err.Error()})
				return
			} else if errors.Is(err, domain.RepositoryError) {
				c.JSON(http.StatusInternalServerError, payload.Error{ErrorMessage: err.Error()})
				return
			}
		}

		// Error is not AppError, return a generic internal server error
		c.JSON(http.StatusInternalServerError, payload.Error{ErrorMessage: "Internal Server Errror"})
		return
	}
}
