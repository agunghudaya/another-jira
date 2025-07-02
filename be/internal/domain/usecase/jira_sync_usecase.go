package usecase

import (
	"context"
	"be/internal/domain/entity"
	"be/internal/domain/repository"
)

// JiraSyncUseCase defines the interface for Jira synchronization operations
type JiraSyncUseCase interface {
	// SyncIssues synchronizes issues from Jira to the local database
	SyncIssues(ctx context.Context) error
	
	// GetIssues retrieves issues based on filter criteria
	GetIssues(ctx context.Context, filter repository.JiraFilter) ([]entity.JiraIssue, error)
	
	// GetIssueByKey retrieves a single issue by its key
	GetIssueByKey(ctx context.Context, key string) (*entity.JiraIssue, error)
}

// jiraSyncUseCase implements JiraSyncUseCase
type jiraSyncUseCase struct {
	jiraRepo repository.JiraRepository
}

// NewJiraSyncUseCase creates a new instance of JiraSyncUseCase
func NewJiraSyncUseCase(jiraRepo repository.JiraRepository) JiraSyncUseCase {
	return &jiraSyncUseCase{
		jiraRepo: jiraRepo,
	}
}

// SyncIssues implements the synchronization logic
func (uc *jiraSyncUseCase) SyncIssues(ctx context.Context) error {
	// TODO: Implement synchronization logic
	// 1. Fetch issues from Jira API
	// 2. Transform to domain entities
	// 3. Save to local database
	return nil
}

// GetIssues implements the issue retrieval logic
func (uc *jiraSyncUseCase) GetIssues(ctx context.Context, filter repository.JiraFilter) ([]entity.JiraIssue, error) {
	return uc.jiraRepo.GetIssues(ctx, filter)
}

// GetIssueByKey implements the single issue retrieval logic
func (uc *jiraSyncUseCase) GetIssueByKey(ctx context.Context, key string) (*entity.JiraIssue, error) {
	return uc.jiraRepo.GetIssueByKey(ctx, key)
} 