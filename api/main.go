package main

import (
	"las_api/controllers"
	"las_api/database"
	"las_api/models"
	"log"

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
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func serveApplication() {
	r := gin.Default()

	auth := r.Group("/auth")
	auth.POST("/login", controllers.Login)
	auth.POST("/register", controllers.Register)

	r.Run()
}
