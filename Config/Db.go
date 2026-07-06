package Config

import (
	"log"

	"tutorial/go/Models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=103.109.172.228 port=9834 user=herris password=Hastalavista007 dbname=alex sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			NoLowerCase:   true, // Forces PascalCase for tables and columns
			SingularTable: true, // Prevents pluralization
		},
	})
	if err != nil {
		log.Fatal("Gagal koneksi database: ", err)
	}

	if err := db.AutoMigrate(
		&Models.User{}, &Models.Product{}, &Models.Vendor{}, &Models.PurchaseOrder{},
		&Models.Country{}, &Models.Province{}, &Models.City{},
		&Models.Bank{}, &Models.Currency{}, &Models.IndustryType{}, &Models.PhoneCode{},
	); err != nil {
		log.Fatal("Gagal AutoMigrate: ", err)
	}

	// Eksekusi injeksi data awal (seeding) untuk Master Data
	SeedMasterData(db)

	DB = db
	log.Println("Koneksi database berhasil!")
}
