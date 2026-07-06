package Models

import "gorm.io/gorm"

type Country struct {
	gorm.Model
	Code      string `json:"Code" gorm:"unique;not null"`
	Name      string `json:"Name" gorm:"not null"`
	PhoneCode string `json:"PhoneCode"`
}

func (Country) TableName() string {
	return "Country"
}

type Province struct {
	gorm.Model
	CountryID uint   `json:"CountryId" gorm:"not null"`
	Name      string `json:"Name" gorm:"not null"`
}

func (Province) TableName() string {
	return "Province"
}

type City struct {
	gorm.Model
	ProvinceID uint   `json:"ProvinceId" gorm:"not null"`
	Name       string `json:"Name" gorm:"not null"`
}

func (City) TableName() string {
	return "City"
}

type Bank struct {
	gorm.Model
	Code string `json:"Code" gorm:"unique;not null"`
	Name string `json:"Name" gorm:"not null"`
}

func (Bank) TableName() string {
	return "Bank"
}

type Currency struct {
	gorm.Model
	Code   string `json:"Code" gorm:"unique;not null"`
	Symbol string `json:"Symbol"`
	Name   string `json:"Name" gorm:"not null"`
}

func (Currency) TableName() string {
	return "Currency"
}

type IndustryType struct {
	gorm.Model
	Name string `json:"Name" gorm:"unique;not null"`
}

func (IndustryType) TableName() string {
	return "IndustryType"
}

type PhoneCode struct {
	gorm.Model
	CountryID  uint   `json:"CountryId"`
	DialCode   string `json:"DialCode" gorm:"not null"`
	RegionName string `json:"RegionName" gorm:"not null"`
}

func (PhoneCode) TableName() string {
	return "PhoneCode"
}
