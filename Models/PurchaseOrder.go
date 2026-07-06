package Models

import "gorm.io/gorm"

type PurchaseOrder struct {
	gorm.Model
	VendorID    uint    `json:"VendorId" gorm:"index"`
	OrderNumber string  `json:"OrderNumber" gorm:"uniqueIndex;not null"`
	TotalAmount float64 `json:"TotalAmount" gorm:"not null"`
	Status      string  `json:"Status" gorm:"default:'Draft'"`
}
