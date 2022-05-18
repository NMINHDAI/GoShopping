package database

import (
	"log"

	"GoShopping/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dbURL := "postgres://pg:pass@localhost:5432/crud"
	connection, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	DB = connection
	connection.AutoMigrate(models.User{})
}
