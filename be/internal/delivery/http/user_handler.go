package delivery

import (
	"context"
	"net/http"
	"time"

	"be/internal/infrastructure/logger"

	"github.com/gin-gonic/gin"

	ucIssue "be/internal/usecase/uc_issue"
	ucUser "be/internal/usecase/uc_user"
)

// UserHandler struct
type UserHandler struct {
	logger  logger.Logger
	ucUser  ucUser.UsecaseUser
	ucIssue ucIssue.UsecaseIssue
}

// NewUserHandler registers routes
func NewUserHandler(r *gin.Engine, logger logger.Logger, ucUser ucUser.UsecaseUser, ucIssue ucIssue.UsecaseIssue) *UserHandler {
	return &UserHandler{logger: logger, ucUser: ucUser, ucIssue: ucIssue}
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	_, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
	defer cancel()

	resp, err := h.ucUser.GetAllUsers(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"data": resp})

	return

}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	_, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
	defer cancel()

	// Get the parameter from the URL
	userID := c.Param("id")

	resp, err := h.ucUser.GetUserByID(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": resp})
}

func (h *UserHandler) GetAssignedIssuesByUserID(c *gin.Context) {
	_, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
	defer cancel()

	// Get the parameter from the URL
	userID := c.Param("id")

	resp, err := h.ucIssue.GetAssignedIssueByUserID(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": resp})
}
