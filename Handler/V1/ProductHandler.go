package V1

import (
	"net/http"
	"strconv"

	"tutorial/go/Models"
	"tutorial/go/Repository"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	repo Repository.Repository[Models.Product]
}

func NewProductHandler(repo Repository.Repository[Models.Product]) *ProductHandler {
	return &ProductHandler{repo: repo}
}

type CreateProductRequest struct {
	Name  string  `json:"Name"  binding:"required" example:"Laptop"`
	Price float64 `json:"Price" binding:"required" example:"15000000"`
}

type UpdateProductRequest struct {
	Name  string  `json:"Name"  example:"Laptop Pro"`
	Price float64 `json:"Price" example:"18000000"`
}

// GetAll godoc
// @Summary      Get all products
// @Description  Returns a paginated list of products with optional sorting, filtering, and search
// @Tags         products
// @Produce      json
// @Param        top     query     int     false  "Number of records to return (default: 10)"
// @Param        skip    query     int     false  "Number of records to skip (default: 0)"
// @Param        sortby  query     string  false  "Field name to sort by (e.g. name, price)"
// @Param        order   query     string  false  "Sort direction: asc or desc (default: asc)"
// @Param        search  query     string  false  "Search term matched against product name"
// @Param        filter  query     string  false  "Filter by field value (e.g. filter[name]=Laptop)"
// @Success      200  {object}  Repository.PaginationResult[Models.Product]
// @Failure      500  {object}  ErrorResponse
// @Router       /api/v1/products [get]
func (h *ProductHandler) GetAll(c *gin.Context) {
	top, _ := strconv.Atoi(c.DefaultQuery("top", "10"))
	skip, _ := strconv.Atoi(c.DefaultQuery("skip", "0"))

	params := Repository.QueryParams{
		Top:          top,
		Skip:         skip,
		SortBy:       c.Query("sortby"),
		Order:        c.Query("order"),
		Filter:       c.QueryMap("filter"),
		Search:       c.Query("search"),
		SearchFields: []string{"name"},
	}

	result, err := h.repo.FindAll(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

// GetByID godoc
// @Summary      Get product by ID
// @Description  Returns a single product by ID
// @Tags         products
// @Produce      json
// @Param        id   path      int  true  "Product ID"
// @Success      200  {object}  Models.Product
// @Failure      400  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Router       /api/v1/products/{id} [get]
func (h *ProductHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid ID"})
		return
	}
	product, err := h.repo.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Error: "Product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

// Create godoc
// @Summary      Create a product
// @Description  Creates a new product with name and price
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        body  body      CreateProductRequest  true  "Product data"
// @Success      201   {object}  Models.Product
// @Failure      400   {object}  ErrorResponse
// @Failure      500   {object}  ErrorResponse
// @Router       /api/v1/products [post]
func (h *ProductHandler) Create(c *gin.Context) {
	var req CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}
	product := Models.Product{Name: req.Name, Price: req.Price}
	if err := h.repo.Create(&product); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, product)
}

// Update godoc
// @Summary      Update a product
// @Description  Updates a product by ID
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id    path      int                   true  "Product ID"
// @Param        body  body      UpdateProductRequest  true  "Updated product data"
// @Success      200   {object}  Models.Product
// @Failure      400   {object}  ErrorResponse
// @Failure      500   {object}  ErrorResponse
// @Router       /api/v1/products/{id} [put]
func (h *ProductHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid ID"})
		return
	}
	var req UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}
	product := Models.Product{Name: req.Name, Price: req.Price}
	if err := h.repo.Update(uint(id), &product); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}
	updated, _ := h.repo.FindByID(uint(id))
	c.JSON(http.StatusOK, updated)
}

// Delete godoc
// @Summary      Delete a product
// @Description  Deletes a product by ID (soft delete via GORM)
// @Tags         products
// @Produce      json
// @Param        id   path      int  true  "Product ID"
// @Success      200  {object}  MessageResponse
// @Failure      400  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /api/v1/products/{id} [delete]
func (h *ProductHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid ID"})
		return
	}
	if err := h.repo.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, MessageResponse{Message: "Product deleted successfully"})
}
