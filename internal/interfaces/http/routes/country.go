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

func CountriesRoute(group *gin.RouterGroup, countryService service.CountryService) {
	// List
	group.GET("/", func(c *gin.Context) {
		request := payload.NewCountryGetListRequest()

		if err := c.Bind(&request); err != nil {
			appError := domain.NewError(domain.ValidationError, err)
			_ = c.Error(appError)
			return
		}

		filter := mapper.CountryFilterToEntity(request.Filter)

		list, err := countryService.GetCountryList(c, filter)
		if err != nil {
			_ = c.Error(err)
			return
		}

		output := mapper.CountryListFromEntity(list)

		c.JSON(http.StatusOK, output)
	})

	// One
	group.GET("/:id", func(c *gin.Context) {
		countryID := c.Param("id")
		countryIDUUID, err := uuid.Parse(countryID)
		if err != nil {
			_ = c.Error(domain.NewError(domain.ValidationError, errors.New("invalid uuid value")))
			return
		}

		item, err := countryService.GetCountryByID(c, countryIDUUID)
		if err != nil {
			_ = c.Error(err)
			return
		}

		output := mapper.CountryFromEntity(item)

		c.JSON(http.StatusOK, output)
	})
}
