package router

import (
	httpDelivery "be/internal/delivery/http"
	middleware "be/internal/middleware/auth"
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/spf13/viper"
)

func InitRouter() *chi.Mux {
	r := chi.NewRouter()

	// Get frontend URL from Viper
	frontendURL := viper.GetString("fe.url")

	// Log the frontend URL
	log.Printf("Frontend URL: %s", frontendURL)

	// Add CORS middleware
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{frontendURL},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// Add middleware
	r.Use(middleware.AuthMiddleware)

	// Define routes
	r.Get("/api/health", httpDelivery.HandleRequest)

	return r
}
