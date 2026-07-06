package Router

import (
	"tutorial/go/Config"
	"tutorial/go/Controller/V1"
	"tutorial/go/Controller/V2"
	"tutorial/go/Controller/V6"
	"tutorial/go/Models"
	"tutorial/go/Repository"
	"tutorial/go/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Enable CORS for all origins (useful for Flutter Web local testing)
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// Expose uploads directory to public
	r.Static("/uploads", "./uploads")

	r.GET("/swagger/v1/swagger.json", docs.ServeOpenAPIJSON)
	r.GET("/swaggerui/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("/swagger/v1/swagger.json")))

	userRepo := Repository.NewGormRepository[Models.User](Config.DB)
	productRepo := Repository.NewGormRepository[Models.Product](Config.DB)
	vendorRepo := Repository.NewGormRepository[Models.Vendor](Config.DB)
	poRepo := Repository.NewGormRepository[Models.PurchaseOrder](Config.DB)

	userController := V1.NewUserController(userRepo)
	productController := V1.NewProductController(productRepo)
	userODataController := V2.NewUserODataController(userRepo)
	productODataController := V2.NewProductODataController(productRepo)
	metadataController := V2.NewMetadataController()

	userV6Controller := V6.NewUserODataController(userRepo)
	vendorV6Controller := V6.NewVendorODataController(vendorRepo)
	poV6Controller := V6.NewPurchaseOrderODataController(poRepo)
	metadataV6Controller := V6.NewMetadataController()

	countryRepo := Repository.NewGormRepository[Models.Country](Config.DB)
	provinceRepo := Repository.NewGormRepository[Models.Province](Config.DB)
	cityRepo := Repository.NewGormRepository[Models.City](Config.DB)
	bankRepo := Repository.NewGormRepository[Models.Bank](Config.DB)
	currencyRepo := Repository.NewGormRepository[Models.Currency](Config.DB)
	industryRepo := Repository.NewGormRepository[Models.IndustryType](Config.DB)
	phoneRepo := Repository.NewGormRepository[Models.PhoneCode](Config.DB)

	masterV6Controller := V6.NewMasterDataController(countryRepo, provinceRepo, cityRepo, bankRepo, currencyRepo, industryRepo, phoneRepo)

	v1 := r.Group("/go")
	{
		users := v1.Group("/users")
		{
			users.GET("", userController.GetAll)
			users.GET("/:id", userController.GetByID)
			users.POST("", userController.Create)
			users.PUT("/:id", userController.Update)
			users.DELETE("/:id", userController.Delete)
		}

		products := v1.Group("/products")
		{
			products.GET("", productController.GetAll)
			products.GET("/:id", productController.GetByID)
			products.POST("", productController.Create)
			products.PUT("/:id", productController.Update)
			products.DELETE("/:id", productController.Delete)
		}
	}

	v2 := r.Group("/go/odata")
	{
		v2.GET("/", metadataController.ServiceDocument)
		v2.GET("/$metadata", metadataController.Metadata)

		users := v2.Group("/users")
		{
			users.GET("", userODataController.GetAll)
			users.GET("/$count", userODataController.Count)
			users.GET("/:id", userODataController.GetByID)
			users.POST("", userController.Create)
			users.PUT("/:id", userController.Update)
			users.DELETE("/:id", userController.Delete)
		}

		products := v2.Group("/products")
		{
			products.GET("", productODataController.GetAll)
			products.GET("/$count", productODataController.Count)
			products.GET("/:id", productODataController.GetByID)
			products.POST("", productController.Create)
			products.PUT("/:id", productController.Update)
			products.DELETE("/:id", productController.Delete)
		}
	}

	v6 := r.Group("/api/v6/odata")
	{
		v6.GET("/", metadataV6Controller.ServiceDocument)
		v6.GET("/$metadata", metadataV6Controller.Metadata)

		usersV6 := v6.Group("/User")
		{
			usersV6.GET("", userV6Controller.GetAll)
			usersV6.GET("/$count", userV6Controller.Count)
			usersV6.GET("/:id", userV6Controller.GetByID)
			usersV6.POST("", userV6Controller.Create)
			usersV6.PUT("/:id", userV6Controller.Update)
			usersV6.DELETE("/:id", userV6Controller.Delete)
		}

		vendorsV6 := v6.Group("/Vendor")
		{
			vendorsV6.GET("", vendorV6Controller.GetAll)
			vendorsV6.GET("/$count", vendorV6Controller.Count)
			vendorsV6.GET("/:id", vendorV6Controller.GetByID)
			vendorsV6.POST("", vendorV6Controller.Create)
			vendorsV6.PUT("/:id", vendorV6Controller.Update)
			vendorsV6.DELETE("/:id", vendorV6Controller.Delete)
		}

		poV6 := v6.Group("/PurchaseOrder")
		{
			poV6.GET("", poV6Controller.GetAll)
			poV6.GET("/$count", poV6Controller.Count)
			poV6.GET("/:id", poV6Controller.GetByID)
			poV6.POST("", poV6Controller.Create)
			poV6.PUT("/:id", poV6Controller.Update)
			poV6.DELETE("/:id", poV6Controller.Delete)
		}

		v6.GET("/Country", masterV6Controller.GetCountries)
		v6.GET("/Province", masterV6Controller.GetProvinces)
		v6.GET("/City", masterV6Controller.GetCities)
		v6.GET("/Bank", masterV6Controller.GetBanks)
		v6.GET("/Currency", masterV6Controller.GetCurrencies)
		v6.GET("/IndustryType", masterV6Controller.GetIndustryTypes)
		v6.GET("/PhoneCode", masterV6Controller.GetPhoneCodes)
	}

	return r
}
