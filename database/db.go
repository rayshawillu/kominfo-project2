package database

import (
	"book-challenge/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "54n_L34rn1n9"
	dbPort   = "5432"
	dbName   = "kominfo-project2"
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbName, dbPort)
	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database :", err)
	}

	db.Debug().AutoMigrate(models.Book{})
}

func GetDB() *gorm.DB {
	return db
}
