package jira_db

import (
	domain "be/internal/domain/repository"
	"be/internal/infrastructure/config"
	"context"
	"database/sql"
	"time"

	"github.com/sirupsen/logrus"
)

type JiraDBRepository interface {
	FetchJiraIssue(ctx context.Context, issueKey string) (issue domain.JiraIssue, err error)
	FetchPendingSync(ctx context.Context, jiraID string) ([]domain.SyncHistory, error)
	FetchUserList(ctx context.Context) ([]domain.User, error)

	InsertJiraIssue(ctx context.Context, issue domain.JiraIssue) error
	InsertSyncHistory(ctx context.Context, jiraID string, status string, recordsSynced int, totalExpected int, errMessage string, startedAt time.Time) error

	MarkSyncAsCompleted(ctx context.Context, syncID int, success bool, recordsSynced int, errMessage *string) error

	UpdateJiraIssue(ctx context.Context, issue domain.JiraIssue) error
}

type jiraDBRepository struct {
	cfg *config.Config
	db  *sql.DB
	log *logrus.Logger
}

func NewJiraDBRepository(cfg *config.Config, log *logrus.Logger, db *sql.DB) JiraDBRepository {
	return &jiraDBRepository{cfg: cfg, db: db, log: log}
}
