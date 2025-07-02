package config

import "time"

// Config defines the interface for application configuration
type Config interface {
	// GetJiraConfig returns the Jira configuration
	GetJiraConfig() JiraConfig
	
	// GetDBConfig returns the database configuration
	GetDBConfig() DBConfig
	
	// GetSyncConfig returns the synchronization configuration
	GetSyncConfig() SyncConfig
}

// JiraConfig represents Jira-specific configuration
type JiraConfig struct {
	BaseURL      string
	Username     string
	APIToken     string
	ProjectKey   string
	BatchSize    int
	Timeout      time.Duration
}

// DBConfig represents database-specific configuration
type DBConfig struct {
	Host         string
	Port         int
	User         string
	Password     string
	Database     string
	MaxOpenConns int
	MaxIdleConns int
}

// SyncConfig represents synchronization-specific configuration
type SyncConfig struct {
	Interval     time.Duration
	MaxRetries   int
	RetryDelay   time.Duration
	BatchSize    int
} 