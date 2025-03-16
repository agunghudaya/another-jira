package logger

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// InitLogger initializes the global logger
func InitLogger() *logrus.Logger {
	logger := logrus.New()

	// Set output to stdout
	logger.SetOutput(os.Stdout)

	// Set log format (JSON or Text)
	logger.SetFormatter(&logrus.JSONFormatter{}) // Use JSON for structured logging

	// Set log level (INFO by default)
	logger.SetLevel(logrus.InfoLevel)

	return logger
}

// LoggerMiddleware is a Gin middleware for structured logging
func LoggerMiddleware() gin.HandlerFunc {
	logger := InitLogger()
	return func(c *gin.Context) {
		// Before request
		path := c.Request.URL.Path
		method := c.Request.Method
		clientIP := c.ClientIP()

		// Process request
		c.Next()

		// After request
		statusCode := c.Writer.Status()
		logger.WithFields(logrus.Fields{
			"method": method,
			"path":   path,
			"status": statusCode,
			"ip":     clientIP,
		}).Info("Request completed")
	}
}
