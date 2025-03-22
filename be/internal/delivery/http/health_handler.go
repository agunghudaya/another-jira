package delivery

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// HealthHandler struct
type HealthHandler struct {
	logger *logrus.Logger
}

// NewHealthHandler registers routes
func NewHealthHandler(r *gin.Engine, logger *logrus.Logger) *HealthHandler {
	return &HealthHandler{logger: logger}
}

// HealthCheck handles the /health endpoint
func (h *HealthHandler) HealthCheck(c *gin.Context) {
	_, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
	defer cancel()

	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}
