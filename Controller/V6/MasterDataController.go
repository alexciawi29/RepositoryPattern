package V6

import (
	"net/http"

	"tutorial/go/Models"
	"tutorial/go/Repository"

	"github.com/gin-gonic/gin"
)

type MasterDataController struct {
	repoCountry  Repository.Repository[Models.Country]
	repoProvince Repository.Repository[Models.Province]
	repoCity     Repository.Repository[Models.City]
	repoBank     Repository.Repository[Models.Bank]
	repoCurrency Repository.Repository[Models.Currency]
	repoIndustry Repository.Repository[Models.IndustryType]
	repoPhone    Repository.Repository[Models.PhoneCode]
}

func NewMasterDataController(
	repoCountry Repository.Repository[Models.Country],
	repoProvince Repository.Repository[Models.Province],
	repoCity Repository.Repository[Models.City],
	repoBank Repository.Repository[Models.Bank],
	repoCurrency Repository.Repository[Models.Currency],
	repoIndustry Repository.Repository[Models.IndustryType],
	repoPhone Repository.Repository[Models.PhoneCode],
) *MasterDataController {
	return &MasterDataController{
		repoCountry:  repoCountry,
		repoProvince: repoProvince,
		repoCity:     repoCity,
		repoBank:     repoBank,
		repoCurrency: repoCurrency,
		repoIndustry: repoIndustry,
		repoPhone:    repoPhone,
	}
}

// GetCountries godoc
func (h *MasterDataController) GetCountries(c *gin.Context) {
	conditions, _ := ParseODataFilter(c.Query("$filter"))
	result, _ := h.repoCountry.FindAll(Repository.QueryParams{Top: 1000, Conditions: conditions})
	c.JSON(http.StatusOK, ToODataResponse(c, result, "MasterCountry"))
}

// GetProvinces godoc
func (h *MasterDataController) GetProvinces(c *gin.Context) {
	conditions, _ := ParseODataFilter(c.Query("$filter"))
	result, _ := h.repoProvince.FindAll(Repository.QueryParams{Top: 1000, Conditions: conditions})
	c.JSON(http.StatusOK, ToODataResponse(c, result, "MasterProvince"))
}

// GetCities godoc
func (h *MasterDataController) GetCities(c *gin.Context) {
	conditions, _ := ParseODataFilter(c.Query("$filter"))
	result, _ := h.repoCity.FindAll(Repository.QueryParams{Top: 1000, Conditions: conditions})
	c.JSON(http.StatusOK, ToODataResponse(c, result, "MasterCity"))
}

// GetBanks godoc
func (h *MasterDataController) GetBanks(c *gin.Context) {
	conditions, _ := ParseODataFilter(c.Query("$filter"))
	result, _ := h.repoBank.FindAll(Repository.QueryParams{Top: 1000, Conditions: conditions})
	c.JSON(http.StatusOK, ToODataResponse(c, result, "MasterBank"))
}

// GetCurrencies godoc
func (h *MasterDataController) GetCurrencies(c *gin.Context) {
	conditions, _ := ParseODataFilter(c.Query("$filter"))
	result, _ := h.repoCurrency.FindAll(Repository.QueryParams{Top: 1000, Conditions: conditions})
	c.JSON(http.StatusOK, ToODataResponse(c, result, "MasterCurrency"))
}

// GetIndustryTypes godoc
func (h *MasterDataController) GetIndustryTypes(c *gin.Context) {
	conditions, _ := ParseODataFilter(c.Query("$filter"))
	result, _ := h.repoIndustry.FindAll(Repository.QueryParams{Top: 1000, Conditions: conditions})
	c.JSON(http.StatusOK, ToODataResponse(c, result, "MasterIndustryType"))
}

// GetPhoneCodes godoc
func (h *MasterDataController) GetPhoneCodes(c *gin.Context) {
	conditions, _ := ParseODataFilter(c.Query("$filter"))
	result, _ := h.repoPhone.FindAll(Repository.QueryParams{Top: 1000, Conditions: conditions})
	c.JSON(http.StatusOK, ToODataResponse(c, result, "MasterPhoneCode"))
}
