package database

import (
	"errors"
	"fmt"

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

	return db, nil
}
