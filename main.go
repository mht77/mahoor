package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"

	"github.com/mht77/mahoor/controllers"
	"github.com/mht77/mahoor/docs"
	"github.com/mht77/mahoor/models"
	"github.com/mht77/mahoor/repositories"
	"github.com/mht77/mahoor/services"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"), 5432)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	migrate(db)

	docs.SwaggerInfo.Title = "Charity Swagger"
	docs.SwaggerInfo.Description = "API for Charity products"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, Charity!")
	})

	routes := r.Group("")
	{
		productController := controllers.NewProductController(services.NewProductService(repositories.NewProductRepository(db)))
		products := routes.Group("products")
		{
			products.POST("/", productController.CreateProduct)
			products.GET("/:id", productController.GetProductByID)
			products.GET("/", productController.GetAllProducts)
			products.PUT("/:id", productController.UpdateProduct)
			products.DELETE("/:id", productController.DeleteProduct)
			products.GET("/:id/qr", productController.GetSellQRCode)
		}

		sellController := controllers.NewSellController(services.NewSellService(repositories.NewSellRepository(db)))
		sells := routes.Group("sells")
		{
			sells.GET("/", sellController.CreateSell)
			sells.GET("/:productId", sellController.GetSellsByProductID)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// TODO
	serverErr := r.RunTLS(":7070", os.Getenv("CERT_FILE"), os.Getenv("KEY_FILE"))
	//serverErr := r.Run(":7070")
	if serverErr != nil {
		panic(err)
	}
}

func migrate(db *gorm.DB) {
	modelsInterfaces := []interface{}{
		&models.Product{},
		&models.Sell{},
	}
	err := db.AutoMigrate(modelsInterfaces...)
	if err != nil {
		panic("failed to migrate database")
	}
	// change the product id sequence to start from 1000
	db.Exec("ALTER SEQUENCE products_id_seq START WITH 1000")
	db.Exec("ALTER SEQUENCE products_id_seq RESTART WITH 1000")
}
