package repository

import (
	"context"
	"be/internal/domain/entity"
	"time"
)

// JiraRepository defines the interface for Jira data operations
type JiraRepository interface {
	// SaveIssues saves multiple Jira issues
	SaveIssues(ctx context.Context, issues []entity.JiraIssue) error
	
	// GetIssues retrieves Jira issues based on filter
	GetIssues(ctx context.Context, filter JiraFilter) ([]entity.JiraIssue, error)
	
	// GetIssueByKey retrieves a single Jira issue by its key
	GetIssueByKey(ctx context.Context, key string) (*entity.JiraIssue, error)
	
	// UpdateIssue updates an existing Jira issue
	UpdateIssue(ctx context.Context, issue entity.JiraIssue) error
	
	// DeleteIssue deletes a Jira issue
	DeleteIssue(ctx context.Context, key string) error
}

// JiraFilter represents the filter criteria for querying Jira issues
type JiraFilter struct {
	ProjectKey    string
	Status        string
	AssigneeEmail string
	CreatedAfter  *time.Time
	CreatedBefore *time.Time
	UpdatedAfter  *time.Time
	UpdatedBefore *time.Time
	Limit         int
	Offset        int
} 