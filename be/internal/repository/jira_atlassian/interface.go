package jira_atlassian

import (
	"be/internal/domain/repository"
	"be/internal/infrastructure/config"
	repository "be/internal/repository/jira_db/entity"
	"context"
	"database/sql"

	"github.com/sirupsen/logrus"
)

type JiraAtlassianRepository interface {
	FetchJiraTasksWithFilter(ctx context.Context, jiraUserID string, cfg *config.Config) (repository.JiraIssueResponse, error)
	FetchJiraIssueHistories(ctx context.Context, jiraIssueKey string, cfg *config.Config) (repository.JiraIssueHistory, error)
}

type jiraAtlassianRepository struct {
	cfg *config.Config
	db  *sql.DB
	log *logrus.Logger
}

func NewJiraAtlassianRepository(cfg *config.Config, log *logrus.Logger, db *sql.DB) JiraAtlassianRepository {
	return &jiraAtlassianRepository{cfg: cfg, db: db, log: log}
}
