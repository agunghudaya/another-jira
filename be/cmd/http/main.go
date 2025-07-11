package main

import (
	delivery "be/internal/delivery/http"
	config "be/internal/infrastructure/config"
	db "be/internal/infrastructure/db"
	server "be/internal/infrastructure/http_server"
	logger "be/internal/infrastructure/logger"
	routes "be/internal/routes"

	rpJiraDB "be/internal/repository/jira_db/impl"
	ucUser "be/internal/usecase/uc_user"

	"context"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	// Initialize logger
	log := logger.InitLogger()
	log.Info("Starting server...")

	// Create base context
	_, cancel := context.WithCancel(context.Background())

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Initialize database
	db, err := db.InitDB(cfg)
	if err != nil {
		log.Fatalf("Error initializing database: %s", err)
	}
	defer db.Close()

	// Handle graceful shutdown
	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
		<-quit
		log.Info("Shutting down server...")
		cancel()
	}()

	// Initialize dependencies
	rpJiraDB := rpJiraDB.NewJiraDBRepository(cfg, log, db)
	ucUser := ucUser.NewUsecaseUser(cfg, log, rpJiraDB)

	// Create Gin Router
	r := server.InitServer()

	// Initialize handlers
	hr := &routes.HandlerRegistry{
		HealthHandler: delivery.NewHealthHandler(r, log, db),
		UserHandler:   delivery.NewUserHandler(r, log, ucUser),
	}

	routes.RegisterRoutes(r, hr)

	// Start server
	port := cfg.GetString("server.port")
	address := ":" + port // Ensure correct format ":8080"

	log.Infof("Server is running on %s", address)
	if err := r.Run(address); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
