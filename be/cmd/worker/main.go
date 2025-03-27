package main

import (
	"context"
	"fmt"
	"log"

	"be/internal/delivery/cron"
	"be/internal/infrastructure/config"
	"be/internal/infrastructure/db"
	"be/internal/infrastructure/logger"
	"be/internal/repository"
	"be/internal/usecase"
)

func main() {
	// Load configuration from infrastructure
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	log := logger.InitLogger()
	log.Info("Starting server...")

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
	syncRepo := repository.NewSyncRepository(cfg, log, db)
	jiraSync := usecase.NewJiraSyncUsecase(cfg, log, syncRepo)

	c := cron.NewWorker(log, jiraSync)

	fmt.Println("Starting Cron Jobs...")
	c.Start(ctx)

	// Keep the process running
	select {}
}
