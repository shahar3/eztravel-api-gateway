package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Logger returns a Gin middleware that logs requests.
func Logger(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		// Process request
		c.Next()
		duration := time.Since(startTime)

		logger.WithFields(logrus.Fields{
			"status":   c.Writer.Status(),
			"method":   c.Request.Method,
			"path":     c.Request.RequestURI,
			"duration": duration,
			"clientIP": c.ClientIP(),
		}).Info("handled request")
	}
}

// Recovery returns a middleware that recovers from any panics and writes a 500 if there was one.
func Recovery(logger *logrus.Logger) gin.HandlerFunc {
	return gin.RecoveryWithWriter(logger.Out)
}
