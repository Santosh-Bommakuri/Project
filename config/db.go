package config

import (
	"log"
	"Project/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func InitDB() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=root dbname=testdb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	log.Println("Database connected successfully")

	
	err = db.AutoMigrate(&models.Product{}, &models.Customer{}, &models.Order{})
	if err != nil {
		return nil, err
	}

	log.Println("Database tables migrated successfully")

	return db, nil
}
