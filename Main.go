package main

import (
	"log"

	"tutorial/go/Config"
	"tutorial/go/Router"
	_ "tutorial/go/docs"
)

// @title           Tutorial Go API
// @version         2.0
// @description     REST API with Gin, GORM, and Swagger — v1 uses custom query params, v2 uses OData
// @host            localhost:8080
// @BasePath        /
func main() {
	Config.ConnectDB()
	r := Router.SetupRouter()
	log.Println("Server berjalan di  → http://localhost:8080")
	log.Println("Swagger UI          → http://localhost:8080/swaggerui/index.html")
	log.Println("OpenAPI 3.0.1 JSON  → http://localhost:8080/swagger/v1/swagger.json")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Gagal menjalankan server: ", err)
	}
}
