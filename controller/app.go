package controller

import (
	"time"

	"github.com/JulesMike/api-er/helper"
	"github.com/gin-gonic/gin"
)

// Ping pongs (Get it?)
func Ping(c *gin.Context) {
	helper.ResponseSuccessPayload(c, "pong", time.Now())
}

// Panic simulates a panic
func Panic(c *gin.Context) {
	panic("oh la la!")
}

// NotFound is used as not found handler
func NotFound(c *gin.Context) {
	helper.ResponseNotFound(c, "app:route:unknown")
}
