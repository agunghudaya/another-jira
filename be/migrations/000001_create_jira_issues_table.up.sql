CREATE TABLE IF NOT EXISTS jira_issues (
    key VARCHAR(255) PRIMARY KEY,
    summary TEXT NOT NULL,
    description TEXT,
    status_name VARCHAR(255) NOT NULL,
    status_description TEXT,
    status_category_key VARCHAR(50) NOT NULL,
    status_category_name VARCHAR(255) NOT NULL,
    issue_type_name VARCHAR(255) NOT NULL,
    issue_type_description TEXT,
    priority_name VARCHAR(255) NOT NULL,
    project_id INTEGER NOT NULL,
    project_key VARCHAR(255) NOT NULL,
    project_name VARCHAR(255) NOT NULL,
    assignee_email VARCHAR(255),
    assignee_name VARCHAR(255),
    reporter_email VARCHAR(255),
    reporter_name VARCHAR(255),
    created TIMESTAMP WITH TIME ZONE NOT NULL,
    updated TIMESTAMP WITH TIME ZONE NOT NULL,
    due_date TIMESTAMP WITH TIME ZONE,
    time_estimate DOUBLE PRECISION,
    time_original_estimate DOUBLE PRECISION,
    aggregate_time_estimate DOUBLE PRECISION,
    aggregate_time_original_estimate DOUBLE PRECISION,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_jira_issues_project_key ON jira_issues(project_key);
CREATE INDEX IF NOT EXISTS idx_jira_issues_status_name ON jira_issues(status_name);
CREATE INDEX IF NOT EXISTS idx_jira_issues_assignee_email ON jira_issues(assignee_email);
CREATE INDEX IF NOT EXISTS idx_jira_issues_created ON jira_issues(created);
CREATE INDEX IF NOT EXISTS idx_jira_issues_updated ON jira_issues(updated); 