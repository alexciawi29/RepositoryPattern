package Models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string `json:"Name"  gorm:"not null"`
	Email string `json:"Email" gorm:"unique;not null"`
}
