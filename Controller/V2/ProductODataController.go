package V2

import (
	"net/http"
	"strconv"

	"tutorial/go/Models"
	"tutorial/go/Repository"

	"github.com/gin-gonic/gin"
)

type ProductODataController struct {
	repo Repository.Repository[Models.Product]
}

func NewProductODataController(repo Repository.Repository[Models.Product]) *ProductODataController {
	return &ProductODataController{repo: repo}
}

// GetAll godoc
// @Summary      Get all products (OData v4)
// @Description  Returns a paginated product collection using OData v4 query options
// @Tags         products
// @Produce      json
// @Param        $top      query     int     false  "Number of records to return (default: 10)"
// @Param        $skip     query     int     false  "Number of records to skip (default: 0)"
// @Param        $orderby  query     string  false  "Sort field and direction (e.g. name asc, price desc)"
// @Param        $filter   query     string  false  "OData filter (e.g. name eq 'Laptop' and price gt 10000, contains(name,'lap'), startswith(name,'L'))"
// @Param        $search   query     string  false  "Search term matched against product name"
// @Success      200  {object}  ODataPagedResponse[Models.Product]
// @Failure      400  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /api/v2/products [get]
func (h *ProductODataController) GetAll(c *gin.Context) {
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
		SearchFields: []string{"name"},
	}

	result, err := h.repo.FindAll(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, ToODataResponse(c, result, "Products"))
}

// GetByID godoc
// @Summary      Get product by ID (OData v4)
// @Description  Returns a single product with OData v4 context
// @Tags         products
// @Produce      json
// @Param        id   path      int  true  "Product ID"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Router       /api/v2/products/{id} [get]
func (h *ProductODataController) GetByID(c *gin.Context) {
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

	resp, err := ToODataEntityResponse(product, ContextURL(c, "Products/$entity"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// Count godoc
// @Summary      Get product count (OData v4)
// @Description  Returns the total number of products as plain text
// @Tags         products
// @Produce      plain
// @Param        $filter  query     string  false  "OData filter expression"
// @Success      200  {string}  string
// @Failure      400  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /api/v2/products/$count [get]
func (h *ProductODataController) Count(c *gin.Context) {
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
