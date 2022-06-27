package config

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDb(models ...interface{}) {
	dns := "host=localhost port=5432 user=postgres password=root dbname=shop sslmode=disable TimeZone=Africa/Accra"
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database. \n", err.Error())
		os.Exit(2)
	}

	log.Println("Connected to the database successfully!")
	// db.Logger = logger.Default.LogMode(logger.Info)

	//Add Migrations
	log.Println("Running Migrations")
	db.AutoMigrate(models...)

	DB = db
}
