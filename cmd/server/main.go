package main

import (
	"life-tracker-app-backend/internal/database"
	"life-tracker-app-backend/internal/handlers"
	"life-tracker-app-backend/internal/lifetracker/habits"
	"life-tracker-app-backend/internal/lifetracker/moods"
	"life-tracker-app-backend/internal/lifetracker/tasks"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	router := gin.Default()
	router.SetTrustedProxies(nil)

	tasksRepo := tasks.NewTasksRepository(db)
	tasksService := tasks.NewTasksService(tasksRepo)
	tasksHandler := handlers.NewTasksHandler(tasksService)

	moodsRepo := moods.NewMoodsRepository(db)
	moodsService := moods.NewMoodsService(moodsRepo)
	moodsHandler := handlers.NewMoodsHandler(moodsService)

	habitsRepo := habits.NewHabitsRepository(db)
	habitsService := habits.NewHabitsService(habitsRepo)
	habitsHandler := handlers.NewHabitsHandler(habitsService)

	router.GET("/tasks", tasksHandler.GetTasks)
	router.POST("/tasks", tasksHandler.PostTasks)
	router.DELETE("/tasks/:id", tasksHandler.DeleteTasks)
	router.PATCH("/tasks/:id", tasksHandler.PatchTasks)

	router.GET("/moods", moodsHandler.GetMoods)
	router.POST("/moods", moodsHandler.PostMoods)
	router.DELETE("/moods/:id", moodsHandler.DeleteMoods)
	router.PATCH("/moods/:id", moodsHandler.PatchMoods)

	router.GET("/habits", habitsHandler.GetHabits)
	router.POST("/habits", habitsHandler.PostHabits)
	router.DELETE("/habits/:id", habitsHandler.DeleteHabits)
	router.PATCH("/habits/:id", habitsHandler.PatchHabits)
	router.Run()

}
