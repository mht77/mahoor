package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mht77/mahoor/middlewares"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/mht77/mahoor/controllers"
	"github.com/mht77/mahoor/docs"
	"github.com/mht77/mahoor/models"
	"github.com/mht77/mahoor/repositories"
	"github.com/mht77/mahoor/services"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Enter "Bearer {your_token}" to authenticate.
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

	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173", os.Getenv("ALLOW_ORIGIN"), os.Getenv("ALLOW_ORIGIN2")}
	config.AllowHeaders = []string{"Authorization", "Content-Type", "Accept"}
	config.AllowCredentials = true
	r.Use(cors.New(config))
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, Charity!")
	})

	productController := controllers.NewProductController(services.NewProductService(repositories.NewProductRepository(db)))
	products := r.Group("products")
	{
		products.POST("/", middlewares.AuthMiddleware(), productController.CreateProduct)
		products.GET("/:id", productController.GetProductByID)
		products.GET("/", productController.GetAllProducts)
		products.PUT("/:id", middlewares.AuthMiddleware(), productController.UpdateProduct)
		products.DELETE("/:id", middlewares.AuthMiddleware(), productController.DeleteProduct)
	}

	sellController := controllers.NewSellController(services.NewSellService(repositories.NewSellRepository(db)))
	sells := r.Group("sells")
	{
		sells.GET("/", middlewares.AuthMiddleware(), sellController.GetAllSells)
		sells.GET("/:productId", middlewares.AuthMiddleware(), sellController.GetSellsByProductID)
		sells.POST("/", sellController.CreateSell)
		sells.DELETE("/:id", middlewares.AuthMiddleware(), sellController.DeleteSell)
	}

	userController := controllers.NewUserController(services.NewUserService(repositories.NewUserRepository(db)))
	users := r.Group("users")
	{
		users.POST("/", userController.CreateUser)
		users.GET("/", middlewares.AuthMiddleware(), userController.GetAllUsers)
		users.POST("/token", userController.GetToken)
	}

	tikkieController := controllers.NewTikkieController(services.NewTikkieService(repositories.NewTikkieRepository(db)))
	tikkies := r.Group("tikkies", middlewares.AuthMiddleware())
	{
		tikkies.GET("/", tikkieController.GetTikkies)
		tikkies.POST("/", tikkieController.CreateTikkie)
	}

	attendenceController := controllers.NewAttendanceController(services.NewAttendanceService(repositories.NewAttendanceRepository(db)))
	attendences := r.Group("attendances")
	{
		attendences.GET("/", middlewares.AuthMiddleware(), attendenceController.GetAllattendances)
		attendences.DELETE("/:id", middlewares.AuthMiddleware(), attendenceController.DeleteAttendance)
		attendences.POST("/", attendenceController.CreateAttendance)
	}

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Static("/files", "./files")

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "healthy"})
	})

	serverErr := r.Run("0.0.0.0:7777")
	if serverErr != nil {
		panic(serverErr)
	}
}

func migrate(db *gorm.DB) {
	modelsInterfaces := []interface{}{
		&models.Product{},
		&models.Sell{},
		&models.User{},
		&models.Tikkie{},
		&models.Attendance{},
	}
	err := db.AutoMigrate(modelsInterfaces...)
	if err != nil {
		panic("failed to migrate database")
	}
}
