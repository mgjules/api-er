package controller

import (
	"net/http"

	"github.com/JulesMike/api-er/entity"
	"github.com/gin-gonic/gin"
)

type createUser struct {
	Username string `json:"username" binding:"required,min=8"`
	Password string `json:"password" binding:"required,min=8"`
}

// CreateUser creates a new user
func CreateUser(c *gin.Context) {
	var json createUser

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := entity.User{Username: json.Username, Password: json.Password}

	if err := _db.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "user created",
	})
}

// ListUsers creates a new user
func ListUsers(c *gin.Context) {
	users := []entity.User{}

	if err := _db.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}
