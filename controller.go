package main

import "github.com/gin-gonic/gin"

// Controller defines a generic controller
type Controller interface {
	AttachRoutes(r *gin.RouterGroup)
}
