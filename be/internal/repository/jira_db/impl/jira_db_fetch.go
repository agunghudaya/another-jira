package jiradbimpl

import (
	repository "be/internal/domain/repository"
	"context"
	"database/sql"
	"fmt"
)

func (r *jiraDBRepository) FetchJiraIssue(ctx context.Context, issueKey string) (repository.JiraIssueEntity, error) {
	query := `
        SELECT "key", "self", url, assignee_email, assignee_name, reporter_email, reporter_name, 
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

	var issue repository.JiraIssueEntity
	err := r.db.QueryRowContext(ctx, query, issueKey).Scan(
		&issue.Key, &issue.Self, &issue.URL,
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
			return repository.JiraIssueEntity{}, nil
		}
		r.log.Errorf("Error fetching Jira issue for key %s: %v", issueKey, err)
		return repository.JiraIssueEntity{}, err
	}

	return issue, nil
}

func (r *jiraDBRepository) FetchJiraAssignedIssuesByEmail(ctx context.Context, email string) ([]repository.JiraIssueEntity, error) {
	query := `
        SELECT "key", "self", url, assignee_email, assignee_name, reporter_email, reporter_name, 
               creator_email, creator_name, summary, description, created, updated, duedate, 
               statuscategorychangedate, issue_type_name, issue_type_desc, project_id, project_key, 
               project_name, priority_name, status_name, status_desc, status_category_name, 
               status_category_key, time_estimate, time_original_estimate, aggregate_time_estimate, 
               aggregate_time_original_estimate
        FROM public.jira_issues
        WHERE 
            assignee_email = $1;
    `

	rows, err := r.db.QueryContext(ctx, query, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var issues []repository.JiraIssueEntity
	for rows.Next() {
		var issue repository.JiraIssueEntity
		err := rows.Scan(
			&issue.Key, &issue.Self, &issue.URL,
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
			return nil, err
		}
		issues = append(issues, issue)
	}

	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return issues, nil
}

func (r *jiraDBRepository) FetchPendingSync(ctx context.Context, jiraID string) ([]repository.SyncHistory, error) {
	query := `
		SELECT id, jira_id, sync_date, started_at, finished_at, status, 
		       error_message, records_synced, total_expected_records, sync_attempt, created_at
		FROM jira_user_sync_history 
		WHERE 
			jira_id = $1 
			AND sync_date = CURRENT_DATE
		ORDER BY started_at DESC ;
	`

	rows, err := r.db.QueryContext(ctx, query, jiraID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var syncs []repository.SyncHistory

	for rows.Next() {
		var sync repository.SyncHistory
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

func (r *jiraDBRepository) FetchUserList(ctx context.Context) ([]repository.UserEntity, error) {

	select {
	case <-ctx.Done():
		r.log.Infof("FetchUserList operation was canceled!")
		return nil, ctx.Err()
	default:
		// Continue processing
	}

	query := `select ID, username, jira_user_id, status from users u where status = 'active';`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close() // Ensure rows are closed when function exits

	var users []repository.UserEntity

	for rows.Next() {
		var user repository.UserEntity
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

func (r *jiraDBRepository) FetchUserByID(ctx context.Context, jiraID string) (repository.UserEntity, error) {
	select {
	case <-ctx.Done():
		r.log.Infof("FetchUserByID operation was canceled!")
		return repository.UserEntity{}, ctx.Err()
	default:
		// Continue processing
	}

	query := `SELECT ID, username, jira_user_id, status, email FROM users u WHERE id = $1 AND status = 'active';`

	var user repository.UserEntity
	err := r.db.QueryRowContext(ctx, query, jiraID).Scan(&user.ID, &user.Username, &user.JiraUserID, &user.Status, &user.Email)
	if err != nil || err == sql.ErrNoRows {
		if err == sql.ErrNoRows {
			return repository.UserEntity{}, fmt.Errorf("error fetching user by ID %s", jiraID)
		}
		return repository.UserEntity{}, err
	}

	return user, nil
}
