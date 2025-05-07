package database

import (
	"log"
	"os"
	"todo-service/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB-Verbindung fehlgeschlagen:", err)
	}

	err = db.AutoMigrate(&models.Todo{}, &models.Project{}, &models.Category{})
	if err != nil {
		log.Fatal("Migration fehlgeschlagen:", err)
	}

	DB = db
}
