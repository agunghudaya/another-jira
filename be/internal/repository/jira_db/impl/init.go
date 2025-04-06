package jira_db_impl

import (
	repository "be/internal/domain/repository"
	"be/internal/infrastructure/config"

	"be/internal/repository/jira_db"
	"context"
	"database/sql"

	"github.com/sirupsen/logrus"
)

func NewJiraDBRepository(cfg *config.Config, log *logrus.Logger, db *sql.DB) jira_db.JiraDBRepository {
	return &jiraDBRepository{cfg: cfg, db: db, log: log}
}

type jiraDBRepository struct {
	cfg *config.Config
	db  *sql.DB
	log *logrus.Logger
}

func (r *jiraDBRepository) UpdateJiraIssue(ctx context.Context, issue repository.JiraIssueEntity) error {
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
