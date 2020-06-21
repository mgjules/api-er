package app

import (
	"time"

	"github.com/JulesMike/api-er/helper"
	"github.com/gin-gonic/gin"
)

// Controller represents the app controller
type Controller struct{}

// NewController returns a new app controller
func NewController() *Controller {
	return &Controller{}
}

// AttachRoutes attaches the controller's routes to gin.RouterGroup
func (c *Controller) AttachRoutes(r *gin.RouterGroup) {
	r.GET("/ping", c.Ping)

	r.GET("/panic", c.Panic)
}

// Ping pongs (Get it?)
func (c *Controller) Ping(ctx *gin.Context) {
	helper.ResponseSuccessPayload(ctx, "pong", time.Now())
}

// Panic simulates a panic
func (c *Controller) Panic(ctx *gin.Context) {
	panic("oh la la!")
}

// NotFound is used as not found handler
func (c *Controller) NotFound(ctx *gin.Context) {
	helper.ResponseNotFound(ctx, "app:route:unknown")
}
