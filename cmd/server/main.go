package main

import (
	"life-tracker-app-backend/internal/app"
	"life-tracker-app-backend/internal/database"
	"life-tracker-app-backend/internal/routes"
	"log"

	"github.com/gin-contrib/cors"
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
	router.Use(cors.Default())

	handlersContainer := app.InitializeHandlers(db)

	routeHandlers := &routes.Handlers{
		TasksHandler:  handlersContainer.TasksHandler,
		MoodsHandler:  handlersContainer.MoodsHandler,
		HabitsHandler: handlersContainer.HabitsHandler,
	}

	routes.SetupRoutes(router, routeHandlers)

	if err := router.Run(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}

}
