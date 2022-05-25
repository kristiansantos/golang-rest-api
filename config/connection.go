package config

import (
	"log"

	"github.com/kristiansantos/go-api-rest/api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Start() {
	connectionString := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(connectionString))

	if err != nil {
		log.Panic("Database connection error")
	}

	DB.AutoMigrate(&models.Planet{})
}
