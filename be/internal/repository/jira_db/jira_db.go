package jira_db

import (
	domain "be/internal/domain/repository"
	"context"
	"database/sql"
	"time"
)

func (r *jiraDBRepository) FetchJiraIssue(ctx context.Context, issueKey string) (domain.JiraIssue, error) {
	query := `
        SELECT id, "key", "self", url, assignee_email, assignee_name, reporter_email, reporter_name, 
               creator_email, creator_name, summary, description, created, updated, duedate, 
               statuscategorychangedate, issue_type_name, issue_type_desc, project_id, project_key, 
               project_name, priority_name, status_name, status_desc, status_category_name, 
               status_category_key, time_estimate, time_original_estimate, aggregate_time_estimate, 
               aggregate_time_original_estimate
        FROM public.jira_issues
        WHERE 
            "key" = $1
        LIMIT 1;
    `

	var issue domain.JiraIssue
	err := r.db.QueryRowContext(ctx, query, issueKey).Scan(
		&issue.ID, &issue.Key, &issue.Self, &issue.URL,
		&issue.AssigneeEmail, &issue.AssigneeName,
		&issue.ReporterEmail, &issue.ReporterName,
		&issue.CreatorEmail, &issue.CreatorName,
		&issue.Summary, &issue.Description,
		&issue.Created, &issue.Updated, &issue.DueDate, &issue.StatusCategoryChange,
		&issue.IssueTypeName, &issue.IssueTypeDescription,
		&issue.ProjectID, &issue.ProjectKey, &issue.ProjectName,
		&issue.PriorityName, &issue.StatusName, &issue.StatusDescription,
		&issue.StatusCategoryName, &issue.StatusCategoryKey,
		&issue.TimeEstimate, &issue.TimeOriginalEstimate,
		&issue.AggregateTimeEstimate, &issue.AggregateTimeOriginalEstimate,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			// Handle case where no rows are returned
			r.log.Warnf("No Jira issue found for key: %s", issueKey)
			return domain.JiraIssue{}, nil
		}
		// Handle other errors
		r.log.Errorf("Error fetching Jira issue for key %s: %v", issueKey, err)
		return domain.JiraIssue{}, err
	}

	return issue, nil
}

func (r *jiraDBRepository) FetchPendingSync(ctx context.Context, jiraID string) ([]domain.SyncHistory, error) {
	query := `
		SELECT id, jira_id, sync_date, started_at, finished_at, status, 
		       error_message, records_synced, total_expected_records, sync_attempt, created_at
		FROM jira_user_sync_history 
		WHERE 
			jira_id = $1 
			AND sync_date = CURRENT_DATE
		ORDER BY started_at DESC ;
	`

	rows, err := r.db.Query(query, jiraID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var syncs []domain.SyncHistory

	for rows.Next() {
		var sync domain.SyncHistory
		err := rows.Scan(
			&sync.ID, &sync.JiraID, &sync.SyncDate, &sync.StartedAt, &sync.FinishedAt,
			&sync.Status, &sync.ErrorMessage, &sync.RecordsSynced, &sync.TotalExpectedRecords,
			&sync.SyncAttempt, &sync.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		syncs = append(syncs, sync)
	}

	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return syncs, nil
}

func (r *jiraDBRepository) FetchUserList(ctx context.Context) ([]domain.User, error) {

	select {
	case <-ctx.Done():
		r.log.Warn("FetchUserList operation was canceled!")
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

func (r *jiraDBRepository) InsertJiraIssue(ctx context.Context, issue domain.JiraIssue) error {
	query := `
    INSERT INTO public.jira_issues (
        "key", "self", url, 
        assignee_email, assignee_name, 
        reporter_email, reporter_name, 
        creator_email, creator_name, 
        summary, description, 
        created, updated, duedate, statuscategorychangedate, 
        issue_type_name, issue_type_desc, 
        project_id, project_key, project_name, 
        priority_name, 
        status_name, status_desc, status_category_name, status_category_key,
        aggregate_time_estimate,
        aggregate_time_original_estimate,
        time_estimate,
        time_original_estimate
    ) VALUES (
        $1, $2, $3, 
        $4, $5, 
        $6, $7, 
        $8, $9, 
        $10, $11, 
        $12, $13, $14, $15, 
        $16, $17, $18, 
        $19, $20, $21, 
        $22, $23, 
        $24, $25, $26, $27, $28, $29
    ) RETURNING id;`

	var id int
	err := r.db.QueryRowContext(ctx, query,
		issue.Key, issue.Self, issue.URL,
		issue.AssigneeEmail, issue.AssigneeName,
		issue.ReporterEmail, issue.ReporterName,
		issue.CreatorEmail, issue.CreatorName,
		issue.Summary, issue.Description,
		issue.Created, issue.Updated, issue.DueDate, issue.StatusCategoryChange,
		issue.IssueTypeName, issue.IssueTypeDescription,
		issue.ProjectID, issue.ProjectKey, issue.ProjectName,
		issue.PriorityName,
		issue.StatusName, issue.StatusDescription, issue.StatusCategoryName, issue.StatusCategoryKey,
		issue.AggregateTimeEstimate, issue.AggregateTimeOriginalEstimate,
		issue.TimeEstimate, issue.TimeOriginalEstimate,
	).Scan(&id)

	if err != nil {
		return err
	}

	r.log.Infof("Inserted Jira issue with ID: %d\n", id)
	return nil
}

func (r *jiraDBRepository) InsertSyncHistory(ctx context.Context, jiraID string, status string, recordsSynced int, totalExpected int, errMessage string, startedAt time.Time) error {
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

	r.log.Infof("Inserted/Updated sync history ID: %d\n", id)
	return nil
}

func (r *jiraDBRepository) MarkSyncAsCompleted(ctx context.Context, syncID int, success bool, recordsSynced int, errMessage *string) error {
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

func (r *jiraDBRepository) UpdateJiraIssue(ctx context.Context, issue domain.JiraIssue) error {
	query := `
        UPDATE public.jira_issues
        SET 
            assignee_email = $1,
            assignee_name = $2,
            reporter_email = $3,
            reporter_name = $4,
            creator_email = $5,
            creator_name = $6,
            summary = $7,
            description = $8,
            updated = $9,
            status_name = $10,
            status_desc = $11,
            status_category_name = $12,
            status_category_key = $13
        WHERE 
            "key" = $14;
    `

	_, err := r.db.ExecContext(ctx, query,
		issue.AssigneeEmail, issue.AssigneeName,
		issue.ReporterEmail, issue.ReporterName,
		issue.CreatorEmail, issue.CreatorName,
		issue.Summary, issue.Description,
		issue.Updated,
		issue.StatusName, issue.StatusDescription,
		issue.StatusCategoryName, issue.StatusCategoryKey,
		issue.Key,
	)

	if err != nil {
		r.log.Errorf("Error updating Jira issue with key %s: %v", issue.Key, err)
		return err
	}

	r.log.Infof("Updated Jira issue with key: %s", issue.Key)
	return nil
}
