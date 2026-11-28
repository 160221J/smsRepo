package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type DemoHandler struct{}

func NewDemoHandler() *DemoHandler { return &DemoHandler{} }

func (h *DemoHandler) Slow(c *gin.Context) {

	ms := 2500

	if raw := c.Query("ms"); raw != "" {
		if v, err := strconv.Atoi(raw); err == nil && v > 0 && v < 30_000 {
			ms = v
		}
	}

	ctx := c.Request.Context()
	timer := time.NewTimer(time.Duration(ms) * time.Millisecond)
	defer timer.Stop()

	select {
	case <-ctx.Done():
		c.JSON(http.StatusGatewayTimeout, gin.H{
			"error":  "canceled or deadline exceeded",
			"detail": ctx.Err().Error(),
		})
	case <-timer.C:
		c.JSON(http.StatusOK, gin.H{"slept_ms": ms, "ok": true})
	}
}
