package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ResponseUnauthorized wraps c.JSON
func ResponseUnauthorized(c *gin.Context, msg string) {
	c.JSON(http.StatusUnauthorized, gin.H{"error": msg})
	c.Abort()
}

// ResponseBadRequest wraps c.JSON
func ResponseBadRequest(c *gin.Context, msg string) {
	c.JSON(http.StatusBadRequest, gin.H{"error": msg})
	c.Abort()
}

// ResponseNotFound wraps c.JSON
func ResponseNotFound(c *gin.Context, msg string) {
	c.JSON(http.StatusNotFound, gin.H{"error": msg})
	c.Abort()
}

// ResponseInternalServerError wraps c.JSON
func ResponseInternalServerError(c *gin.Context, msg string) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
	c.Abort()
}

// ResponseSuccess wraps c.JSON
func ResponseSuccess(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{"success": msg})
	c.Abort()
}

// ResponseSuccessPayload wraps c.JSON
func ResponseSuccessPayload(c *gin.Context, msg string, payload interface{}) {
	c.JSON(http.StatusOK, gin.H{"success": msg, "payload": payload})
	c.Abort()
}
