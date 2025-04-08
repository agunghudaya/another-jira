CREATE TABLE jira_issues (
    key VARCHAR(50) PRIMARY KEY, -- Jira issue key (e.g., PROJ-123)
    self TEXT NOT NULL, -- Jira API self URL
    url TEXT NOT NULL, -- Jira web URL

    assignee_email VARCHAR(255), -- Email of the assignee
    assignee_name VARCHAR(255), -- Name of the assignee
    reporter_email VARCHAR(255), -- Email of the reporter
    reporter_name VARCHAR(255), -- Name of the reporter
    creator_email VARCHAR(255), -- Email of the creator
    creator_name VARCHAR(255), -- Name of the creator

    summary TEXT NOT NULL, -- Issue summary/title
    description TEXT, -- Issue description

    created TIMESTAMP NOT NULL, -- Issue creation timestamp
    updated TIMESTAMP NOT NULL, -- Last update timestamp
    duedate DATE, -- Due date of the issue
    statuscategorychangedate TIMESTAMP, -- Status category change timestamp

    issue_type_name VARCHAR(255), -- Name of the issue type
    issue_type_desc TEXT, -- Description of the issue type

    project_id INT NOT NULL, -- Project ID
    project_key VARCHAR(50) NOT NULL, -- Project key (e.g., PROJ)
    project_name VARCHAR(255) NOT NULL, -- Project name

    priority_name VARCHAR(255), -- Priority name

    status_name VARCHAR(255), -- Status name
    status_desc TEXT, -- Status description
    status_category_name VARCHAR(255), -- Status category name
    status_category_key VARCHAR(255), -- Status category key

    time_estimate DOUBLE PRECISION,
    time_original_estimate DOUBLE PRECISION, 
    aggregate_time_estimate DOUBLE PRECISION,
    aggregate_time_original_estimate DOUBLE PRECISION

);




