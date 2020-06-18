package controller

import (
	"github.com/JulesMike/api-er/helper"
	"github.com/gin-gonic/gin"
)

// TokenMismatch is used for ErrorFunc of CSRF middleware
func TokenMismatch(c *gin.Context) {
	helper.ResponseBadRequest(c, "CSRF Token mismatch")
}
