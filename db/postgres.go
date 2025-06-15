package db

import (
	"errors"
	"gozam/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresClient() (*gorm.DB, error) {
	dbURL := os.Getenv("postgresURL")
	if dbURL == "" {
		return nil, errors.New("postgresURL environment variable is not set")
	}

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Automigrate models
	if err := db.AutoMigrate(&models.User{}); err != nil {
		return nil, err // Return migration error
	}

	return db, nil
}
