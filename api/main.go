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
	database.DB.AutoMigrate(&models.Admin{})
	database.DB.AutoMigrate(&models.Patron{})
	database.DB.AutoMigrate(&models.Book{})
	database.DB.AutoMigrate(&models.Transaction{})
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

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
