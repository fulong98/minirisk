package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/minirisk/utils"
)

// LoggerMiddleware creates a middleware that logs HTTP requests
func LoggerMiddleware(logger *utils.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method
		ip := c.ClientIP()

		// Process request
		c.Next()

		// Calculate duration
		duration := time.Since(start)

		// Log request
		logger.LogRequest(method, path, ip, duration)
	}
}

// CORSMiddleware creates a middleware that handles CORS
func CORSMiddleware(allowedOrigins []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		if origin == "" {
			c.Next()
			return
		}

		// Check if origin is allowed
		allowed := false
		for _, allowedOrigin := range allowedOrigins {
			if origin == allowedOrigin {
				allowed = true
				break
			}
		}

		if allowed {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		}

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

// ErrorHandlerMiddleware creates a middleware that handles errors
func ErrorHandlerMiddleware(logger *utils.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// Check if there are any errors
		if len(c.Errors) > 0 {
			for _, err := range c.Errors {
				logger.LogError(err.Err, "HTTP Error")
			}

			// Send error response
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Internal Server Error",
			})
		}
	}
}

// AuthMiddleware creates a middleware that handles authentication
func AuthMiddleware(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get token from header
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
			c.Abort()
			return
		}

		// TODO: Implement JWT validation
		// For now, we'll just pass through
		c.Next()
	}
}

// RateLimitMiddleware creates a middleware that limits request rate
func RateLimitMiddleware(requestsPerMinute int) gin.HandlerFunc {
	// Simple in-memory rate limiter
	requests := make(map[string][]time.Time)

	return func(c *gin.Context) {
		ip := c.ClientIP()
		now := time.Now()

		// Clean up old requests
		requests[ip] = cleanOldRequests(requests[ip], now)

		// Check rate limit
		if len(requests[ip]) >= requestsPerMinute {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "Rate limit exceeded",
			})
			c.Abort()
			return
		}

		// Add new request
		requests[ip] = append(requests[ip], now)

		c.Next()
	}
}

// cleanOldRequests removes requests older than 1 minute
func cleanOldRequests(times []time.Time, now time.Time) []time.Time {
	var result []time.Time
	for _, t := range times {
		if now.Sub(t) < time.Minute {
			result = append(result, t)
		}
	}
	return result
}
