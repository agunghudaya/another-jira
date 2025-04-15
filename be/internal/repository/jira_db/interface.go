package jiradb

import (
	repository "be/internal/domain/repository"
	"context"
	"time"
)

type JiraDBRepository interface {
	FetchJiraIssue(ctx context.Context, issueKey string) (issue repository.JiraIssueEntity, err error)
	FetchPendingSync(ctx context.Context, jiraID string) ([]repository.SyncHistory, error)
	FetchUserList(ctx context.Context) ([]repository.UserEntity, error)

	InsertJiraIssue(ctx context.Context, issue repository.JiraIssueEntity) error
	InsertJiraIssueHistory(ctx context.Context, history repository.JiraIssueHistoryEntity) error
	InsertSyncHistory(ctx context.Context, jiraID string, status string, recordsSynced int, totalExpected int, errMessage string, startedAt time.Time) error

	UpdateJiraIssue(ctx context.Context, issue repository.JiraIssueEntity) error
}
