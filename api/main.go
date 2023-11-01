package main

import (
	"las_api/controllers"
	"las_api/database"
	"las_api/middleware"
	"las_api/models"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	loadDatabase()
	serveApplication()
}

func loadDatabase() {
	database.Connect()
	if err := database.DB.AutoMigrate(
		&models.Admin{},
		&models.Patron{},
		&models.Book{},
		&models.Transaction{},
	); err != nil {
		log.Fatal("Error migrating/updating database schemas")
	}
}

func loadEnv() {
	env := os.Getenv("LAS_ENV")
	if env != "production" {
		if err := godotenv.Load(".env.local"); err != nil {
			log.Fatal("Error loading .env file. Are you sure one exists?")
		}
	}
}

func serveApplication() {
	r := gin.Default()

	api := r.Group("/api")
	api.GET("/healthz", controllers.Health)

	auth := api.Group("/auth")
	auth.POST("/login", controllers.Login)
	auth.POST("/register", middleware.AuthRequired(), controllers.Register)

	inventory := api.Group("/inventory", middleware.AuthRequired())
	inventory.POST("/", controllers.Import)
	// inventory.GET("/")                 // get current inventory and status of each item
	// inventory.GET("/:id")              // get inventory item and its recent transactions
	// inventory.PUT("/:id")              // update inventory item
	// inventory.POST("/:id/transaction") // checkout/issue item
	// inventory.PUT("/:id/transaction")  // mark item returned

	// patrons := api.Group("/patrons")
	// patrons.GET("/")    // get list of existing patrons
	// patrons.POST("/")   // create new patron affiliated with library
	// patrons.GET("/:id") // get specific patron and their recent transactions
	// patrons.PUT("/:id") // update patron information; pay outstanding balance

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Error starting HTTP server")
	}
}
