package middleware

import (
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

// CSRF is the csrf middleware
func CSRF(secret string) gin.HandlerFunc {
	return csrf.Middleware(csrf.Options{
		Secret: secret,
		ErrorFunc: func(c *gin.Context) {
			c.JSON(400, gin.H{"error": "CSRF Token mismatch"})
			c.Abort()
		},
	})
}
