package V6

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"tutorial/go/Models"
	"tutorial/go/Repository"
	"tutorial/go/Utils"

	"github.com/gin-gonic/gin"
)

type VendorODataController struct {
	repo Repository.Repository[Models.Vendor]
}

func NewVendorODataController(repo Repository.Repository[Models.Vendor]) *VendorODataController {
	return &VendorODataController{repo: repo}
}

// GetAll godoc
// @Summary      Get all vendors (OData v6)
// @Description  Returns a paginated vendor collection using OData query options
// @Tags         vendors
// @Produce      json
// @Param        $top      query     int     false  "Number of records to return (default: 10)"
// @Param        $skip     query     int     false  "Number of records to skip (default: 0)"
// @Param        $orderby  query     string  false  "Sort field and direction"
// @Param        $filter   query     string  false  "OData filter"
// @Param        $search   query     string  false  "Search term"
// @Success      200  {object}  ODataPagedResponse[Models.Vendor]
// @Failure      400  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /api/v6/odata/vendors [get]
func (h *VendorODataController) GetAll(c *gin.Context) {
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
		SearchFields: []string{}, 
	}

	result, err := h.repo.FindAll(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, ToODataResponse(c, result, "Vendor"))
}

// GetByID godoc
// @Summary      Get vendor by ID (OData v6)
// @Description  Returns a single vendor with OData context
// @Tags         vendors
// @Produce      json
// @Param        id   path      int  true  "Vendor ID"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Router       /api/v6/odata/vendors/{id} [get]
func (h *VendorODataController) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid ID"})
		return
	}

	vendor, err := h.repo.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Error: "Vendor not found"})
		return
	}

	resp, err := ToODataEntityResponse(vendor, ContextURL(c, "Vendors/$entity"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// Count godoc
// @Summary      Get vendor count (OData v6)
// @Description  Returns the total number of vendors as plain text
// @Tags         vendors
// @Produce      plain
// @Param        $filter  query     string  false  "OData filter expression"
// @Success      200  {string}  string
// @Failure      400  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /api/v6/odata/vendors/$count [get]
func (h *VendorODataController) Count(c *gin.Context) {
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
// @Summary      Create a vendor
// @Description  Creates a new vendor from full JSON payload
// @Tags         vendors
// @Accept       json
// @Produce      json
// @Param        body  body      Models.Vendor  true  "Vendor data"
// @Success      201   {object}  Models.Vendor
// @Failure      400   {object}  ErrorResponse
// @Failure      500   {object}  ErrorResponse
// @Router       /api/v6/odata/vendors [post]
func (h *VendorODataController) Create(c *gin.Context) {
	var vendor Models.Vendor
	if err := c.ShouldBindJSON(&vendor); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	processVendorAttachments(&vendor)

	// Generate ReferenceNumber
	now := time.Now()
	randomStr := fmt.Sprintf("%04d", rand.Intn(10000))
	vendor.ReferenceNumber = fmt.Sprintf("WARUNA-%04d-%02d-%s", now.Year(), now.Month(), randomStr)

	if err := h.repo.Create(&vendor); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, vendor)
}

// Update godoc
// @Summary      Update a vendor
// @Description  Updates a vendor by ID
// @Tags         vendors
// @Accept       json
// @Produce      json
// @Param        id    path      int            true  "Vendor ID"
// @Param        body  body      Models.Vendor  true  "Updated vendor data"
// @Success      200   {object}  Models.Vendor
// @Failure      400   {object}  ErrorResponse
// @Failure      500   {object}  ErrorResponse
// @Router       /api/v6/odata/vendors/{id} [put]
func (h *VendorODataController) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid ID"})
		return
	}
	var vendor Models.Vendor
	if err := c.ShouldBindJSON(&vendor); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	processVendorAttachments(&vendor)

	if err := h.repo.Update(uint(id), &vendor); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}
	updated, _ := h.repo.FindByID(uint(id))
	c.JSON(http.StatusOK, updated)
}

// Delete godoc
// @Summary      Delete a vendor
// @Description  Deletes a vendor by ID
// @Tags         vendors
// @Produce      json
// @Param        id   path      int  true  "Vendor ID"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /api/v6/odata/vendors/{id} [delete]
func (h *VendorODataController) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid ID"})
		return
	}
	if vendor, err := h.repo.FindByID(uint(id)); err == nil {
		vendor.Deleted = "T"
		h.repo.Update(uint(id), &vendor)
	}

	if err := h.repo.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, MessageResponse{Message: "Vendor deleted successfully"})
}

func processVendorAttachments(vendor *Models.Vendor) {
	uploadDir := "uploads/vendors"

	processBase64Field := func(field **string) {
		if *field != nil && len(**field) > 100 { // simple heuristic to distinguish base64 from a URL path
			path, err := Utils.SaveBase64ToFile(**field, uploadDir)
			if err == nil {
				*field = &path
			}
		}
	}

	processBase64Field(&vendor.CustomPaymentTermApproval)
	processBase64Field(&vendor.InformationScreenshot)
	processBase64Field(&vendor.CompanyStampInvoiceHeader)
	processBase64Field(&vendor.FrontOfficePhoto)
	processBase64Field(&vendor.InsideOfficePhoto)
	processBase64Field(&vendor.OfficeVideo)
	processBase64Field(&vendor.BankAccountProof)
	processBase64Field(&vendor.OwnerIdentityCard)
}
