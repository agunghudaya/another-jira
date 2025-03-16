package routes

import (
	handlers "be/internal/delivery/http"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes sets up API routes
func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		userRoutes(api)
		healthRoutes(api)
	}
}

// userRoutes defines user-related routes
func userRoutes(r *gin.RouterGroup) {
	users := r.Group("/users")
	{
		users.GET("/:id", handlers.GetUserHandler)
		users.POST("/", handlers.CreateUserHandler)
	}
}

// healthRoutes defines health check routes
func healthRoutes(r *gin.RouterGroup) {
	r.GET("/health", handlers.GetHealthHandler)
}
