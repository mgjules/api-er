package main

import (
	"github.com/JulesMike/api-er/controller"
	"github.com/gin-gonic/gin"
)

func attachRoutes(r *gin.Engine) {
	r.GET("/ping", controller.Ping)

	r.GET("/panic", controller.Panic)

	r.GET("/token", controller.Token)

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
