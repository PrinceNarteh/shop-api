package config

import (
	"log"
	"os"

	"shop_api/modules/order"
	"shop_api/modules/product"
	"shop_api/modules/user"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	db, err := gorm.Open(sqlite.Open("shop.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database. \n", err.Error())
		os.Exit(2)
	}

	log.Println("Connected to the database successfully!")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migrations")

	// TODO: Add Migrations
	db.AutoMigrate(&user.User{}, &product.Product{}, &order.Order{})

	Database = DbInstance{Db: db}
}
