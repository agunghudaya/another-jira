package logger

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func InitLogger() Logger {
	base := logrus.New()

	base.SetOutput(os.Stdout)
	base.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
	})
	base.SetLevel(logrus.InfoLevel)

	return &LogrusAdapter{logger: base}
}

func LoggerMiddleware(logger Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		method := c.Request.Method
		clientIP := c.ClientIP()

		c.Next()

		statusCode := c.Writer.Status()
		logger.WithFields(map[string]any{
			"method": method,
			"path":   path,
			"status": statusCode,
			"ip":     clientIP,
		}).Info("Request completed")
	}
}
