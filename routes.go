package main

import (
	"github.com/JulesMike/api-er/controller"
	"github.com/gin-gonic/gin"
)

func attachRoutes(r *gin.Engine) {
	r.GET("/ping", controller.Ping)

	r.GET("/panic", controller.Panic)

	r.GET("/token", controller.Token)

	// Security routes
	r.POST("/login", controller.Login)
	r.POST("/logout", controller.Logout)

	// Users routes
	users := r.Group("/users")
	{
		users.POST("/", controller.CreateUser)
		users.GET("/", controller.ListUsers)
		users.GET("/:id", controller.GetUser)
		users.PATCH("/:id", controller.UpdateUser)
		users.DELETE("/:id", controller.DeleteUser)
	}
}
