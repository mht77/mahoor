package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
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

	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173", os.Getenv("ALLOW_ORIGIN")}
	r.Use(cors.New(config))
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
		}

		sellController := controllers.NewSellController(services.NewSellService(repositories.NewSellRepository(db)))
		sells := routes.Group("sells")
		{
			sells.GET("/", sellController.GetAllSells)
			sells.GET("/:productId", sellController.GetSellsByProductID)
			sells.POST("/", sellController.CreateSell)
			sells.DELETE("/:id", sellController.DeleteSell)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "healthy"})
	})

	serverErr := r.Run("0.0.0.0:7777")
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
}
