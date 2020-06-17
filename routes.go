package main

import (
	"net/http"
	"time"

	"github.com/JulesMike/api-er/controller"
	"github.com/gin-gonic/gin"
)

func attachRoutes(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"api-er": time.Now(),
		})
	})

	r.GET("/panic", func(c *gin.Context) {
		panic("oh la la!")
	})

	// Users routes
	users := r.Group("/users")
	{
		users.POST("/", controller.CreateUser)
		users.GET("/", controller.ListUsers)
	}
}
