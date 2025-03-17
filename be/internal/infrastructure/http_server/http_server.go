package httpserver

import (
	"be/internal/infrastructure/logger"

	"github.com/gin-gonic/gin"
)

// InitServer initializes the HTTP server with custom logging
func InitServer() *gin.Engine {
	r := gin.New()
	r.Use(logger.LoggerMiddleware(), gin.Recovery()) // Use custom logger middleware
	return r
}
