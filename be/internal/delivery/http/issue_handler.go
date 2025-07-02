package delivery

import (
	"context"
	"net/http"
	"time"

	"be/internal/domain/errors"
	"be/internal/infrastructure/logger"

	"github.com/gin-gonic/gin"

	ucIssue "be/internal/usecase/uc_issue"
)

// IssueHandler struct
type IssueHandler struct {
	logger  logger.Logger
	ucIssue ucIssue.UsecaseIssue
}

// NewIssueHandler registers routes
func NewIssueHandler(r *gin.Engine, logger logger.Logger, ucIssue ucIssue.UsecaseIssue) *IssueHandler {
	handler := &IssueHandler{logger: logger, ucIssue: ucIssue}

	// Register routes
	issueGroup := r.Group("/api/v1/issues")
	{
		issueGroup.POST("/sync/:issueKey", handler.SyncIssueByIssueKey)
	}

	return handler
}

func (h *IssueHandler) SyncIssueByIssueKey(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
	defer cancel()

	issueKey := c.Param("issueKey")
	if issueKey == "" {
		response := errors.NewErrorResponse(errors.InvalidInputError)
		_ = response.WriteJSON(c.Writer, http.StatusBadRequest)
		return
	}

	err := h.ucIssue.SyncIssueByIssueKey(ctx, issueKey)
	if err != nil {
		h.logger.Error("Failed to sync issue", "error", err, "issueKey", issueKey)
		response := errors.NewErrorResponse(err)
		status := errors.GetHTTPStatus(err)
		_ = response.WriteJSON(c.Writer, status)
		return
	}

	response := errors.NewSuccessResponse("Issue sync initiated", nil)
	_ = response.WriteJSON(c.Writer, http.StatusOK)
}
