package delivery

import (
	"be/internal/infrastructure/logger"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// HealthHandler struct
type HealthHandler struct {
	logger logger.Logger
}

// NewHealthHandler registers routes
func NewHealthHandler(r *gin.Engine, logger logger.Logger) *HealthHandler {
	return &HealthHandler{logger: logger}
}

// HealthCheck handles the /health endpoint
func (h *HealthHandler) HealthCheck(c *gin.Context) {
	_, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
	defer cancel()

	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}
