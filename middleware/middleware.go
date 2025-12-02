package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// RequestLogger logs method, path, status and latency for each request.
func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		// before request
		method := c.Request.Method
		path := c.Request.URL.Path

		c.Next()

		// after request
		latency := time.Since(start)
		status := c.Writer.Status()
		log.Printf("%s %s -> %d (%s)", method, path, status, latency)
	}
}
