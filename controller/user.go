package controller

import (
	"net/http"

	uuid "github.com/satori/go.uuid"

	"github.com/JulesMike/api-er/entity"
	"github.com/gin-gonic/gin"
)

type createUser struct {
	Username string `json:"username" binding:"required,min=8"` // TODO: add validation
	Password string `json:"password" binding:"required,min=8"` // TODO: add validation
}

type updateUser struct {
	Username string `json:"username" binding:""` // TODO: add validation
	Password string `json:"password" binding:""` // TODO: add validation
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

// UpdateUser updates a user
func UpdateUser(c *gin.Context) {
	id, err := uuid.FromString(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var json updateUser

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user entity.User
	user.ID = id

	if err := _db.First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := _db.Model(&user).Updates(json).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}
