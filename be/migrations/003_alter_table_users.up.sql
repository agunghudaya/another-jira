ALTER TABLE users
ADD COLUMN jira_user_id VARCHAR(255) UNIQUE, -- Assuming Jira ID is a string
ADD COLUMN status VARCHAR(20) DEFAULT 'active', -- Example: 'active', 'inactive'
ADD COLUMN created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
ADD COLUMN updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP;
