package entity

import (
	"be/internal/domain/errors"
	"time"
)

// JiraIssue represents the core domain entity for a Jira issue
type JiraIssue struct {
	ID           string
	Key          string
	Summary      string
	Description  string
	Status       IssueStatus
	Type         IssueType
	Priority     Priority
	Project      Project
	Assignee     User
	Reporter     User
	TimeTracking TimeTracking
	Created      time.Time
	Updated      time.Time
	DueDate      *time.Time
}

// IssueStatus represents the status of a Jira issue
type IssueStatus struct {
	Name        string
	Description string
	Category    StatusCategory
}

// StatusCategory represents the category of a status
type StatusCategory struct {
	Key  string
	Name string
}

// IssueType represents the type of a Jira issue
type IssueType struct {
	Name        string
	Description string
}

// Priority represents the priority of a Jira issue
type Priority struct {
	Name string
}

// Project represents a Jira project
type Project struct {
	ID   int
	Key  string
	Name string
}

// User represents a Jira user
type User struct {
	Email       string
	DisplayName string
}

// TimeTracking represents time tracking information
type TimeTracking struct {
	OriginalEstimate          float64
	CurrentEstimate           float64
	AggregateEstimate         float64
	AggregateOriginalEstimate float64
}

// Validate implements domain validation
func (i *JiraIssue) Validate() error {
	if i.Key == "" {
		return ErrInvalidIssueKey
	}
	if i.Summary == "" {
		return ErrInvalidSummary
	}
	return nil
}

// Domain errors
var (
	ErrInvalidIssueKey = errors.NewDomainError("invalid_issue_key", "Issue key cannot be empty")
	ErrInvalidSummary  = errors.NewDomainError("invalid_summary", "Summary cannot be empty")
)
