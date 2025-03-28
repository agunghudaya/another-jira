package repository

import (
	"be/internal/domain"
	"be/internal/infrastructure/config"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/sirupsen/logrus"
)

type SyncRepository interface {
	FetchJiraTasksWithFilter(ctx context.Context, jiraUserID string, cfg *config.Config) (domain.JiraResponse, error)
	FetchPendingSync(ctx context.Context, jiraID string) (*domain.SyncHistory, error)
	FetchUserList(ctx context.Context) ([]domain.User, error)
	InsertSyncHistory(ctx context.Context, jiraID string, status string, recordsSynced int, totalExpected int, errMessage string, startedAt time.Time) error
	MarkSyncAsCompleted(ctx context.Context, syncID int, success bool, recordsSynced int, errMessage *string) error
}

type syncRepository struct {
	cfg *config.Config
	db  *sql.DB
	log *logrus.Logger
}

func NewSyncRepository(cfg *config.Config, log *logrus.Logger, db *sql.DB) SyncRepository {
	return &syncRepository{cfg: cfg, db: db, log: log}
}

func (r *syncRepository) FetchPendingSync(ctx context.Context, jiraID string) (*domain.SyncHistory, error) {
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

func (r *syncRepository) FetchUserList(ctx context.Context) ([]domain.User, error) {

	select {
	case <-ctx.Done():
		r.log.Warn("SaveOrder operation was canceled!")
		return nil, ctx.Err()
	default:
		// Continue processing
	}

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

func (r *syncRepository) MarkSyncAsCompleted(ctx context.Context, syncID int, success bool, recordsSynced int, errMessage *string) error {
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

func (r *syncRepository) FetchJiraTasksWithFilter(ctx context.Context, jiraUserID string, cfg *config.Config) (domain.JiraResponse, error) {

	jiraData := domain.JiraResponse{}
	startAt := 0

	jql := fmt.Sprintf(`
	assignee = %s 
	and status not in (CANCELLED, OPEN) 
	and issuetype != Epic 
	and (created >= "2025-01-01" or resolved >= "2025-01-01") 
	order by status DESC, created DESC`, jiraUserID)

	client := &http.Client{}
	reqURL := fmt.Sprintf("%s%s?jql=%s&&maxResults=50&startAt=%d", r.cfg.GetString("jira.baseurl"), r.cfg.GetString("jira.searchurl"), url.QueryEscape(jql), startAt)

	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return jiraData, err
	}

	req.SetBasicAuth(cfg.GetString("jira_username"), cfg.GetString("jira_token"))
	req.Header.Add("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return jiraData, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return jiraData, fmt.Errorf("failed to fetch data: %s", resp.Status)
	}

	var tmpJiraData *domain.JiraResponse

	err = json.NewDecoder(resp.Body).Decode(&tmpJiraData)
	if err != nil {
		return jiraData, err
	}

	jiraData.Issues = append(jiraData.Issues, tmpJiraData.Issues...)

	return jiraData, nil
}

func (r *syncRepository) InsertSyncHistory(ctx context.Context, jiraID string, status string, recordsSynced int, totalExpected int, errMessage string, startedAt time.Time) error {
	now := time.Now()
	syncDate := now.Format("2006-01-02")

	query := `
		INSERT INTO public.jira_user_sync_history 
			(jira_id, sync_date, started_at, finished_at, status, error_message, records_synced, total_expected_records, sync_attempt, created_at)
		VALUES 
			($1, $2, $3, $4, $5, $6, $7, $8, 1, NOW())
		ON CONFLICT (jira_id, sync_date, sync_attempt) DO UPDATE 
		SET 
			finished_at = EXCLUDED.finished_at, 
			status = EXCLUDED.status,
			error_message = EXCLUDED.error_message,
			records_synced = EXCLUDED.records_synced,
			total_expected_records = EXCLUDED.total_expected_records
		RETURNING id;
	`

	var id int
	err := r.db.QueryRow(query, jiraID, syncDate, startedAt, now, status, errMessage, recordsSynced, totalExpected).Scan(&id)
	if err != nil {
		return err
	}

	fmt.Printf("Inserted/Updated sync history ID: %d\n", id)
	return nil
}
