package utils

import (
	"github.com/gin-gonic/gin"
)

// SendError sends a standardized JSON error response
func SendError(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{"error": message})
}

// SendSuccess sends a standardized JSON success response
func SendSuccess(c *gin.Context, status int, data interface{}) {
	c.JSON(status, data)
}
