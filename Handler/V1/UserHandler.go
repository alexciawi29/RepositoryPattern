package V1

import (
	"net/http"
	"strconv"

	"tutorial/go/Models"
	"tutorial/go/Repository"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	repo Repository.Repository[Models.User]
}

func NewUserHandler(repo Repository.Repository[Models.User]) *UserHandler {
	return &UserHandler{repo: repo}
}

type CreateUserRequest struct {
	Name  string `json:"Name"  binding:"required"       example:"Herris"`
	Email string `json:"Email" binding:"required,email" example:"herris@example.com"`
}

type UpdateUserRequest struct {
	Name  string `json:"Name"  example:"Herris Updated"`
	Email string `json:"Email" example:"herris.new@example.com"`
}

// GetAll godoc
// @Summary      Get all users
// @Description  Returns a paginated list of users with optional sorting, filtering, and search
// @Tags         users
// @Produce      json
// @Param        top     query     int     false  "Number of records to return (default: 10)"
// @Param        skip    query     int     false  "Number of records to skip (default: 0)"
// @Param        sortby  query     string  false  "Field name to sort by (e.g. name, email)"
// @Param        order   query     string  false  "Sort direction: asc or desc (default: asc)"
// @Param        search  query     string  false  "Search term matched against name and email"
// @Param        filter  query     string  false  "Filter by field value (e.g. filter[name]=Herris)"
// @Success      200  {object}  Repository.PaginationResult[Models.User]
// @Failure      500  {object}  ErrorResponse
// @Router       /api/v1/users [get]
func (h *UserHandler) GetAll(c *gin.Context) {
	top, _ := strconv.Atoi(c.DefaultQuery("top", "10"))
	skip, _ := strconv.Atoi(c.DefaultQuery("skip", "0"))

	params := Repository.QueryParams{
		Top:          top,
		Skip:         skip,
		SortBy:       c.Query("sortby"),
		Order:        c.Query("order"),
		Filter:       c.QueryMap("filter"),
		Search:       c.Query("search"),
		SearchFields: []string{"name", "email"},
	}

	result, err := h.repo.FindAll(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

// GetByID godoc
// @Summary      Get user by ID
// @Description  Returns a single user by ID
// @Tags         users
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  Models.User
// @Failure      400  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Router       /api/v1/users/{id} [get]
func (h *UserHandler) GetByID(c *gin.Context) {
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
	c.JSON(http.StatusOK, user)
}

// Create godoc
// @Summary      Create a user
// @Description  Creates a new user with name and email
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        body  body      CreateUserRequest  true  "User data"
// @Success      201   {object}  Models.User
// @Failure      400   {object}  ErrorResponse
// @Failure      500   {object}  ErrorResponse
// @Router       /api/v1/users [post]
func (h *UserHandler) Create(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}
	user := Models.User{Name: req.Name, Email: req.Email}
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
// @Param        id    path      int               true  "User ID"
// @Param        body  body      UpdateUserRequest  true  "Updated user data"
// @Success      200   {object}  Models.User
// @Failure      400   {object}  ErrorResponse
// @Failure      500   {object}  ErrorResponse
// @Router       /api/v1/users/{id} [put]
func (h *UserHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid ID"})
		return
	}
	var req UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}
	user := Models.User{Name: req.Name, Email: req.Email}
	if err := h.repo.Update(uint(id), &user); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}
	updated, _ := h.repo.FindByID(uint(id))
	c.JSON(http.StatusOK, updated)
}

// Delete godoc
// @Summary      Delete a user
// @Description  Deletes a user by ID (soft delete via GORM)
// @Tags         users
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  MessageResponse
// @Failure      400  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /api/v1/users/{id} [delete]
func (h *UserHandler) Delete(c *gin.Context) {
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
