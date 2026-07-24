package V6

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"tutorial/go/Config"
	"tutorial/go/Models"
	"tutorial/go/Repository/V6"

	"github.com/gin-gonic/gin"
)

// InternalApprovalController struct
type InternalApprovalController struct{}

// GetByVendorId handles fetching existing approval data
func (ctrl *InternalApprovalController) GetByVendorId(c *gin.Context) {
	vendorIdStr := c.Param("vendorId")
	vendorId, err := strconv.ParseUint(vendorIdStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Vendor ID"})
		return
	}

	var approval Models.InternalApproval
	if err := V6.GetInternalApprovalByVendorId(uint(vendorId), &approval); err != nil {
		// Auto-fill from Vendor database if InternalApproval not found
		var vendor Models.Vendor
		if errVendor := Config.DB.First(&vendor, vendorId).Error; errVendor == nil {
			approval.VendorId = vendor.ID
			approval.TeamName = vendor.DepartmentDivision
			approval.PicName = vendor.InternalPicFullName
			approval.VendorType = vendor.VendorType
			if vendor.ItemName != nil { approval.ProcuredGoodsServices = *vendor.ItemName }
			if vendor.SelectedCurrency != nil { approval.Currency = *vendor.SelectedCurrency }
			if vendor.EstimatedTotalValue != nil { approval.EstimatedTotalPoValue = *vendor.EstimatedTotalValue }
			approval.BankAccountHolder = vendor.BankAccountHolder
			approval.DiscoverySource = vendor.InformationSource
			approval.ContactPerson = vendor.ContactPersonFullName
			approval.CompanyWebsite = vendor.Website
			if vendor.NpwpDocument != nil { approval.NpwpDocument = *vendor.NpwpDocument }
			if vendor.NibDocument != nil { approval.NibDocument = *vendor.NibDocument }
			if vendor.SppkpDocument != nil { approval.SppkpDocument = *vendor.SppkpDocument }
			if vendor.FrontOfficePhoto != nil { approval.PhysicalLocationPhoto = *vendor.FrontOfficePhoto }
			
			c.JSON(http.StatusOK, gin.H{"data": approval})
			return
		}
		
		c.JSON(http.StatusNotFound, gin.H{"error": "Approval form not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": approval})
}

// Submit handles the multipart form submission (Update-or-Create)
func (ctrl *InternalApprovalController) Submit(c *gin.Context) {
	// 1. Parse Multipart Form (10MB limit)
	err := c.Request.ParseMultipartForm(10 << 20)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse multipart form: " + err.Error()})
		return
	}

	// 2. Map basic fields
	vendorIdStr := c.PostForm("VendorId")
	var vendorId uint
	if vendorIdStr != "" && vendorIdStr != "null" {
		if vid, err := strconv.ParseUint(vendorIdStr, 10, 32); err == nil {
			vendorId = uint(vid)
		}
	}

	estimatedValue, _ := strconv.ParseFloat(c.PostForm("EstimatedTotalPoValue"), 64)
	isCvPersonalAccount := c.PostForm("IsCvPersonalAccount") == "true"

	var approval Models.InternalApproval
	isUpdate := false
	if vendorId != 0 {
		if err := V6.GetInternalApprovalByVendorId(vendorId, &approval); err == nil {
			isUpdate = true
		}
	}

	approval.VendorId = vendorId
	approval.TeamName = c.PostForm("TeamName")
	approval.PicName = c.PostForm("PicName")
	approval.VendorType = c.PostForm("VendorType")
	approval.ProcuredGoodsServices = c.PostForm("ProcuredGoodsServices")
	approval.AgreedPaymentTerm = c.PostForm("AgreedPaymentTerm")
	approval.AgreedPaymentTermCheck = c.PostForm("AgreedPaymentTermCheck") == "true"
	approval.NegotiationHistory = c.PostForm("NegotiationHistory")
	approval.EstimatedTotalPoValue = estimatedValue
	approval.Currency = c.PostForm("Currency")
	approval.BankAccountHolder = c.PostForm("BankAccountHolder")
	approval.BankAccountHolderCheck = c.PostForm("BankAccountHolderCheck") == "true"
	approval.FraudCheckStatus = c.PostForm("FraudCheckStatus")
	approval.FraudCheckStatusCheck = c.PostForm("FraudCheckStatusCheck") == "true"
	approval.DiscoverySource = c.PostForm("DiscoverySource")
	approval.VendorHistory = c.PostForm("VendorHistory")
	approval.NumberOfEmployees = c.PostForm("NumberOfEmployees")
	approval.ContactPerson = c.PostForm("ContactPerson")
	approval.IsCvPersonalAccount = &isCvPersonalAccount
	approval.ThirdPartyValidation = c.PostForm("ThirdPartyValidation")
	approval.AlibabaStatus = c.PostForm("AlibabaStatus")
	approval.CompanyWebsite = c.PostForm("CompanyWebsite")
	approval.NpwpStatus = c.PostForm("NpwpStatus")
	approval.NibStatus = c.PostForm("NibStatus")
	approval.SppkpStatus = c.PostForm("SppkpStatus")
	approval.UsccStatus = c.PostForm("UsccStatus")
	approval.PersonalAccountStatus = c.PostForm("PersonalAccountStatus")
	approval.GoogleStreetViewStatus = c.PostForm("GoogleStreetViewStatus")
	approval.GoogleStreetViewNotes = c.PostForm("GoogleStreetViewNotes")
	approval.AdditionalAttachmentsJson = c.PostForm("AdditionalAttachmentsJson")
	approval.HighlightedFields = c.PostForm("HighlightedFields")

	// 3. Helper to save uploaded files
	saveFile := func(formField string) string {
		file, err := c.FormFile(formField)
		if err != nil {
			return "" // No file uploaded for this field
		}
		ext := filepath.Ext(file.Filename)
		filename := fmt.Sprintf("%s_%d%s", formField, time.Now().Unix(), ext)
		savePath := filepath.Join("uploads", filename)
		if err := c.SaveUploadedFile(file, savePath); err != nil {
			return ""
		}
		return savePath
	}

	// 4. Save attached files or fallback to previous string path
	if newPath := saveFile("NpwpDocument"); newPath != "" { approval.NpwpDocument = newPath } else { approval.NpwpDocument = c.PostForm("NpwpDocument") }
	if newPath := saveFile("NibDocument"); newPath != "" { approval.NibDocument = newPath } else { approval.NibDocument = c.PostForm("NibDocument") }
	if newPath := saveFile("SppkpDocument"); newPath != "" { approval.SppkpDocument = newPath } else { approval.SppkpDocument = c.PostForm("SppkpDocument") }
	if newPath := saveFile("UsccDocument"); newPath != "" { approval.UsccDocument = newPath } else { approval.UsccDocument = c.PostForm("UsccDocument") }
	if newPath := saveFile("PersonalAccountStatement"); newPath != "" { approval.PersonalAccountStatement = newPath } else { approval.PersonalAccountStatement = c.PostForm("PersonalAccountStatement") }
	if newPath := saveFile("GoogleStreetView"); newPath != "" { approval.GoogleStreetView = newPath } else { approval.GoogleStreetView = c.PostForm("GoogleStreetView") }
	if newPath := saveFile("PhysicalLocationPhoto"); newPath != "" { approval.PhysicalLocationPhoto = newPath } else { approval.PhysicalLocationPhoto = c.PostForm("PhysicalLocationPhoto") }
	if newPath := saveFile("AccreditationCertificates"); newPath != "" { approval.AccreditationCertificates = newPath } else { approval.AccreditationCertificates = c.PostForm("AccreditationCertificates") }

	// 5. Save or Update to database
	if isUpdate {
		if err := V6.UpdateInternalApproval(&approval); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update in database"})
			return
		}
	} else {
		if err := V6.CreateInternalApproval(&approval); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save to database"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Internal approval processed successfully",
		"data":    approval,
	})
}
