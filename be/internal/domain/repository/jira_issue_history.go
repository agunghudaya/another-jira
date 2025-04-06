package repository

import (
	"time"
)

type JiraIssueHistory struct {
	ID       int
	IssueID  int
	Field    string
	Oldvalue string
	NewValue string

	Created time.Time
}
