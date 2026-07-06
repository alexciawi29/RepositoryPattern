package V2

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ServiceDocumentEntry struct {
	Name string `json:"name"`
	Kind string `json:"kind"`
	URL  string `json:"url"`
}

type ServiceDocumentResponse struct {
	Context string                 `json:"@odata.context"`
	Value   []ServiceDocumentEntry `json:"value"`
}

const edmx = `<?xml version="1.0" encoding="utf-8"?>
<edmx:Edmx Version="4.0" xmlns:edmx="http://docs.oasis-open.org/odata/ns/edmx">
  <edmx:DataServices>
    <Schema Namespace="TutorialGoApi" xmlns="http://docs.oasis-open.org/odata/ns/edm">
      <EntityType Name="Product">
        <Key><PropertyRef Name="ID"/></Key>
        <Property Name="ID"        Type="Edm.Int64"          Nullable="false"/>
        <Property Name="Name"      Type="Edm.String"         Nullable="false"/>
        <Property Name="Price"     Type="Edm.Decimal"        Nullable="false"/>
        <Property Name="CreatedAt" Type="Edm.DateTimeOffset" Nullable="false"/>
        <Property Name="UpdatedAt" Type="Edm.DateTimeOffset" Nullable="false"/>
        <Property Name="DeletedAt" Type="Edm.DateTimeOffset"/>
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
      <EntityContainer Name="DefaultContainer">
        <EntitySet Name="Products" EntityType="TutorialGoApi.Product"/>
        <EntitySet Name="Users"    EntityType="TutorialGoApi.User"/>
      </EntityContainer>
    </Schema>
  </edmx:DataServices>
</edmx:Edmx>`

type MetadataController struct{}

func NewMetadataController() *MetadataController {
	return &MetadataController{}
}

// ServiceDocument godoc
// @Summary      OData v4 service document
// @Description  Lists all available entity sets in the v2 API
// @Tags         odata
// @Produce      json
// @Success      200  {object}  ServiceDocumentResponse
// @Router       /api/v2/ [get]
func (h *MetadataController) ServiceDocument(c *gin.Context) {
	root := serviceRoot(c)
	c.JSON(http.StatusOK, ServiceDocumentResponse{
		Context: fmt.Sprintf("%s/$metadata", root),
		Value: []ServiceDocumentEntry{
			{Name: "Products", Kind: "EntitySet", URL: "Products"},
			{Name: "Users", Kind: "EntitySet", URL: "Users"},
		},
	})
}

// Metadata godoc
// @Summary      OData v4 metadata document
// @Description  Returns the EDMX metadata document describing all entity types and their properties
// @Tags         odata
// @Produce      xml
// @Success      200  {string}  string
// @Router       /api/v2/$metadata [get]
func (h *MetadataController) Metadata(c *gin.Context) {
	c.Data(http.StatusOK, "application/xml;charset=utf-8", []byte(edmx))
}
