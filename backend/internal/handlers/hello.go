package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewHelloHandler() *HelloHandler {
	return &HelloHandler{}
}

type HelloHandler struct{}

func (h *HelloHandler) Get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "hello from SMS"})
}
