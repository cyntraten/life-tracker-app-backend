package handlers

import (
	"life-tracker-app-backend/internal/lifetracker/habits"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HabitsHandler struct {
	service habits.HabitsService
}

func NewHabitsHandler(s habits.HabitsService) *HabitsHandler {
	return &HabitsHandler{service: s}
}

func (h *HabitsHandler) GetHabits(c *gin.Context) {
	habits, err := h.service.GetAllHabits()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not get Habits"})
		return
	}

	c.JSON(http.StatusOK, habits)
}

func (h *HabitsHandler) PostHabits(c *gin.Context) {
	var habit habits.Habit
	if err := c.ShouldBindJSON(&habit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	habit, err := h.service.CreateHabit(habit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not create calculation"})
		return
	}

	c.JSON(http.StatusCreated, habit)

}

func (h *HabitsHandler) PatchHabits(c *gin.Context) {
	id := c.Param("id")

	var habit habits.Habit
	if err := c.ShouldBindJSON(&habit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	habit, err := h.service.UpdateHabit(id, habit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not update calculation"})
		return
	}

	c.JSON(http.StatusOK, habit)

}

func (h *HabitsHandler) DeleteHabits(c *gin.Context) {
	id := c.Param("id")

	if err := h.service.DeleteHabit(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete calculation"})
		return
	}

	c.Status(http.StatusNoContent)

}
