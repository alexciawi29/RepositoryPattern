package V2

import (
	"net/http"
	"strconv"

	"tutorial/go/Models"
	"tutorial/go/Repository"

	"github.com/gin-gonic/gin"
)

type UserODataHandler struct {
	repo Repository.Repository[Models.User]
}

func NewUserODataHandler(repo Repository.Repository[Models.User]) *UserODataHandler {
	return &UserODataHandler{repo: repo}
}

// GetAll godoc
// @Summary      Get all users (OData v4)
// @Description  Returns a paginated user collection using OData v4 query options
// @Tags         users
// @Produce      json
// @Param        $top      query     int     false  "Number of records to return (default: 10)"
// @Param        $skip     query     int     false  "Number of records to skip (default: 0)"
// @Param        $orderby  query     string  false  "Sort field and direction (e.g. name asc, email desc)"
// @Param        $filter   query     string  false  "OData filter (e.g. name eq 'Herris', contains(email,'@example.com'))"
// @Param        $search   query     string  false  "Search term matched against name and email"
// @Success      200  {object}  ODataPagedResponse[Models.User]
// @Failure      400  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /api/v2/users [get]
func (h *UserODataHandler) GetAll(c *gin.Context) {
	top, _ := strconv.Atoi(c.DefaultQuery("$top", "10"))
	skip, _ := strconv.Atoi(c.DefaultQuery("$skip", "0"))

	sortBy, order := ParseODataOrderBy(c.Query("$orderby"))

	conditions, err := ParseODataFilter(c.Query("$filter"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	params := Repository.QueryParams{
		Top:          top,
		Skip:         skip,
		SortBy:       sortBy,
		Order:        order,
		Conditions:   conditions,
		Search:       c.Query("$search"),
		SearchFields: []string{"name", "email"},
	}

	result, err := h.repo.FindAll(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, ToODataResponse(c, result, "Users"))
}

// GetByID godoc
// @Summary      Get user by ID (OData v4)
// @Description  Returns a single user with OData v4 context
// @Tags         users
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Router       /api/v2/users/{id} [get]
func (h *UserODataHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid ID"})
		return
	}

	user, err := h.repo.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Error: "User not found"})
		return
	}

	resp, err := ToODataEntityResponse(user, ContextURL(c, "Users/$entity"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// Count godoc
// @Summary      Get user count (OData v4)
// @Description  Returns the total number of users as plain text
// @Tags         users
// @Produce      plain
// @Param        $filter  query     string  false  "OData filter expression"
// @Success      200  {string}  string
// @Failure      400  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /api/v2/users/$count [get]
func (h *UserODataHandler) Count(c *gin.Context) {
	conditions, err := ParseODataFilter(c.Query("$filter"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	result, err := h.repo.FindAll(Repository.QueryParams{
		Top:        1,
		Conditions: conditions,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.Data(http.StatusOK, "text/plain", []byte(strconv.FormatInt(result.Total, 10)))
}
