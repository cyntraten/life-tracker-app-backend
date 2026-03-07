package database

import (
	"life-tracker-app-backend/internal/lifetrackerService"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&lifetrackerService.Task{}, &lifetrackerService.Habit{}, &lifetrackerService.Mood{})
}
