package repository

import (
	"be/internal/domain"
	"database/sql"
)

type SyncRepository interface {
	FetchPendingSync(jiraID string) (*domain.SyncHistory, error)
	FetchUserList() ([]domain.User, error)
	MarkSyncAsCompleted(syncID int, success bool, recordsSynced int, errMessage *string) error
}

type syncRepository struct {
	db *sql.DB
}

func NewSyncRepository(db *sql.DB) SyncRepository {
	return &syncRepository{db: db}
}

func (r *syncRepository) FetchPendingSync(jiraID string) (*domain.SyncHistory, error) {
	query := `
		SELECT id, jira_id, sync_date, started_at, finished_at, status, 
		       error_message, records_synced, total_expected_records, sync_attempt, created_at
		FROM jira_user_sync_history 
		WHERE 
			jira_id = $1 
			AND sync_date = now()
		ORDER BY started_at DESC ;
	`

	var sync domain.SyncHistory
	err := r.db.QueryRow(query, jiraID).Scan(
		&sync.ID, &sync.JiraID, &sync.SyncDate, &sync.StartedAt, &sync.FinishedAt,
		&sync.Status, &sync.ErrorMessage, &sync.RecordsSynced, &sync.TotalExpectedRecords,
		&sync.SyncAttempt, &sync.CreatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &sync, nil
}

func (r *syncRepository) FetchUserList() ([]domain.User, error) {
	query := `select ID, username, jira_user_id, status from users u where status = 'active';`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close() // Ensure rows are closed when function exits

	var users []domain.User

	for rows.Next() {
		var user domain.User
		err := rows.Scan(&user.ID, &user.Username, &user.JiraUserID, &user.Status)
		if err != nil {
			return nil, err // Return immediately on scan error
		}
		users = append(users, user)
	}

	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil // Return slice of users
}

func (r *syncRepository) MarkSyncAsCompleted(syncID int, success bool, recordsSynced int, errMessage *string) error {
	status := "success"
	if !success {
		status = "failed"
	}

	query := `
		UPDATE jira_user_sync_history 
		SET finished_at = NOW(), status = $1, records_synced = $2, error_message = $3 
		WHERE id = $4;
	`
	_, err := r.db.Exec(query, status, recordsSynced, errMessage, syncID)
	return err
}
