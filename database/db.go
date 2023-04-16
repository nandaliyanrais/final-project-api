package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"mygram-api/models/domain"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "/"
	dbPort   = 5432
	dbName   = "mygram-api"
	db       *gorm.DB
	err      error
)

func StartDB() *gorm.DB {

	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", host, user, password, dbName, dbPort)
	dsn := config
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database :", err)
	}

	if err := db.AutoMigrate(&domain.User{}, &domain.Photo{}, &domain.Comment{}, &domain.SocialMedia{}); err != nil {
		log.Fatal(err.Error())
	}

	return db
}
