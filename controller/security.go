package controller

import (
	"github.com/JulesMike/api-er/helper"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

// TokenMismatch is used for ErrorFunc of CSRF middleware
func TokenMismatch(c *gin.Context) {
	helper.ResponseBadRequest(c, "CSRF Token mismatch")
}

// Token replies with a CSRF token
func Token(c *gin.Context) {
	token := csrf.GetToken(c)
	helper.ResponseSuccessPayload(c, "CSRF Token", token)
}
