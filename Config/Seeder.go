package Config

import (
	"tutorial/go/Models"

	"gorm.io/gorm"
)

func SeedMasterData(db *gorm.DB) {
	var count int64
	db.Model(&Models.Country{}).Count(&count)
	if count > 0 {
		return // Data sudah ada, tidak perlu melakukan seeding ulang
	}

	countries := []Models.Country{
		{Code: "ID", Name: "INDONESIA", PhoneCode: "+62"},
		{Code: "SG", Name: "SINGAPORE", PhoneCode: "+65"},
		{Code: "MY", Name: "MALAYSIA", PhoneCode: "+60"},
	}
	db.Create(&countries)

	provinces := []Models.Province{
		{CountryID: countries[0].ID, Name: "DKI JAKARTA"},
		{CountryID: countries[0].ID, Name: "JAWA BARAT"},
		{CountryID: countries[0].ID, Name: "JAWA TENGAH"},
		{CountryID: countries[0].ID, Name: "JAWA TIMUR"},
		{CountryID: countries[0].ID, Name: "BALI"},
		{CountryID: countries[0].ID, Name: "SUMATERA UTARA"},
	}
	db.Create(&provinces)

	cities := []Models.City{
		{ProvinceID: provinces[0].ID, Name: "JAKARTA PUSAT"},
		{ProvinceID: provinces[0].ID, Name: "JAKARTA SELATAN"},
		{ProvinceID: provinces[0].ID, Name: "JAKARTA BARAT"},
		{ProvinceID: provinces[1].ID, Name: "BANDUNG"},
		{ProvinceID: provinces[3].ID, Name: "SURABAYA"},
		{ProvinceID: provinces[5].ID, Name: "MEDAN"},
	}
	db.Create(&cities)

	banks := []Models.Bank{
		{Code: "BCA", Name: "BANK CENTRAL ASIA / BCA"},
		{Code: "MANDIRI", Name: "BANK MANDIRI (PERSERO) / MANDIRI"},
		{Code: "BRI", Name: "BANK RAKYAT INDONESIA (PERSERO) / BRI"},
		{Code: "BNI", Name: "BANK NEGARA INDONESIA (PERSERO) / BNI"},
		{Code: "CIMB", Name: "BANK CIMB NIAGA / CIMB NIAGA"},
		{Code: "DANAMON", Name: "BANK DANAMON INDONESIA / DANAMON"},
		{Code: "BSI", Name: "BANK SYARIAH INDONESIA / BSI"},
		{Code: "BTPN", Name: "BANK BTPN SYARIAH / BTPN SYARIAH"},
	}
	db.Create(&banks)

	currencies := []Models.Currency{
		{Code: "IDR", Symbol: "Rp", Name: "Indonesian Rupiah"},
		{Code: "USD", Symbol: "$", Name: "United States Dollar"},
		{Code: "SGD", Symbol: "S$", Name: "Singapore Dollar"},
		{Code: "EUR", Symbol: "€", Name: "Euro"},
		{Code: "JPY", Symbol: "¥", Name: "Japanese Yen"},
		{Code: "MYR", Symbol: "RM", Name: "Malaysian Ringgit"},
	}
	db.Create(&currencies)

	industries := []Models.IndustryType{
		{Name: "ENGINE SPARE PARTS"},
		{Name: "ANCHORS & CHAINS"},
		{Name: "ROPES & MARINE CABLES"},
		{Name: "PIPES & FITTINGS"},
		{Name: "SHIP REPAIR & MAINTENANCE"},
		{Name: "NAVIGATION EQUIPMENT"},
		{Name: "COMMUNICATION SYSTEMS"},
		{Name: "ELECTRICAL SYSTEMS"},
		{Name: "MARINE LOGISTICS"},
		{Name: "SHIPPING/VESSEL TRANSPORTATION"},
		{Name: "OTHER (MANUAL INPUT)"},
	}
	db.Create(&industries)

	phoneCodes := []Models.PhoneCode{
		{CountryID: countries[0].ID, DialCode: "+62", RegionName: "(INDONESIA)"},
		{CountryID: countries[1].ID, DialCode: "+65", RegionName: "(SINGAPORE)"},
		{CountryID: countries[2].ID, DialCode: "+60", RegionName: "(MALAYSIA)"},
		{CountryID: countries[0].ID, DialCode: "(021)", RegionName: "DKI Jakarta - Kota Jakarta, Kota Tangerang"},
		{CountryID: countries[0].ID, DialCode: "(022)", RegionName: "Jawa Barat - Kota Bandung, Kota Cimahi"},
		{CountryID: countries[0].ID, DialCode: "(031)", RegionName: "Jawa Timur - Kota Surabaya, Kabupaten Sidoarjo"},
		{CountryID: countries[0].ID, DialCode: "(061)", RegionName: "Sumatera Utara - Kota Medan, Kota Binjai"},
	}
	db.Create(&phoneCodes)
}
