package V6

import (
	"net/http"
	"strconv"

	"tutorial/go/Models"
	"tutorial/go/Repository"

	"github.com/gin-gonic/gin"
)

type UserODataController struct {
	repo Repository.Repository[Models.User]
}

func NewUserODataController(repo Repository.Repository[Models.User]) *UserODataController {
	return &UserODataController{repo: repo}
}

// GetAll godoc
// @Summary      Get all users (OData v6)
// @Description  Returns a paginated user collection using OData query options
// @Tags         users
// @Produce      json
// @Param        $top      query     int     false  "Number of records to return (default: 10)"
// @Param        $skip     query     int     false  "Number of records to skip (default: 0)"
// @Param        $orderby  query     string  false  "Sort field and direction"
// @Param        $filter   query     string  false  "OData filter"
// @Param        $search   query     string  false  "Search term matched against name and email"
// @Success      200  {object}  ODataPagedResponse[Models.User]
// @Failure      400  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /api/v6/odata/users [get]
func (h *UserODataController) GetAll(c *gin.Context) {
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

	c.JSON(http.StatusOK, ToODataResponse(c, result, "User"))
}

// GetByID godoc
// @Summary      Get user by ID (OData v6)
// @Description  Returns a single user with OData v4 context
// @Tags         users
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Router       /api/v6/odata/users/{id} [get]
func (h *UserODataController) GetByID(c *gin.Context) {
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
// @Summary      Get user count (OData v6)
// @Description  Returns the total number of users as plain text
// @Tags         users
// @Produce      plain
// @Param        $filter  query     string  false  "OData filter expression"
// @Success      200  {string}  string
// @Failure      400  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /api/v6/odata/users/$count [get]
func (h *UserODataController) Count(c *gin.Context) {
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

// Create godoc
// @Summary      Create a user
// @Description  Creates a new user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        body  body      Models.User  true  "User data"
// @Success      201   {object}  Models.User
// @Failure      400   {object}  ErrorResponse
// @Failure      500   {object}  ErrorResponse
// @Router       /api/v6/odata/users [post]
func (h *UserODataController) Create(c *gin.Context) {
	var user Models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}
	if err := h.repo.Create(&user); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}

// Update godoc
// @Summary      Update a user
// @Description  Updates a user by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id    path      int          true  "User ID"
// @Param        body  body      Models.User  true  "Updated user data"
// @Success      200   {object}  Models.User
// @Failure      400   {object}  ErrorResponse
// @Failure      500   {object}  ErrorResponse
// @Router       /api/v6/odata/users/{id} [put]
func (h *UserODataController) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid ID"})
		return
	}
	var user Models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}
	if err := h.repo.Update(uint(id), &user); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}
	updated, _ := h.repo.FindByID(uint(id))
	c.JSON(http.StatusOK, updated)
}

// Delete godoc
// @Summary      Delete a user
// @Description  Deletes a user by ID
// @Tags         users
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /api/v6/odata/users/{id} [delete]
func (h *UserODataController) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid ID"})
		return
	}
	if err := h.repo.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, MessageResponse{Message: "User deleted successfully"})
}
