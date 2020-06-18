package controller

import (
	"github.com/jinzhu/gorm"
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
		return
	}

	user := entity.User{Username: json.Username, Password: json.Password}

	if err := _db.Create(&user).Error; err != nil {
		helper.ResponseBadRequest(c, err.Error())
		return
	}

	helper.ResponseSuccess(c, "User created")
}

// GetUser retrieves a single user
func GetUser(c *gin.Context) {
	id, err := uuid.FromString(c.Param("id"))
	if err != nil {
		helper.ResponseBadRequest(c, err.Error())
		return
	}

	var user entity.User
	user.ID = id

	if err := _db.First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			helper.ResponseNotFound(c, err.Error())
			return
		}

		helper.ResponseInternalServerError(c, err.Error())
		return
	}

	helper.ResponseSuccessPayload(c, "User retrieved", user)
}

// ListUsers retrieves a list of user
func ListUsers(c *gin.Context) {
	users := []entity.User{}

	if err := _db.Find(&users).Error; err != nil {
		helper.ResponseInternalServerError(c, err.Error())
		return
	}

	helper.ResponseSuccessPayload(c, "Users retrieved", users)
}

// UpdateUser updates a user
func UpdateUser(c *gin.Context) {
	id, err := uuid.FromString(c.Param("id"))
	if err != nil {
		helper.ResponseBadRequest(c, err.Error())
		return
	}

	var json updateUser

	if err := c.ShouldBindJSON(&json); err != nil {
		helper.ResponseBadRequest(c, err.Error())
		return
	}

	var user entity.User
	user.ID = id

	if err := _db.First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			helper.ResponseNotFound(c, err.Error())
			return
		}

		helper.ResponseInternalServerError(c, err.Error())
		return
	}

	if err := _db.Model(&user).Updates(json).Error; err != nil {
		helper.ResponseInternalServerError(c, err.Error())
		return
	}

	helper.ResponseSuccessPayload(c, "User updated", user)
}

// DeleteUser deletes a user
func DeleteUser(c *gin.Context) {
	id, err := uuid.FromString(c.Param("id"))
	if err != nil {
		helper.ResponseBadRequest(c, err.Error())
		return
	}

	var user entity.User
	user.ID = id

	if err := _db.Delete(&user).Error; err != nil {
		helper.ResponseInternalServerError(c, err.Error())
		return
	}

	helper.ResponseSuccess(c, "User deleted")
}
