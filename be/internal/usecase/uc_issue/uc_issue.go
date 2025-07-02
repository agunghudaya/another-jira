package ucissue

import (
	"context"
	"log"

	repository "be/internal/domain/repository"
)

// UsecaseIssue defines the interface for issue-related use cases
type UsecaseIssue interface {
	SyncIssueByIssueKey(ctx context.Context, issueKey string) error
}

type usecaseIssue struct {
	jiraDB repository.JiraRepository
	log    *log.Logger
}

// NewUsecaseIssue creates a new issue usecase
func NewUsecaseIssue(jiraDB repository.JiraRepository, log *log.Logger) UsecaseIssue {
	return &usecaseIssue{
		jiraDB: jiraDB,
		log:    log,
	}
}

func (uc *usecaseIssue) SyncIssueByIssueKey(ctx context.Context, issueKey string) error {
	// TODO: Implement issue sync logic
	return nil
}
