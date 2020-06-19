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
	helper.ResponseBadRequest(c, "csrftoken:invalid")
}

// Token replies with a CSRF token
func Token(c *gin.Context) {
	token := csrf.GetToken(c)
	helper.ResponseSuccessPayload(c, "csrftoken:retrieved", token)
}

// Login authenticates the user
func Login(c *gin.Context) {
	if _, ok := helper.UserFromContext(c); ok {
		helper.ResponseBadRequest(c, "auth:alreadyloggedin")
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

	session := sessions.Default(c)

	session.Set(entity.UserSessionKey, user.ID.String())
	if err := session.Save(); err != nil {
		helper.ResponseInternalServerError(c, err.Error())
		return
	}

	helper.ResponseSuccess(c, "auth:loggedin")
}

// Logout deauthenticates the user
func Logout(c *gin.Context) {
	if _, ok := helper.UserFromContext(c); !ok {
		helper.ResponseBadRequest(c, "auth:notloggedin")
		return
	}

	session := sessions.Default(c)

	session.Delete(entity.UserSessionKey)
	if err := session.Save(); err != nil {
		helper.ResponseInternalServerError(c, err.Error())
		return
	}

	helper.ResponseSuccess(c, "auth:loggedout")
}

// Me returns self user
func Me(c *gin.Context) {
	user, ok := helper.UserFromContext(c)
	if !ok {
		helper.ResponseBadRequest(c, "auth:notloggedin")
		return
	}

	helper.ResponseSuccessPayload(c, "auth:me", user)
}

// Status confirms if logged in
func Status(c *gin.Context) {
	helper.ResponseSuccess(c, "auth:ok")
}
