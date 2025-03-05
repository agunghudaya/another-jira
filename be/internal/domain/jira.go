package domain

import (
	"database/sql"
	"time"
)

type JiraItem struct {
	ID     int
	Title  string
	Status string
}

type SyncHistory struct {
	ID                   int
	JiraID               string
	SyncDate             time.Time
	StartedAt            time.Time
	FinishedAt           sql.NullTime
	Status               string
	ErrorMessage         sql.NullString
	RecordsSynced        int
	TotalExpectedRecords int
	SyncAttempt          int
	CreatedAt            time.Time
}
