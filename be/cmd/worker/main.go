package main

import (
	"context"
	"log"

	"be/internal/delivery/cron"
	"be/internal/infrastructure/config"
	"be/internal/infrastructure/db"
	"be/internal/infrastructure/logger"
	rpJiraAtlassian "be/internal/repository/jira_atlassian"
	rpJiraDB "be/internal/repository/jira_db/impl"
	ucJiraSync "be/internal/usecase/uc_jira_sync"
)

func main() {
	// Load configuration from infrastructure
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	log := logger.InitLogger()

	// Create base context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Initialize database
	db, err := db.InitDB(cfg)
	if err != nil {
		log.Fatalf("Error initializing database: %s", err)
	}
	defer db.Close()

	// Initialize dependencies
	rpJiraDB := rpJiraDB.NewJiraDBRepository(cfg, log, db)
	rpJiraAtlassian := rpJiraAtlassian.NewJiraAtlassianRepository(cfg, log, db)

	ucJiraSync := ucJiraSync.NewJiraSyncUsecase(cfg, log, rpJiraDB, rpJiraAtlassian)

	c := cron.NewWorker(log, ucJiraSync)

	log.Info("Starting Cron Jobs...")
	c.Start(ctx)

	// Keep the process running
	select {}
}
