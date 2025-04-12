package db

import (
	"database/sql"
	"fmt"

	"be/internal/infrastructure/config"

	_ "github.com/lib/pq"
)

func InitDB(cfg *config.Config) (DB, error) {
	dbHost := cfg.GetString("database.host")
	dbPort := cfg.GetInt("database.port")
	dbUser := cfg.GetString("database.user")
	dbPassword := cfg.GetString("database.password")
	dbName := cfg.GetString("database.dbname")

	// Create database connection string
	dbConnStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	// Open database connection
	db, err := sql.Open("postgres", dbConnStr)
	if err != nil {
		return nil, fmt.Errorf("error connecting to the database: %w", err)
	}

	// Verify the connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error verifying connection to the database: %w", err)
	}

	return db, nil
}
