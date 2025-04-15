package httpserver

import (
	"be/internal/infrastructure/logger"

	"github.com/gin-gonic/gin"
)

// InitServer initializes the HTTP server with custom logging
func InitServer() *gin.Engine {
	r := gin.New()
	log := logger.InitLogger()
	r.Use(logger.LoggerMiddleware(log), gin.Recovery())
	return r
}
