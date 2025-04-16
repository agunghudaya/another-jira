package db

import (
	"context"
	"database/sql"
)

type DB interface {
	Close() error
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	Ping() error
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}

type SQLDB struct {
	db *sql.DB
}

// Ping checks the connection to the database.
func (s *SQLDB) Ping() error {
	return s.db.Ping()
}

// Close closes the database connection.
func (s *SQLDB) Close() error {
	return s.db.Close()
}

// Exec executes a query without returning any rows.
func (s *SQLDB) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return s.db.ExecContext(ctx, query, args...)
}

// Query executes a query that returns rows.
func (s *SQLDB) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return s.db.QueryContext(ctx, query, args...)
}

// QueryRow executes a query that is expected to return at most one row.
func (s *SQLDB) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return s.db.QueryRowContext(ctx, query, args...)
}
