package main

import (
	"github.com/JulesMike/api-er/controller"
	"github.com/gin-gonic/gin"
)

func attachRoutes(
	r *gin.Engine,
	appCtrl *controller.App,
	securityCtrl *controller.Security,
	userCtrl *controller.User,
) {
	api := r.Group("/api")
	{
		// App routes
		api.GET("/ping", appCtrl.Ping)

		api.GET("/panic", appCtrl.Panic)

		// Security routes
		api.GET("/token", securityCtrl.Token)
		api.POST("/login", securityCtrl.Login)
		api.POST("/logout", securityCtrl.Logout)
		api.GET("/me", securityCtrl.Me)
		api.GET("/status", securityCtrl.Status)

		// Users routes
		users := api.Group("/users")
		{
			users.POST("/", userCtrl.Create)
			users.GET("/", userCtrl.List)
			users.GET("/:id", userCtrl.Get)
			users.PATCH("/:id", userCtrl.Update)
			users.DELETE("/:id", userCtrl.Delete)
		}
	}

	r.NoRoute(appCtrl.NotFound)
}
