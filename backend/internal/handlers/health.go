package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

type HealthHandler struct{}

func (h *HealthHandler) Get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "Backend is running"})
}
