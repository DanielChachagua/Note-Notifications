package database

import (
	"note_notifications/internal/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConectDB(uri string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(uri), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&models.Note{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func CloseDB(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}