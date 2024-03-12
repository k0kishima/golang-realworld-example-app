package middlewares

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	allowedOrigins := getAllowedOrigins()
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")

		if contains(allowedOrigins, origin) {
			c.Header("Access-Control-Allow-Origin", origin)
		} else {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Origin not allowed"})
			return
		}

		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func getAllowedOrigins() []string {
	return strings.Split(os.Getenv("ALLOWED_ORIGINS"), ",")
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
