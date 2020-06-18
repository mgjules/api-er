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
		users.GET("/:id", controller.GetUser)
		users.PATCH("/:id", controller.UpdateUser)
		users.DELETE("/:id", controller.DeleteUser)
	}

	// r.GET("/protected", func(c *gin.Context) {
	// 	c.String(200, csrf.GetToken(c))
	// })

	// r.POST("/protected", func(c *gin.Context) {
	// 	c.String(200, "CSRF token is valid")
	// })
}
