package entrypoints

import (
	"fmt"
	"go-simple-project/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Search Index For Value Endpoint
// @Description Searches for the index for the given value
// @Param value path int true "Value to look for."
// @Success 200 {object} services.SearchResultDto
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /numbers/{value} [get]
func SearchIndexEndpoint(c *gin.Context) {
	service, isOk := c.Request.Context().Value("SearchValueService").(*services.SearchValueService)
	if !isOk || service == nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Details: "User service not available"})
		return
	}

	searchValue := c.Param("value")
	num, err := strconv.Atoi(searchValue)

	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Details: fmt.Sprintf("Given value '%s' is not an integer", searchValue),
		})
		return
	}

	searchedValue := service.SearchIndex(num)

	if searchedValue.Found {
		c.JSON(http.StatusOK, searchedValue)
	} else {
		c.JSON(http.StatusNotFound, ErrorResponse{Details: "Index not found for this value."})
	}

}
