package handlers

import (
	"github.com/gin-gonic/gin"
)

func respondWithError(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{"error": message})
}
