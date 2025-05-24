package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"user-service/controllers"
	"user-service/database"
	"user-service/middleware"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	database.Connect()

	r := gin.Default() // statt gin.New() → damit OPTIONS automatisch behandelt wird

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

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
