package db

import (
	"github.com/kekubhai/zyntic/models"
	"gorm.io/gorm"
)

func DBMigrator(db *gorm.DB) error {
	return db.AutoMigrate(&models.Event{}, )
}