package database

import (
	"errors"
	"fmt"

	"github.com/sstulio/order-service-demo/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase(dns string) (*gorm.DB, error) {
	fmt.Println("connecting to database...")

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		return nil, errors.New("could not connect to database")
	}
	fmt.Println("connected to database!")

	if err := db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`).Error; err != nil {
		return nil, errors.New("could not create uuid extension")
	}

	if err := db.AutoMigrate(&models.Order{}); err != nil {
		return nil, errors.New("could not migrate orders table")
	}

	return db, nil
}
