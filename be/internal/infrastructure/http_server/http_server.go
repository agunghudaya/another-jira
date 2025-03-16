package httpserver

import (
	"be/internal/infrastructure/logger"
	"be/internal/routes"

	"github.com/gin-gonic/gin"
)

// InitServer initializes the HTTP server with custom logging
func InitServer() *gin.Engine {
	r := gin.New()
	r.Use(logger.LoggerMiddleware(), gin.Recovery()) // Use custom logger middleware

	routes.RegisterRoutes(r)

	return r
}
