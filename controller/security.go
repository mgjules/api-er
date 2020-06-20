package controller

import (
	"github.com/JulesMike/api-er/entity"
	"github.com/JulesMike/api-er/helper"
	"github.com/JulesMike/api-er/repository"
	"github.com/JulesMike/api-er/service"
	"github.com/gin-gonic/gin"
)

// TODO: add proper validation
type loginJSON struct {
	Username string `json:"username" binding:"required,min=8"`
	Password string `json:"password" binding:"required,min=8"`
}

// Security represents the security controller
type Security struct {
	userRepo    *repository.User
	securitySvc *service.Security
}

// NewSecurity returns a new Security
func NewSecurity(userRepo *repository.User, securitySvc *service.Security) *Security {
	return &Security{userRepo: userRepo, securitySvc: securitySvc}
}

// TokenMismatch is used for ErrorFunc of CSRF middleware
func (c *Security) TokenMismatch(ctx *gin.Context) {
	helper.ResponseBadRequest(ctx, "csrftoken:invalid")
}

// Token replies with a CSRF token
func (c *Security) Token(ctx *gin.Context) {
	helper.ResponseSuccessPayload(ctx, "csrftoken:retrieved", c.securitySvc.Token(ctx))
}

// Login authenticates the user
func (c *Security) Login(ctx *gin.Context) {
	if _, ok := c.securitySvc.UserFromContext(ctx); ok {
		helper.ResponseBadRequest(ctx, "auth:alreadyloggedin")
		return
	}

	var json loginJSON

	if err := ctx.ShouldBindJSON(&json); err != nil {
		helper.ResponseBadRequest(ctx, err.Error())
		return
	}

	user := &entity.User{}
	user.Username = json.Username

	user, err := c.userRepo.Get(user)
	if err != nil {
		if err == repository.ErrRecordNotFound {
			helper.ResponseNotFound(ctx, "auth:user:notfound")
			return
		}

		helper.ResponseInternalServerError(ctx, "auth:user:internalerror")
		return
	}

	if err := c.securitySvc.ComparePassword([]byte(user.Password), []byte(json.Password)); err != nil {
		helper.ResponseUnauthorized(ctx, err.Error())
		return
	}

	// TODO: when I get the register user done
	// if !user.Verified {
	// 	helper.ResponseUnauthorized(ctx, "auth:user:notverified")
	// 	return
	// }

	if err := c.securitySvc.SetUserIDSessionContext(ctx, user.ID.String()); err != nil {
		helper.ResponseInternalServerError(ctx, err.Error())
		return
	}

	helper.ResponseSuccess(ctx, "auth:loggedin")
}

// Logout deauthenticates the user
func (c *Security) Logout(ctx *gin.Context) {
	if _, ok := c.securitySvc.UserFromContext(ctx); !ok {
		helper.ResponseBadRequest(ctx, "auth:notloggedin")
		return
	}

	if err := c.securitySvc.DeleteUserIDSessionContext(ctx); err != nil {
		helper.ResponseInternalServerError(ctx, err.Error())
		return
	}

	helper.ResponseSuccess(ctx, "auth:loggedout")
}

// Me returns self user
func (c *Security) Me(ctx *gin.Context) {
	user, ok := c.securitySvc.UserFromContext(ctx)
	if !ok {
		helper.ResponseBadRequest(ctx, "auth:notloggedin")
		return
	}

	helper.ResponseSuccessPayload(ctx, "auth:me", user)
}

// Status confirms if logged in
func (c *Security) Status(ctx *gin.Context) {
	helper.ResponseSuccess(ctx, "auth:ok")
}
