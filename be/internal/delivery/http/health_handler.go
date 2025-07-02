package delivery

import (
	"be/internal/infrastructure/logger"
	"context"
	"net/http"
	"time"

	"be/internal/infrastructure/db"

	"github.com/gin-gonic/gin"
)

// HealthHandler struct
type HealthHandler struct {
	logger logger.Logger
	db     db.DB
}

// NewHealthHandler registers routes
func NewHealthHandler(r *gin.Engine, logger logger.Logger, db db.DB) *HealthHandler {
	return &HealthHandler{logger: logger, db: db}
}

// HealthCheck handles the /health endpoint
func (h *HealthHandler) HealthCheck(c *gin.Context) {
	// ctx, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
	_, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
	defer cancel()

	dbErr := h.db.Ping()
	workerHealthy := false // TODO: Implement worker heartbeat check

	if dbErr != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"status": "fail", "db": "unhealthy", "worker": workerHealthy})
		return
	}

	// Placeholder: check worker heartbeat from DB
	// if worker is healthy, set workerHealthy = true

	if !workerHealthy {
		c.JSON(http.StatusServiceUnavailable, gin.H{"status": "fail", "db": "healthy", "worker": "unhealthy"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "OK", "db": "healthy", "worker": "healthy"})
}
