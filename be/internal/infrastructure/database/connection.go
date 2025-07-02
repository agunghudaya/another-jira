package database

import (
	"be/internal/domain/config"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Connection represents a database connection
type Connection struct {
	*sql.DB
}

// NewConnection creates a new database connection
func NewConnection(cfg config.DBConfig) (*Connection, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.Database,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	// Set connection pool settings
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)

	// Test the connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &Connection{db}, nil
}

// Close closes the database connection
func (c *Connection) Close() error {
	return c.DB.Close()
}
