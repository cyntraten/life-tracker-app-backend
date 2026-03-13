package app

import (
	"life-tracker-app-backend/internal/handlers"
	"life-tracker-app-backend/internal/lifetracker/habits"
	"life-tracker-app-backend/internal/lifetracker/moods"
	"life-tracker-app-backend/internal/lifetracker/tasks"

	"gorm.io/gorm"
)

type HandlersContainer struct {
	TasksHandler  *handlers.TasksHandler
	MoodsHandler  *handlers.MoodsHandler
	HabitsHandler *handlers.HabitsHandler
}

func InitializeHandlers(db *gorm.DB) *HandlersContainer {
	tasksRepo := tasks.NewTasksRepository(db)
	tasksService := tasks.NewTasksService(tasksRepo)
	tasksHandler := handlers.NewTasksHandler(tasksService)

	moodsRepo := moods.NewMoodsRepository(db)
	moodsService := moods.NewMoodsService(moodsRepo)
	moodsHandler := handlers.NewMoodsHandler(moodsService)

	habitsRepo := habits.NewHabitsRepository(db)
	habitsService := habits.NewHabitsService(habitsRepo)
	habitsHandler := handlers.NewHabitsHandler(habitsService)

	return &HandlersContainer{
		TasksHandler:  tasksHandler,
		MoodsHandler:  moodsHandler,
		HabitsHandler: habitsHandler,
	}
}
