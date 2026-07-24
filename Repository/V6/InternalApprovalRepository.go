package V6

import (
	"tutorial/go/Config"
	"tutorial/go/Models"
)

// CreateInternalApproval creates a new internal approval record in the database
func CreateInternalApproval(approval *Models.InternalApproval) error {
	// Auto migrate if not exists
	Config.DB.AutoMigrate(&Models.InternalApproval{})
	
	result := Config.DB.Create(approval)
	return result.Error
}

// GetInternalApprovalByVendorId retrieves an approval record by its VendorId
func GetInternalApprovalByVendorId(vendorId uint, approval *Models.InternalApproval) error {
	return Config.DB.Where("vendor_id = ?", vendorId).First(approval).Error
}

// UpdateInternalApproval updates an existing record
func UpdateInternalApproval(approval *Models.InternalApproval) error {
	return Config.DB.Save(approval).Error
}
