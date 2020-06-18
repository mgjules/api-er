package main

import (
	"github.com/JulesMike/api-er/controller"
	"github.com/gin-gonic/gin"
)

func attachRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/ping", controller.Ping)

		api.GET("/panic", controller.Panic)

		api.GET("/token", controller.Token)

		// Security routes
		api.POST("/login", controller.Login)
		api.POST("/logout", controller.Logout)
		api.GET("/me", controller.Me)

		// Users routes
		users := api.Group("/users")
		{
			users.POST("/", controller.CreateUser)
			users.GET("/", controller.ListUsers)
			users.GET("/:id", controller.GetUser)
			users.PATCH("/:id", controller.UpdateUser)
			users.DELETE("/:id", controller.DeleteUser)
		}
	}

	r.NoRoute(controller.NotFound)
}
