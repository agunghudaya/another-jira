package main

import (
	"be/pkg/handlers"
	"be/pkg/middleware"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")    // File name without extension
	viper.SetConfigType("json")      // Config format
	viper.AddConfigPath("./configs") // Search in ./configs directory
	viper.AddConfigPath(".")         // (Optional) Fallback to current directory

	// Read the config file
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	feUrl := viper.GetString("fe.url")

	r := chi.NewRouter()

	// Add CORS middleware
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{feUrl}, // Allow your frontend URL
		AllowedMethods:   []string{"GET", "POST"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// Add middleware
	r.Use(middleware.AuthMiddleware)

	r.Get("/", handlers.HandleRequest)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
