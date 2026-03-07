package database

import (
	"life-tracker-app-backend/internal/lifetracker/habits"
	"life-tracker-app-backend/internal/lifetracker/moods"
	"life-tracker-app-backend/internal/lifetracker/tasks"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&tasks.Task{}, &habits.Habit{}, &moods.Mood{})
}
