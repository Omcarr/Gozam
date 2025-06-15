package models

import (
	"gorm.io/gorm"
)

type Song struct {
	gorm.Model        // provides standard fields
	ID         int    `json:"ID" gorm:"primaryKey"`
	Title      string `json:"title"`
	Artist     string `json:"artist"`
	YtID       string `json:"ytID" gorm:"unique"`
}
