package main

import (
	config "be/internal/infrastructure/config"
	db "be/internal/infrastructure/db"
	server "be/internal/infrastructure/http_server"
	logger "be/internal/infrastructure/logger"
)

func main() {

	// Initialize logger
	log := logger.InitLogger()
	log.Info("Starting server...")

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

	// Create Gin Router
	r := server.InitServer()

	// Start server
	port := cfg.GetString("server.port")
	address := ":" + port // Ensure correct format ":8080"

	log.Infof("Server is running on %s", address)
	if err := r.Run(address); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
