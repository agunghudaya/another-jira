package ucissue

import (
	repository "be/internal/domain/repository"
	"be/internal/infrastructure/config"
	"be/internal/infrastructure/logger"

	jiraDBRp "be/internal/repository/jira_db"
	"context"
)

type UsecaseIssue interface {
	GetAssignedIssueByUserID(ctx context.Context, email string) ([]repository.JiraIssueEntity, error)
}

type usecaseIssue struct {
	cfg    config.Config
	jiraDB jiraDBRp.JiraDBRepository
	log    logger.Logger
}

func NewUsecaseIssue(cfg config.Config, log logger.Logger, jiraDB jiraDBRp.JiraDBRepository) UsecaseIssue {
	return &usecaseIssue{
		cfg:    cfg,
		jiraDB: jiraDB,
		log:    log,
	}
}
