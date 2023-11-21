package api

// import (
// 	"las_api/database"
// 	"las_api/routes"
// 	"log"
// 	"os"

// 	"github.com/gin-gonic/gin"
// 	"github.com/joho/godotenv"
// )

// func main() {
// 	env := os.Getenv("ENVIRONMENT")
// 	if env != "production" {
// 		if err := godotenv.Load(".env.local"); err != nil {
// 			log.Fatal("error loading .env file. Are you sure one exists?")
// 		}
// 	}

// 	s, err := database.NewStore(os.Getenv("DB_URL"))
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	r := gin.Default()
// 	routes.Attach(r, s)

// 	port := os.Getenv("PORT")
// 	if port == "" {
// 		port = "8080"
// 	}
// 	if err := r.Run(":" + port); err != nil {
// 		log.Fatal(err)
// 	}
// }
