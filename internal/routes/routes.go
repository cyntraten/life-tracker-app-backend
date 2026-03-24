package routes

import (
	"life-tracker-app-backend/internal/handlers"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	TasksHandler  *handlers.TasksHandler
	HabitsHandler *handlers.HabitsHandler
	MoodsHandler  *handlers.MoodsHandler
}

func SetupRoutes(router *gin.Engine, handlers *Handlers) {
	taskRoutes := router.Group("/api/tasks")
	{
		taskRoutes.GET("", handlers.TasksHandler.GetTasks)
		taskRoutes.POST("", handlers.TasksHandler.PostTasks)
		taskRoutes.DELETE("/tasks/:id", handlers.TasksHandler.DeleteTasks)
		taskRoutes.PATCH("/tasks/:id", handlers.TasksHandler.PatchTasks)
	}

	moodRoutes := router.Group("/moods")
	{
		moodRoutes.GET("", handlers.MoodsHandler.GetMoods)
		moodRoutes.POST("", handlers.MoodsHandler.PostMoods)
		moodRoutes.DELETE("/:id", handlers.MoodsHandler.DeleteMoods)
		moodRoutes.PATCH("/:id", handlers.MoodsHandler.PatchMoods)
	}

	habitRoutes := router.Group("/habits")
	{
		habitRoutes.GET("", handlers.HabitsHandler.GetHabits)
		habitRoutes.POST("", handlers.HabitsHandler.PostHabits)
		habitRoutes.DELETE("/:id", handlers.HabitsHandler.DeleteHabits)
		habitRoutes.PATCH("/:id", handlers.HabitsHandler.PatchHabits)
	}
}
