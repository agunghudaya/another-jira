package delivery

import (
	"context"
	"net/http"
	"time"

	"be/internal/domain/errors"
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
	handler := &UserHandler{logger: logger, ucUser: ucUser}

	// Register routes
	userGroup := r.Group("/api/v1/users")
	{
		userGroup.GET("", handler.GetUsers)
	}

	return handler
}

// GetUsers handles the GET /users endpoint
func (h *UserHandler) GetUsers(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
	defer cancel()

	users, err := h.ucUser.GetAllJiraUsers(ctx)
	if err != nil {
		h.logger.Error("Failed to get users", "error", err)
		response := errors.NewErrorResponse(err)
		status := errors.GetHTTPStatus(err)
		_ = response.WriteJSON(c.Writer, status)
		return
	}

	response := errors.NewSuccessResponse(users, nil)
	_ = response.WriteJSON(c.Writer, http.StatusOK)
}
