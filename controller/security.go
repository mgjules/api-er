package controller

import (
	"github.com/JulesMike/api-er/entity"
	"github.com/JulesMike/api-er/helper"
	"github.com/JulesMike/api-er/security"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	csrf "github.com/utrack/gin-csrf"
)

type login struct {
	Username string `json:"username" binding:"required,min=8"` // TODO: add validation
	Password string `json:"password" binding:"required,min=8"` // TODO: add validation
}

// TokenMismatch is used for ErrorFunc of CSRF middleware
func TokenMismatch(c *gin.Context) {
	helper.ResponseBadRequest(c, "CSRF Token mismatch")
}

// Token replies with a CSRF token
func Token(c *gin.Context) {
	token := csrf.GetToken(c)
	helper.ResponseSuccessPayload(c, "CSRF Token", token)
}

// Login authenticates the user
func Login(c *gin.Context) {
	session := sessions.Default(c)

	if session.Get(entity.UserSessionKey) != nil {
		helper.ResponseBadRequest(c, "You are already logged in!")
		return
	}

	var json login

	if err := c.ShouldBindJSON(&json); err != nil {
		helper.ResponseBadRequest(c, err.Error())
		return
	}

	var user entity.User
	user.Username = json.Username

	if err := _db.Where(&user).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			helper.ResponseNotFound(c, err.Error())
			return
		}

		helper.ResponseInternalServerError(c, err.Error())
		return
	}

	if err := security.ComparePassword([]byte(user.Password), []byte(json.Password)); err != nil {
		helper.ResponseUnauthorized(c, err.Error())
		return
	}

	// if !user.Verified {
	// 	helper.ResponseUnauthorized(c, "User not verified")
	// 	return
	// }

	session.Set(entity.UserSessionKey, user.ID.String())
	if err := session.Save(); err != nil {
		helper.ResponseInternalServerError(c, err.Error())
		return
	}

	helper.ResponseSuccess(c, "Login successful")
}

// Logout deauthenticates the user
func Logout(c *gin.Context) {
	session := sessions.Default(c)

	if session.Get(entity.UserSessionKey) == nil {
		helper.ResponseBadRequest(c, "Invalid session token")
		return
	}

	session.Delete(entity.UserSessionKey)
	if err := session.Save(); err != nil {
		helper.ResponseInternalServerError(c, err.Error())
		return
	}

	helper.ResponseSuccess(c, "Logout successful")
}
