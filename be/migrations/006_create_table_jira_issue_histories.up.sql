CREATE TABLE jira_issue_histories (
    id SERIAL PRIMARY KEY, -- Unique identifier for the issue history
    key VARCHAR(255), 
    field VARCHAR(255), 
    old_value text, 
    new_value text, 

    created TIMESTAMP NOT NULL
);




