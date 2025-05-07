package main

import (
	"log"
	"os"
	"user-service/controllers"
	"user-service/database"
	"user-service/middleware"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	database.Connect()

	r := gin.New()
	r.Use(middleware.Logger())

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	r.GET("/me", controllers.Me)
	r.GET("/users/:id", controllers.GetUserByID)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8001"
	}
	log.Fatal(r.Run(":" + port))
}
