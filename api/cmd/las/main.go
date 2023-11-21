package main

import (
	"fmt"
	"las_api/database"
	"las_api/routes"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println(`
   __      ______  ____       
  /\ \    /\  _  \/\  _'\     
  \ \ \   \ \ \L\ \ \,\L\_\   
   \ \ \  _\ \  __ \/_\__ \   
    \ \ \L\ \ \ \/\ \/\ \L\ \ 
     \ \____/\ \_\ \_\ '\____\
      \/___/  \/_/\/_/\/_____/
      Libraries  Are   Sacred
  `) // Larry 3D ASCII font

	env := os.Getenv("ENVIRONMENT")
	if env != "production" {
		if err := godotenv.Load(".env.local"); err != nil {
			log.Fatal("error loading .env file. Are you sure one exists?")
		}
	}

	s, err := database.NewStore(os.Getenv("DB_URL"))
	if err != nil {
		log.Fatal(err)
	}

	r := routes.NewRouter()
	r.Attach(s)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
