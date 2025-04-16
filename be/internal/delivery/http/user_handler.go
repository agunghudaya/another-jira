package delivery

import (
	"context"
	"net/http"
	"time"

	"be/internal/infrastructure/logger"

	"github.com/gin-gonic/gin"

	ucUser "be/internal/usecase/uc_user"
)

// UserHandler struct
type UserHandler struct {
	logger logger.Logger
	ucUser ucUser.UsecaseUser
}

// NewUserHandler registers routes
func NewUserHandler(r *gin.Engine, logger logger.Logger, ucUser ucUser.UsecaseUser) *UserHandler {
	return &UserHandler{logger: logger, ucUser: ucUser}
}

// HealthCheck handles the /health endpoint
func (h *UserHandler) GetUsers(c *gin.Context) {
	_, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
	defer cancel()

	resp, err := h.ucUser.GetAllJiraUsers(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"data": resp})

	return

}
