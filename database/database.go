package database

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDB() {
	str_connection := "user=apisample dbname=apisample password=apisample123456 port=15432 host=localhost sslmode=disable"
	database, err := gorm.Open(postgres.Open(str_connection), &gorm.Config{})
	if err != nil {
		log.Fatal("Error to open connection:", err)
	}
	db = database
	config, err := db.DB()
	if err != nil {
		log.Fatal("Error to use ORM:", err)
	}
	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)
}

func GetDatabase() *gorm.DB {
	return db
}
