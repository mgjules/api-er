package middleware

import (
	"github.com/JulesMike/api-er/controller"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

// CSRF is the csrf middleware
func CSRF(secret string) gin.HandlerFunc {
	return csrf.Middleware(csrf.Options{
		Secret:    secret,
		ErrorFunc: controller.TokenMismatch,
	})
}
