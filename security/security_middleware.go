package security

import (
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

// CSRF is the csrf middleware
func CSRF(secret string, securityCtrl *Controller) gin.HandlerFunc {
	return csrf.Middleware(csrf.Options{
		Secret:    secret,
		ErrorFunc: securityCtrl.TokenMismatch,
	})
}
