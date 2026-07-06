package V6

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ServiceDocumentEntry struct {
	Name string `json:"Name"`
	Kind string `json:"Kind"`
	URL  string `json:"Url"`
}

type ServiceDocumentResponse struct {
	Context string                 `json:"@odata.context"`
	Value   []ServiceDocumentEntry `json:"Value"`
}

const edmx = `<?xml version="1.0" encoding="utf-8"?>
<edmx:Edmx Version="4.0" xmlns:edmx="http://docs.oasis-open.org/odata/ns/edmx">
  <edmx:DataServices>
    <Schema Namespace="TutorialGoApi" xmlns="http://docs.oasis-open.org/odata/ns/edm">
      <EntityType Name="Vendor">
        <Key><PropertyRef Name="ID"/></Key>
        <Property Name="ID"        Type="Edm.Int64"          Nullable="false"/>
        <Property Name="CreatedAt" Type="Edm.DateTimeOffset" Nullable="false"/>
        <Property Name="UpdatedAt" Type="Edm.DateTimeOffset" Nullable="false"/>
        <Property Name="DeletedAt" Type="Edm.DateTimeOffset"/>
        <Property Name="RegistrationPurpose" Type="Edm.String"/>
        <Property Name="VendorType" Type="Edm.String"/>
        <Property Name="Detail" Type="Edm.String"/>
        <Property Name="OrderType" Type="Edm.String"/>
        <Property Name="ShipmentMethod" Type="Edm.String"/>
        <Property Name="ReimburseTo" Type="Edm.String"/>
        <Property Name="DepartmentDivision" Type="Edm.String"/>
        <Property Name="DeliveryType" Type="Edm.String"/>
        <Property Name="WorkOrderRemark" Type="Edm.String"/>
        <Property Name="DiverWork" Type="Edm.String"/>
        <Property Name="ExistingVendorName" Type="Edm.String"/>
        <Property Name="DisableExistingVendor" Type="Edm.String"/>
        <Property Name="ChangeReason" Type="Edm.String"/>
        <Property Name="PaymentTermType" Type="Edm.String"/>
        <Property Name="CreditTermDay" Type="Edm.String"/>
        <Property Name="AdvanceTermType" Type="Edm.String"/>
        <Property Name="CustomPaymentTerm" Type="Edm.String"/>
        <Property Name="TaxType" Type="Edm.String"/>
        <Property Name="PpnType" Type="Edm.String"/>
        <Property Name="DefaultDiscount" Type="Edm.String"/>
        <Property Name="Step2DeliveryType" Type="Edm.String"/>
        <Property Name="DeliveryMethod" Type="Edm.String"/>
        <Property Name="DeliveryCityArea" Type="Edm.String"/>
        <Property Name="Country" Type="Edm.String"/>
        <Property Name="StateCity" Type="Edm.String"/>
        <Property Name="InformationSource" Type="Edm.String"/>
        <Property Name="InternalPicFullName" Type="Edm.String"/>
        <Property Name="InternalPicJobTitle" Type="Edm.String"/>
        <Property Name="InternalPicWhatsappCode" Type="Edm.String"/>
        <Property Name="InternalPicWhatsappNumber" Type="Edm.String"/>
        <Property Name="RecommenderVendorName" Type="Edm.String"/>
        <Property Name="RecommenderPicFullName" Type="Edm.String"/>
        <Property Name="RecommenderPicJobTitle" Type="Edm.String"/>
        <Property Name="RecommenderPicWhatsappCode" Type="Edm.String"/>
        <Property Name="RecommenderPicWhatsappNumber" Type="Edm.String"/>
        <Property Name="ContactPersonFullName" Type="Edm.String"/>
        <Property Name="ContactPersonJobTitle" Type="Edm.String"/>
        <Property Name="ContactPersonWhatsappCode" Type="Edm.String"/>
        <Property Name="ContactPersonWhatsappNumber" Type="Edm.String"/>
        <Property Name="SubmissionType" Type="Edm.String"/>
        <Property Name="BusinessType" Type="Edm.String"/>
        <Property Name="CompanyName" Type="Edm.String"/>
        <Property Name="YearEstablished" Type="Edm.String"/>
        <Property Name="Website" Type="Edm.String"/>
        <Property Name="MainContactName" Type="Edm.String"/>
        <Property Name="MainContactTitle" Type="Edm.String"/>
        <Property Name="MainContactEmail" Type="Edm.String"/>
        <Property Name="MainContactPhoneCode" Type="Edm.String"/>
        <Property Name="MainContactPhone" Type="Edm.String"/>
        <Property Name="InquirySameAsMain" Type="Edm.Boolean"/>
        <Property Name="InquiryContactName" Type="Edm.String"/>
        <Property Name="InquiryContactTitle" Type="Edm.String"/>
        <Property Name="InquiryContactEmail" Type="Edm.String"/>
        <Property Name="InquiryContactPhoneCode" Type="Edm.String"/>
        <Property Name="InquiryContactPhone" Type="Edm.String"/>
        <Property Name="PoSameAsMain" Type="Edm.Boolean"/>
        <Property Name="PoContactName" Type="Edm.String"/>
        <Property Name="PoContactTitle" Type="Edm.String"/>
        <Property Name="PoContactEmail" Type="Edm.String"/>
        <Property Name="PoContactPhoneCode" Type="Edm.String"/>
        <Property Name="PoContactPhone" Type="Edm.String"/>
        <Property Name="PaymentSameAsMain" Type="Edm.Boolean"/>
        <Property Name="PaymentContactName" Type="Edm.String"/>
        <Property Name="PaymentContactTitle" Type="Edm.String"/>
        <Property Name="PaymentContactEmail" Type="Edm.String"/>
        <Property Name="PaymentContactPhoneCode" Type="Edm.String"/>
        <Property Name="PaymentContactPhone" Type="Edm.String"/>
        <Property Name="StreetAddress" Type="Edm.String"/>
        <Property Name="Province" Type="Edm.String"/>
        <Property Name="CityArea" Type="Edm.String"/>
        <Property Name="PostalCode" Type="Edm.String"/>
        <Property Name="FullAddress" Type="Edm.String"/>
        <Property Name="MapUrl" Type="Edm.String"/>
        <Property Name="OfficePhoneCode" Type="Edm.String"/>
        <Property Name="OfficePhone" Type="Edm.String"/>
        <Property Name="IndustryType" Type="Edm.String"/>
        <Property Name="CustomIndustryType" Type="Edm.String"/>
        <Property Name="CompanySize" Type="Edm.String"/>
        <Property Name="VendorBackground" Type="Edm.String"/>
        <Property Name="Certification" Type="Edm.String"/>
        <Property Name="HasCompanyBankAccount" Type="Edm.Boolean"/>
        <Property Name="BankName" Type="Edm.String"/>
        <Property Name="CustomBankName" Type="Edm.String"/>
        <Property Name="BankAccountNumber" Type="Edm.String"/>
        <Property Name="BankAccountHolder" Type="Edm.String"/>
        <Property Name="BankCurrency" Type="Edm.String"/>
        <Property Name="IsIntegrityPactAccepted" Type="Edm.Boolean"/>
        <Property Name="IntegrityPactSignerName" Type="Edm.String"/>
        <Property Name="IntegrityPactSignerRole" Type="Edm.String"/>
        <Property Name="CustomPaymentTermApproval" Type="Edm.String"/>
        <Property Name="InformationScreenshot" Type="Edm.String"/>
        <Property Name="CompanyStampInvoiceHeader" Type="Edm.String"/>
        <Property Name="FrontOfficePhoto" Type="Edm.String"/>
        <Property Name="InsideOfficePhoto" Type="Edm.String"/>
        <Property Name="OfficeVideo" Type="Edm.String"/>
        <Property Name="BankAccountProof" Type="Edm.String"/>
        <Property Name="OwnerIdentityCard" Type="Edm.String"/>
      </EntityType>
      <EntityType Name="PurchaseOrder">
        <Key><PropertyRef Name="ID"/></Key>
        <Property Name="ID"          Type="Edm.Int64"          Nullable="false"/>
        <Property Name="VendorID"    Type="Edm.Int64"          Nullable="false"/>
        <Property Name="OrderNumber" Type="Edm.String"         Nullable="false"/>
        <Property Name="TotalAmount" Type="Edm.Decimal"        Nullable="false"/>
        <Property Name="Status"      Type="Edm.String"         Nullable="false"/>
        <Property Name="CreatedAt"   Type="Edm.DateTimeOffset" Nullable="false"/>
        <Property Name="UpdatedAt"   Type="Edm.DateTimeOffset" Nullable="false"/>
        <Property Name="DeletedAt"   Type="Edm.DateTimeOffset"/>
      </EntityType>
      <EntityType Name="User">
        <Key><PropertyRef Name="ID"/></Key>
        <Property Name="ID"        Type="Edm.Int64"          Nullable="false"/>
        <Property Name="Name"      Type="Edm.String"         Nullable="false"/>
        <Property Name="Email"     Type="Edm.String"         Nullable="false"/>
        <Property Name="CreatedAt" Type="Edm.DateTimeOffset" Nullable="false"/>
        <Property Name="UpdatedAt" Type="Edm.DateTimeOffset" Nullable="false"/>
        <Property Name="DeletedAt" Type="Edm.DateTimeOffset"/>
      </EntityType>
      <EntityType Name="Country">
        <Key><PropertyRef Name="ID"/></Key>
        <Property Name="ID"        Type="Edm.Int64"  Nullable="false"/>
        <Property Name="Code"      Type="Edm.String" Nullable="false"/>
        <Property Name="Name"      Type="Edm.String" Nullable="false"/>
        <Property Name="PhoneCode" Type="Edm.String"/>
      </EntityType>
      <EntityType Name="Province">
        <Key><PropertyRef Name="ID"/></Key>
        <Property Name="ID"        Type="Edm.Int64"  Nullable="false"/>
        <Property Name="CountryID" Type="Edm.Int64"  Nullable="false"/>
        <Property Name="Name"      Type="Edm.String" Nullable="false"/>
      </EntityType>
      <EntityType Name="City">
        <Key><PropertyRef Name="ID"/></Key>
        <Property Name="ID"         Type="Edm.Int64"  Nullable="false"/>
        <Property Name="ProvinceID" Type="Edm.Int64"  Nullable="false"/>
        <Property Name="Name"       Type="Edm.String" Nullable="false"/>
      </EntityType>
      <EntityType Name="Bank">
        <Key><PropertyRef Name="ID"/></Key>
        <Property Name="ID"   Type="Edm.Int64"  Nullable="false"/>
        <Property Name="Code" Type="Edm.String" Nullable="false"/>
        <Property Name="Name" Type="Edm.String" Nullable="false"/>
      </EntityType>
      <EntityType Name="Currency">
        <Key><PropertyRef Name="ID"/></Key>
        <Property Name="ID"     Type="Edm.Int64"  Nullable="false"/>
        <Property Name="Code"   Type="Edm.String" Nullable="false"/>
        <Property Name="Symbol" Type="Edm.String"/>
        <Property Name="Name"   Type="Edm.String" Nullable="false"/>
      </EntityType>
      <EntityType Name="IndustryType">
        <Key><PropertyRef Name="ID"/></Key>
        <Property Name="ID"   Type="Edm.Int64"  Nullable="false"/>
        <Property Name="Name" Type="Edm.String" Nullable="false"/>
      </EntityType>
      <EntityType Name="PhoneCode">
        <Key><PropertyRef Name="ID"/></Key>
        <Property Name="ID"         Type="Edm.Int64"  Nullable="false"/>
        <Property Name="CountryID"  Type="Edm.Int64"/>
        <Property Name="DialCode"   Type="Edm.String" Nullable="false"/>
        <Property Name="RegionName" Type="Edm.String" Nullable="false"/>
      </EntityType>
      <EntityContainer Name="DefaultContainer">
        <EntitySet Name="Vendor"        EntityType="TutorialGoApi.Vendor"/>
        <EntitySet Name="PurchaseOrder" EntityType="TutorialGoApi.PurchaseOrder"/>
        <EntitySet Name="User"          EntityType="TutorialGoApi.User"/>
        <EntitySet Name="Country"     EntityType="TutorialGoApi.Country"/>
        <EntitySet Name="Province"     EntityType="TutorialGoApi.Province"/>
        <EntitySet Name="City"        EntityType="TutorialGoApi.City"/>
        <EntitySet Name="Bank"         EntityType="TutorialGoApi.Bank"/>
        <EntitySet Name="Currency"    EntityType="TutorialGoApi.Currency"/>
        <EntitySet Name="IndustryType" EntityType="TutorialGoApi.IndustryType"/>
        <EntitySet Name="PhoneCode"    EntityType="TutorialGoApi.PhoneCode"/>
      </EntityContainer>
    </Schema>
  </edmx:DataServices>
</edmx:Edmx>`

type MetadataController struct{}

func NewMetadataController() *MetadataController {
	return &MetadataController{}
}

// ServiceDocument godoc
// @Summary      OData v6 service document
// @Description  Lists all available entity sets in the v6 API
// @Tags         odata
// @Produce      json
// @Success      200  {object}  ServiceDocumentResponse
// @Router       /api/v6/odata/ [get]
func (h *MetadataController) ServiceDocument(c *gin.Context) {
	root := serviceRoot(c)
	c.JSON(http.StatusOK, ServiceDocumentResponse{
		Context: fmt.Sprintf("%s/$metadata", root),
		Value: []ServiceDocumentEntry{
			{Name: "Vendor", Kind: "EntitySet", URL: "Vendor"},
			{Name: "PurchaseOrder", Kind: "EntitySet", URL: "PurchaseOrder"},
			{Name: "User", Kind: "EntitySet", URL: "User"},
			{Name: "MasterCountry", Kind: "EntitySet", URL: "Country"},
			{Name: "MasterProvince", Kind: "EntitySet", URL: "Province"},
			{Name: "MasterCity", Kind: "EntitySet", URL: "City"},
			{Name: "MasterBank", Kind: "EntitySet", URL: "Bank"},
			{Name: "MasterCurrency", Kind: "EntitySet", URL: "Currency"},
			{Name: "MasterIndustryType", Kind: "EntitySet", URL: "IndustryType"},
			{Name: "MasterPhoneCode", Kind: "EntitySet", URL: "PhoneCode"},
		},
	})
}

// Metadata godoc
// @Summary      OData v6 metadata document
// @Description  Returns the EDMX metadata document describing all entity types and their properties
// @Tags         odata
// @Produce      xml
// @Success      200  {string}  string
// @Router       /api/v6/odata/$metadata [get]
func (h *MetadataController) Metadata(c *gin.Context) {
	c.Data(http.StatusOK, "application/xml;charset=utf-8", []byte(edmx))
}
