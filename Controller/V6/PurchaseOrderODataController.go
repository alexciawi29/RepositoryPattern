package V6

import (
	"net/http"
	"strconv"

	"tutorial/go/Models"
	"tutorial/go/Repository"

	"github.com/gin-gonic/gin"
)

type PurchaseOrderODataController struct {
	repo Repository.Repository[Models.PurchaseOrder]
}

func NewPurchaseOrderODataController(repo Repository.Repository[Models.PurchaseOrder]) *PurchaseOrderODataController {
	return &PurchaseOrderODataController{repo: repo}
}

// GetAll godoc
// @Summary      Get all purchase orders (OData v6)
// @Description  Returns a paginated PO collection using OData query options
// @Tags         purchase_orders
// @Produce      json
// @Param        $top      query     int     false  "Number of records to return (default: 10)"
// @Param        $skip     query     int     false  "Number of records to skip (default: 0)"
// @Param        $orderby  query     string  false  "Sort field and direction"
// @Param        $filter   query     string  false  "OData filter"
// @Param        $search   query     string  false  "Search term"
// @Success      200  {object}  ODataPagedResponse[Models.PurchaseOrder]
// @Failure      400  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /api/v6/odata/purchase_orders [get]
func (h *PurchaseOrderODataController) GetAll(c *gin.Context) {
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
		SearchFields: []string{"order_number"}, 
	}

	result, err := h.repo.FindAll(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, ToODataResponse(c, result, "PurchaseOrder"))
}

// GetByID godoc
// @Summary      Get PO by ID (OData v6)
// @Description  Returns a single PO with OData context
// @Tags         purchase_orders
// @Produce      json
// @Param        id   path      int  true  "PO ID"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Router       /api/v6/odata/purchase_orders/{id} [get]
func (h *PurchaseOrderODataController) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid ID"})
		return
	}

	po, err := h.repo.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Error: "PO not found"})
		return
	}

	resp, err := ToODataEntityResponse(po, ContextURL(c, "PurchaseOrders/$entity"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// Count godoc
// @Summary      Get PO count (OData v6)
// @Description  Returns the total number of POs as plain text
// @Tags         purchase_orders
// @Produce      plain
// @Param        $filter  query     string  false  "OData filter expression"
// @Success      200  {string}  string
// @Failure      400  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /api/v6/odata/purchase_orders/$count [get]
func (h *PurchaseOrderODataController) Count(c *gin.Context) {
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
// @Summary      Create a PO
// @Description  Creates a new PO
// @Tags         purchase_orders
// @Accept       json
// @Produce      json
// @Param        body  body      Models.PurchaseOrder  true  "PO data"
// @Success      201   {object}  Models.PurchaseOrder
// @Failure      400   {object}  ErrorResponse
// @Failure      500   {object}  ErrorResponse
// @Router       /api/v6/odata/purchase_orders [post]
func (h *PurchaseOrderODataController) Create(c *gin.Context) {
	var po Models.PurchaseOrder
	if err := c.ShouldBindJSON(&po); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}
	if err := h.repo.Create(&po); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, po)
}

// Update godoc
// @Summary      Update a PO
// @Description  Updates a PO by ID
// @Tags         purchase_orders
// @Accept       json
// @Produce      json
// @Param        id    path      int                   true  "PO ID"
// @Param        body  body      Models.PurchaseOrder  true  "Updated PO data"
// @Success      200   {object}  Models.PurchaseOrder
// @Failure      400   {object}  ErrorResponse
// @Failure      500   {object}  ErrorResponse
// @Router       /api/v6/odata/purchase_orders/{id} [put]
func (h *PurchaseOrderODataController) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid ID"})
		return
	}
	var po Models.PurchaseOrder
	if err := c.ShouldBindJSON(&po); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}
	if err := h.repo.Update(uint(id), &po); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}
	updated, _ := h.repo.FindByID(uint(id))
	c.JSON(http.StatusOK, updated)
}

// Delete godoc
// @Summary      Delete a PO
// @Description  Deletes a PO by ID
// @Tags         purchase_orders
// @Produce      json
// @Param        id   path      int  true  "PO ID"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /api/v6/odata/purchase_orders/{id} [delete]
func (h *PurchaseOrderODataController) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid ID"})
		return
	}
	if err := h.repo.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, MessageResponse{Message: "PO deleted successfully"})
}
