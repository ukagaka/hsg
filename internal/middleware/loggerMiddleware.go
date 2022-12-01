package middleware

import (
	"hxsg/base"
	"log"
	"time"
)

func LoggerMiddleware() base.HandlerFunc {
	return func(c *base.Context) {
		// Start timer
		t := time.Now()
		// Process request
		c.Next()
		// Calculate resolution time
		log.Printf("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
