package routes

import (
	delivery "be/internal/delivery/http"

	"github.com/gin-gonic/gin"
)

// HandlerRegistry holds all route handlers
type HandlerRegistry struct {
	HealthHandler *delivery.HealthHandler
	UserHandler   *delivery.UserHandler
}

// RegisterRoutes registers all routes using the handler registry
func RegisterRoutes(r *gin.Engine, hr *HandlerRegistry) {
	api := r.Group("/api")
	{
		api.GET("/health", hr.HealthHandler.HealthCheck)

		userRoutes := api.Group("/users")
		{
			//userRoutes.GET("/:id", hr.UserHandler.GetUser)
			userRoutes.GET("/list", hr.UserHandler.GetUsers)

		}

		// orderRoutes := api.Group("/orders")
		// {
		// 	orderRoutes.GET("/:id", hr.OrderHandler.GetOrder)
		// }

		// productRoutes := api.Group("/products")
		// {
		// 	productRoutes.GET("/:id", hr.ProductHandler.GetProduct)
		// }
	}
}
