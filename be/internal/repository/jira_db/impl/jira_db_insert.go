package jira_db

import (
	repository "be/internal/domain/repository"
	"context"
	"time"
)

func (r *jiraDBRepository) InsertJiraIssue(ctx context.Context, issue repository.JiraIssue) error {
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

func (r *jiraDBRepository) InsertJiraIssueHistory(ctx context.Context, history repository.JiraIssueHistory) (int, error) {
	query := `
		INSERT INTO jira_issue_histories (
            issue_id, 
            field, 
            old_value, 
            new_value, 
            created
        ) VALUES (
            $1, $2, $3, $4, $5
        ) RETURNING id; `

	var id int
	err := r.db.QueryRowContext(ctx, query,
		history.IssueID,
		history.Field,
		history.Oldvalue,
		history.NewValue,
		history.Created,
	).Scan(&id)

	if err != nil {
		r.log.Errorf("Error inserting Jira issue history: %v", err)
		return 0, err
	}

	r.log.Infof("Inserted Jira issue history with ID: %d", id)
	return id, nil
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
