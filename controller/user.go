package controller

import (
	uuid "github.com/satori/go.uuid"

	"github.com/JulesMike/api-er/entity"
	"github.com/JulesMike/api-er/helper"
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
		helper.ResponseBadRequest(c, err.Error())
	}

	user := entity.User{Username: json.Username, Password: json.Password}

	if err := _db.Create(&user).Error; err != nil {
		helper.ResponseBadRequest(c, err.Error())
	}

	helper.ResponseSuccess(c, "User created")
}

// ListUsers creates a new user
func ListUsers(c *gin.Context) {
	users := []entity.User{}

	if err := _db.Find(&users).Error; err != nil {
		helper.ResponseInternalServerError(c, err.Error())
	}

	helper.ResponseSuccessPayload(c, "Users retrieved", users)
}

// UpdateUser updates a user
func UpdateUser(c *gin.Context) {
	id, err := uuid.FromString(c.Param("id"))
	if err != nil {
		helper.ResponseBadRequest(c, err.Error())
	}

	var json updateUser

	if err := c.ShouldBindJSON(&json); err != nil {
		helper.ResponseBadRequest(c, err.Error())
	}

	var user entity.User
	user.ID = id

	if err := _db.First(&user).Error; err != nil {
		helper.ResponseNotFound(c, err.Error())
	}

	if err := _db.Model(&user).Updates(json).Error; err != nil {
		helper.ResponseInternalServerError(c, err.Error())
	}

	helper.ResponseSuccessPayload(c, "User updated", user)
}
