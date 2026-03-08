package handlers

import (
	"life-tracker-app-backend/internal/lifetracker/tasks"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TasksHandler struct {
	service tasks.TasksService
}

func NewTasksHandler(s tasks.TasksService) *TasksHandler {
	return &TasksHandler{service: s}
}

func (h *TasksHandler) GetTasks(c *gin.Context) {
	tasks, err := h.service.GetAllTasks()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not get tasks"})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func (h *TasksHandler) PostTasks(c *gin.Context) {
	var task tasks.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	task, err := h.service.CreateTask(task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not create calculation"})
		return
	}

	c.JSON(http.StatusCreated, task)

}

func (h *TasksHandler) PatchTasks(c *gin.Context) {
	id := c.Param("id")

	var task tasks.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	task, err := h.service.UpdateTask(id, task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not update calculation"})
		return
	}

	c.JSON(http.StatusOK, task)

}

func (h *TasksHandler) DeleteTasks(c *gin.Context) {
	id := c.Param("id")

	if err := h.service.DeleteTask(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete calculation"})
		return
	}

	c.Status(http.StatusNoContent)

}
