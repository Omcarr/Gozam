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
	if err := db.AutoMigrate(
		&models.User{},
		&models.Song{},
	); err != nil {
		return nil, err // Return migration error
	}

	return db, nil
}

func InsertSong(song *models.Song, postgresClient *gorm.DB) error {
	result := postgresClient.Create(&song)
	if result.Error != nil {
		return errors.New("failed to insertsong in song table")
	}
	return nil
}

func IsSonginDB(ytID string, postgresClient *gorm.DB) (bool, error) {
	var count int64

	result := postgresClient.Model(&models.Song{}).Where("yt_id = ?", ytID).Count(&count)
	if result.Error != nil {
		return false, errors.New("postgres query failed")
	}

	return count > 0, nil
}

func GetSongByID(songID uint32, postgresClient *gorm.DB) (models.Song, bool, error) {
	var song models.Song

	result := postgresClient.Where("id = ?", songID).First(&song)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return song, false, nil
		}
		return song, false, result.Error
	}

	return song, true, nil
}
