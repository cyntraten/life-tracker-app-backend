package handlers

import (
	"life-tracker-app-backend/internal/lifetracker/moods"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MoodsHandler struct {
	service moods.MoodsService
}

func NewMoodsHandler(s moods.MoodsService) *MoodsHandler {
	return &MoodsHandler{service: s}
}

func (h *MoodsHandler) GetMoods(c *gin.Context) {
	Moods, err := h.service.GetAllMoods()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not get Moods"})
		return
	}

	c.JSON(http.StatusOK, Moods)
}

func (h *MoodsHandler) PostMoods(c *gin.Context) {
	var Mood moods.Mood
	if err := c.ShouldBindJSON(&Mood); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	Mood, err := h.service.CreateMood(Mood)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not create calculation"})
		return
	}

	c.JSON(http.StatusCreated, Mood)

}

func (h *MoodsHandler) PatchMoods(c *gin.Context) {
	id := c.Param("id")

	var mood moods.Mood
	if err := c.ShouldBindJSON(&mood); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	mood, err := h.service.UpdateMood(id, mood)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not update calculation"})
		return
	}

	c.JSON(http.StatusOK, mood)

}

func (h *MoodsHandler) DeleteMoods(c *gin.Context) {
	id := c.Param("id")

	if err := h.service.DeleteMood(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete calculation"})
		return
	}

	c.Status(http.StatusNoContent)

}
