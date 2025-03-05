CREATE TABLE jira_user_sync_history (
    id SERIAL PRIMARY KEY,               -- Unique sync ID
    jira_id varchar(200) NOT NULL,        
    sync_date DATE NOT NULL,              -- The date of the sync
    started_at TIMESTAMP NOT NULL,        -- When the sync started
    finished_at TIMESTAMP,                -- When the sync finished
    status VARCHAR(20) NOT NULL,          -- "in_progress", "success", "failed"
    error_message TEXT,                   -- Stores error details (if any)
    records_synced INT DEFAULT 0,         -- How many records were synced
    total_expected_records INT DEFAULT 0, -- Total records expected from Jira
    sync_attempt INT DEFAULT 1,           -- Number of sync attempts in a day
    created_at TIMESTAMP DEFAULT NOW(),   -- Auto-generated timestamp

    UNIQUE (jira_id, sync_date, sync_attempt)  -- Prevent duplicate runs for the same attempt
);
