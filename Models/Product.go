package Models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name  string  `json:"Name"  gorm:"not null"`
	Price float64 `json:"Price" gorm:"not null"`
}
