package jiradb

import (
	"be/internal/domain/entity"
	"be/internal/domain/repository"
	"be/internal/infrastructure/database"
	"be/internal/infrastructure/jira"
	"context"
	"database/sql"
	"fmt"
)

// jiraRepository implements repository.JiraRepository
type jiraRepository struct {
	db         *database.Connection
	jiraClient *jira.Client
}

// NewJiraRepository creates a new Jira repository
func NewJiraRepository(db *database.Connection, jiraClient *jira.Client) repository.JiraRepository {
	return &jiraRepository{
		db:         db,
		jiraClient: jiraClient,
	}
}

// SaveIssues implements repository.JiraRepository
func (r *jiraRepository) SaveIssues(ctx context.Context, issues []entity.JiraIssue) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	for _, issue := range issues {
		if err := r.saveIssue(ctx, tx, issue); err != nil {
			return err
		}
	}

	return tx.Commit()
}

// saveIssue saves a single issue
func (r *jiraRepository) saveIssue(ctx context.Context, tx *sql.Tx, issue entity.JiraIssue) error {
	query := `
		INSERT INTO jira_issues (
			key, summary, description, status_name, status_description,
			status_category_key, status_category_name, issue_type_name,
			issue_type_description, priority_name, project_id, project_key,
			project_name, assignee_email, assignee_name, reporter_email,
			reporter_name, created, updated, due_date, time_estimate,
			time_original_estimate, aggregate_time_estimate,
			aggregate_time_original_estimate
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14,
			$15, $16, $17, $18, $19, $20, $21, $22, $23, $24
		) ON CONFLICT (key) DO UPDATE SET
			summary = EXCLUDED.summary,
			description = EXCLUDED.description,
			status_name = EXCLUDED.status_name,
			status_description = EXCLUDED.status_description,
			status_category_key = EXCLUDED.status_category_key,
			status_category_name = EXCLUDED.status_category_name,
			issue_type_name = EXCLUDED.issue_type_name,
			issue_type_description = EXCLUDED.issue_type_description,
			priority_name = EXCLUDED.priority_name,
			project_id = EXCLUDED.project_id,
			project_key = EXCLUDED.project_key,
			project_name = EXCLUDED.project_name,
			assignee_email = EXCLUDED.assignee_email,
			assignee_name = EXCLUDED.assignee_name,
			reporter_email = EXCLUDED.reporter_email,
			reporter_name = EXCLUDED.reporter_name,
			created = EXCLUDED.created,
			updated = EXCLUDED.updated,
			due_date = EXCLUDED.due_date,
			time_estimate = EXCLUDED.time_estimate,
			time_original_estimate = EXCLUDED.time_original_estimate,
			aggregate_time_estimate = EXCLUDED.aggregate_time_estimate,
			aggregate_time_original_estimate = EXCLUDED.aggregate_time_original_estimate
	`

	_, err := tx.ExecContext(ctx, query,
		issue.Key,
		issue.Summary,
		issue.Description,
		issue.Status.Name,
		issue.Status.Description,
		issue.Status.Category.Key,
		issue.Status.Category.Name,
		issue.Type.Name,
		issue.Type.Description,
		issue.Priority.Name,
		issue.Project.ID,
		issue.Project.Key,
		issue.Project.Name,
		issue.Assignee.Email,
		issue.Assignee.DisplayName,
		issue.Reporter.Email,
		issue.Reporter.DisplayName,
		issue.Created,
		issue.Updated,
		issue.DueDate,
		issue.TimeTracking.CurrentEstimate,
		issue.TimeTracking.OriginalEstimate,
		issue.TimeTracking.AggregateEstimate,
		issue.TimeTracking.AggregateOriginalEstimate,
	)

	return err
}

// GetIssues implements repository.JiraRepository
func (r *jiraRepository) GetIssues(ctx context.Context, filter repository.JiraFilter) ([]entity.JiraIssue, error) {
	query := `
		SELECT * FROM jira_issues
		WHERE ($1 = '' OR project_key = $1)
		AND ($2 = '' OR status_name = $2)
		AND ($3 = '' OR assignee_email = $3)
		AND ($4::timestamp IS NULL OR created >= $4)
		AND ($5::timestamp IS NULL OR created <= $5)
		AND ($6::timestamp IS NULL OR updated >= $6)
		AND ($7::timestamp IS NULL OR updated <= $7)
		ORDER BY created DESC
		LIMIT $8 OFFSET $9
	`

	rows, err := r.db.QueryContext(ctx, query,
		filter.ProjectKey,
		filter.Status,
		filter.AssigneeEmail,
		filter.CreatedAfter,
		filter.CreatedBefore,
		filter.UpdatedAfter,
		filter.UpdatedBefore,
		filter.Limit,
		filter.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var issues []entity.JiraIssue
	for rows.Next() {
		var issue entity.JiraIssue
		if err := r.scanIssue(rows, &issue); err != nil {
			return nil, err
		}
		issues = append(issues, issue)
	}

	return issues, rows.Err()
}

// GetIssueByKey implements repository.JiraRepository
func (r *jiraRepository) GetIssueByKey(ctx context.Context, key string) (*entity.JiraIssue, error) {
	query := `SELECT * FROM jira_issues WHERE key = $1`

	var issue entity.JiraIssue
	err := r.db.QueryRowContext(ctx, query, key).Scan(
		&issue.Key,
		&issue.Summary,
		&issue.Description,
		&issue.Status.Name,
		&issue.Status.Description,
		&issue.Status.Category.Key,
		&issue.Status.Category.Name,
		&issue.Type.Name,
		&issue.Type.Description,
		&issue.Priority.Name,
		&issue.Project.ID,
		&issue.Project.Key,
		&issue.Project.Name,
		&issue.Assignee.Email,
		&issue.Assignee.DisplayName,
		&issue.Reporter.Email,
		&issue.Reporter.DisplayName,
		&issue.Created,
		&issue.Updated,
		&issue.DueDate,
		&issue.TimeTracking.CurrentEstimate,
		&issue.TimeTracking.OriginalEstimate,
		&issue.TimeTracking.AggregateEstimate,
		&issue.TimeTracking.AggregateOriginalEstimate,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &issue, nil
}

// scanIssue scans a database row into a JiraIssue
func (r *jiraRepository) scanIssue(rows *sql.Rows, issue *entity.JiraIssue) error {
	return rows.Scan(
		&issue.Key,
		&issue.Summary,
		&issue.Description,
		&issue.Status.Name,
		&issue.Status.Description,
		&issue.Status.Category.Key,
		&issue.Status.Category.Name,
		&issue.Type.Name,
		&issue.Type.Description,
		&issue.Priority.Name,
		&issue.Project.ID,
		&issue.Project.Key,
		&issue.Project.Name,
		&issue.Assignee.Email,
		&issue.Assignee.DisplayName,
		&issue.Reporter.Email,
		&issue.Reporter.DisplayName,
		&issue.Created,
		&issue.Updated,
		&issue.DueDate,
		&issue.TimeTracking.CurrentEstimate,
		&issue.TimeTracking.OriginalEstimate,
		&issue.TimeTracking.AggregateEstimate,
		&issue.TimeTracking.AggregateOriginalEstimate,
	)
}

// DeleteIssue implements repository.JiraRepository
func (r *jiraRepository) DeleteIssue(ctx context.Context, key string) error {
	query := `DELETE FROM jira_issues WHERE key = $1`
	_, err := r.db.ExecContext(ctx, query, key)
	return err
}

// UpdateIssue implements repository.JiraRepository
func (r *jiraRepository) UpdateIssue(ctx context.Context, issue entity.JiraIssue) error {
	return r.saveIssue(ctx, nil, issue)
}
