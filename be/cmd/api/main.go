package main

import (
	"be/internal/repository"
	"be/internal/router"
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

func main() {
	// Initialize Viper
	viper.SetConfigName("config")    // name of config file (without extension)
	viper.SetConfigType("json")      // or viper.SetConfigType("YAML")
	viper.AddConfigPath("./configs") // optionally look for config in the configs directory
	viper.AddConfigPath(".")         // optionally look for config in the working directory
	viper.AutomaticEnv()             // read in environment variables that match

	// Set default values
	viper.SetDefault("fe.url", "http://localhost:3000")

	// Read the config file
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading config file, %s", err)
	}

	// Initialize database
	db, err := repository.InitDB()
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
