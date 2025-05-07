package main

import (
	"log"
	"os"
	"todo-service/database"
	"todo-service/middleware"
	"todo-service/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	database.Connect()

	r := gin.New()
	r.Use(middleware.Logger())
	r.Use(middleware.AuthMiddleware())

	routes.SetupRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8002"
	}
	log.Fatal(r.Run(":" + port))
}
