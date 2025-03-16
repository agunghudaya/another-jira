package main

import (
	config "be/internal/infrastructure/config"
	db "be/internal/infrastructure/db"

	"be/internal/router"
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

func main() {

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

	// Initialize router
	r := router.InitRouter()

	// Start server
	port := viper.GetString("server.port")
	log.Printf("Starting server on :%s", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), r); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
