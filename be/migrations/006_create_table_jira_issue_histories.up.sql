CREATE TABLE jira_issue_histories (
    id SERIAL PRIMARY KEY, -- Unique identifier for the issue history
    issue_id INT NOT NULL, 
    field VARCHAR(255), 
    old_value VARCHAR(255), 
    new_value VARCHAR(255), 

    created TIMESTAMP NOT NULL,

    CONSTRAINT fk_jira_issue
        FOREIGN KEY (issue_id)
        REFERENCES jira_issues (id)
        ON DELETE CASCADE

);




