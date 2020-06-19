package controller

import (
	"time"

	"github.com/JulesMike/api-er/helper"
	"github.com/JulesMike/api-er/repository"
	"github.com/gin-gonic/gin"
)

// App represents the app controller
type App struct{}

// NewApp returns a new App
func NewApp(userRepo *repository.User) *App {
	return &App{}
}

// Ping pongs (Get it?)
func (c *App) Ping(ctx *gin.Context) {
	helper.ResponseSuccessPayload(ctx, "pong", time.Now())
}

// Panic simulates a panic
func (c *App) Panic(ctx *gin.Context) {
	panic("oh la la!")
}

// NotFound is used as not found handler
func (c *App) NotFound(ctx *gin.Context) {
	helper.ResponseNotFound(ctx, "app:route:unknown")
}
