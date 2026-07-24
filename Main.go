package main

import (
	"log"

	"tutorial/go/Config"
	"tutorial/go/Router"
	_ "tutorial/go/docs"
)

// @title           Go API V6
// @version         2.0
// @description     REST API with Gin, GORM, and Swagger — v1 uses custom query params, v2 uses OData
// @host            localhost:3020
// @BasePath        /
func main() {
	Config.ConnectDB()
	r := Router.SetupRouter()
	log.Println("Server berjalan di  → http://localhost:3020")
	log.Println("Swagger UI          → http://localhost:3020/swaggerui/index.html")
	log.Println("Swagger JSON        → http://localhost:3020/swaggerui/doc.json")
	if err := r.Run(":3020"); err != nil {
		log.Fatal("Gagal menjalankan server: ", err)
	}
}
