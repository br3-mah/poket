package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		
		// Process request
		c.Next()
		
		// Log details
		log.Printf(
			"[%d] %s %s %s %v",
			c.Writer.Status(),
			c.Request.Method,
			c.Request.URL.Path,
			c.ClientIP(),
			time.Since(start),
		)
		
		// Log 500 errors with request details
		if c.Writer.Status() >= 500 {
			log.Printf(
				"ERROR [500] | Path: %s | Method: %s | IP: %s | Error: %v",
				c.Request.URL.Path,
				c.Request.Method,
				c.ClientIP(),
				c.Errors.Last(),
			)
		}
	}
}