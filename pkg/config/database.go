package config

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb(models ...interface{}) {
	db, err := gorm.Open(sqlite.Open("shop.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database. \n", err.Error())
		os.Exit(2)
	}

	log.Println("Connected to the database successfully!")
	// db.Logger = logger.Default.LogMode(logger.Info)

	//Add Migrations
	log.Println("Running Migrations")
	db.AutoMigrate(models...)

	Database = DbInstance{Db: db}
}
