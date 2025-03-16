package main

import (
	"fmt"
	"log"

	"be/internal/delivery/cron"
	"be/internal/infrastructure/config"
	"be/internal/infrastructure/db"
	"be/internal/repository"
	"be/internal/usecase"
)

func main() {
	// Load configuration from infrastructure
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	c := cron.New()

	// Initialize database
	db, err := db.InitDB(cfg)
	if err != nil {
		log.Fatalf("Error initializing database: %s", err)
	}
	defer db.Close()

	// Initialize dependencies
	syncRepo := repository.NewSyncRepository(cfg, db)         // Database/API connection
	jiraSync := usecase.NewJiraSyncUsecase(cfg, db, syncRepo) // Business logic

	cron.RegisterJobs(c, jiraSync)

	fmt.Println("Starting Cron Jobs...")
	c.Start()

	// Keep the process running
	select {}
}
