package security

import (
	"github.com/JulesMike/api-er/helper"
	"github.com/JulesMike/api-er/user"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

// TODO: add proper validation
type loginJSON struct {
	Username string `json:"username" binding:"required,min=8"`
	Password string `json:"password" binding:"required,min=8"`
}

// Controller represents the security controller
type Controller struct {
	userRepo *user.Repository
	userSvc  *user.Service
}

// NewController returns a new Security
func NewController(userRepo *user.Repository, userSvc *user.Service) *Controller {
	return &Controller{userRepo: userRepo, userSvc: userSvc}
}

// AttachRoutes attaches the controller's routes to gin.RouterGroup
func (c *Controller) AttachRoutes(r *gin.RouterGroup) {
	r.GET("/csrf-token", c.Token)
	r.POST("/login", c.Login)
	r.POST("/logout", c.Logout)
	r.GET("/me", c.Me)
	r.GET("/status", c.Status)
}

// TokenMismatch is used for ErrorFunc of CSRF middleware
func (c *Controller) TokenMismatch(ctx *gin.Context) {
	helper.ResponseBadRequest(ctx, "csrftoken:invalid")
}

// Token replies with a CSRF token
func (c *Controller) Token(ctx *gin.Context) {
	helper.ResponseSuccessPayload(ctx, "csrftoken:retrieved", csrf.GetToken(ctx))
}

// Login authenticates the user
func (c *Controller) Login(ctx *gin.Context) {
	if _, ok := c.userSvc.UserFromContext(ctx); ok {
		helper.ResponseBadRequest(ctx, "auth:alreadyloggedin")
		return
	}

	var json loginJSON

	if err := ctx.ShouldBindJSON(&json); err != nil {
		helper.ResponseBadRequest(ctx, err.Error())
		return
	}

	userModel := &user.Model{}
	userModel.Username = json.Username

	userModel, err := c.userRepo.Get(ctx, userModel)
	if err != nil {
		if err == user.ErrUserNotFound {
			helper.ResponseNotFound(ctx, "auth:user:notfound")
			return
		}

		helper.ResponseInternalServerError(ctx, "auth:user:internalerror")
		return
	}

	if err := c.userSvc.ComparePassword([]byte(userModel.Password), []byte(json.Password)); err != nil {
		helper.ResponseUnauthorized(ctx, err.Error())
		return
	}

	// TODO: when I get the register user done
	// if !userModel.Verified {
	// 	helper.ResponseUnauthorized(ctx, "auth:user:notverified")
	// 	return
	// }

	if err := c.userSvc.SetUserIDSessionContext(ctx, userModel.ID.Hex()); err != nil {
		helper.ResponseInternalServerError(ctx, err.Error())
		return
	}

	helper.ResponseSuccess(ctx, "auth:loggedin")
}

// Logout deauthenticates the user
func (c *Controller) Logout(ctx *gin.Context) {
	if _, ok := c.userSvc.UserFromContext(ctx); !ok {
		helper.ResponseBadRequest(ctx, "auth:notloggedin")
		return
	}

	if err := c.userSvc.DeleteUserIDSessionContext(ctx); err != nil {
		helper.ResponseInternalServerError(ctx, err.Error())
		return
	}

	helper.ResponseSuccess(ctx, "auth:loggedout")
}

// Me returns self user
func (c *Controller) Me(ctx *gin.Context) {
	user, ok := c.userSvc.UserFromContext(ctx)
	if !ok {
		helper.ResponseBadRequest(ctx, "auth:notloggedin")
		return
	}

	helper.ResponseSuccessPayload(ctx, "auth:me", user)
}

// Status confirms if logged in
func (c *Controller) Status(ctx *gin.Context) {
	helper.ResponseSuccess(ctx, "auth:ok")
}
