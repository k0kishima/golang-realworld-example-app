package handlers

import (
	"time"

	"github.com/gin-gonic/gin"
)

func respondWithError(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{"error": message})
}

func formatTimeForAPI(t time.Time) string {
	return t.Format("2006-01-02T15:04:05.000Z")
}
