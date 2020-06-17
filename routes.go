package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func attachRoutes(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"api-er": time.Now(),
		})
	})
}
