package user

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/JulesMike/api-er/helper"
	"github.com/gin-gonic/gin"
)

// TODO: add proper validation
type reqJSON struct {
	Username string `json:"username" binding:"required,min=8"`
	Password string `json:"password" binding:"required,min=8"`
	Email    string `json:"email" binding:"required,email"`
	Verified bool   `json:"verified" binding:""`
}

// Controller represents the user controller
type Controller struct {
	userRepo *Repository
}

// NewController returns a new User
func NewController(userRepo *Repository) *Controller {
	return &Controller{userRepo: userRepo}
}

// AttachRoutes attaches the controller's routes to gin.RouterGroup
func (c *Controller) AttachRoutes(r *gin.RouterGroup) {
	users := r.Group("/users")
	{
		users.POST("/", c.Create)
		users.GET("/", c.List)
		users.GET("/:id", c.Get)
		users.PATCH("/:id", c.Update)
		users.DELETE("/:id", c.Delete)
	}
}

// Create creates a new user
func (c *Controller) Create(ctx *gin.Context) {
	var json reqJSON

	if err := ctx.ShouldBindJSON(&json); err != nil {
		helper.ResponseBadRequest(ctx, err.Error())
		return
	}

	user := &Model{
		Username: json.Username,
		Password: json.Password,
		Email:    json.Email,
		Verified: json.Verified,
	}

	user, err := c.userRepo.Create(ctx, user)
	if err != nil {
		helper.ResponseInternalServerError(ctx, err.Error())
		return
	}

	helper.ResponseSuccessPayload(ctx, "user:created", user)
}

// Get retrieves a single user
func (c *Controller) Get(ctx *gin.Context) {
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		helper.ResponseBadRequest(ctx, err.Error())
		return
	}

	user := &Model{}
	user.ID = id

	user, err = c.userRepo.Get(ctx, user)
	if err != nil {
		if err == ErrUserNotFound {
			helper.ResponseNotFound(ctx, "user:notfound")
			return
		}

		helper.ResponseInternalServerError(ctx, "user:internalerror")
		return
	}

	helper.ResponseSuccessPayload(ctx, "user:retrieved", user)
}

// List retrieves a list of user
func (c *Controller) List(ctx *gin.Context) {
	users, err := c.userRepo.List(ctx, bson.M{})
	if err != nil {
		helper.ResponseInternalServerError(ctx, err.Error())
		return
	}

	helper.ResponseSuccessPayload(ctx, "users:retrieved", users)
}

// Update updates a user
func (c *Controller) Update(ctx *gin.Context) {
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		helper.ResponseBadRequest(ctx, err.Error())
		return
	}

	var json reqJSON

	if err := ctx.ShouldBindJSON(&json); err != nil {
		helper.ResponseBadRequest(ctx, err.Error())
		return
	}

	user := &Model{
		ID:       id,
		Username: json.Username,
		Password: json.Password,
		Email:    json.Email,
		Verified: json.Verified,
	}

	user, err = c.userRepo.Update(ctx, user)
	if err != nil {
		if err == ErrUserNotFound {
			helper.ResponseNotFound(ctx, "user:notfound")
			return
		}

		helper.ResponseInternalServerError(ctx, err.Error())
		return
	}

	helper.ResponseSuccessPayload(ctx, "user:updated", user)
}

// Delete deletes a user
func (c *Controller) Delete(ctx *gin.Context) {
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		helper.ResponseBadRequest(ctx, err.Error())
		return
	}

	user := &Model{}
	user.ID = id

	user, err = c.userRepo.Delete(ctx, user)
	if err != nil {
		if err == ErrUserNotFound {
			helper.ResponseNotFound(ctx, "user:notfound")
			return
		}

		helper.ResponseInternalServerError(ctx, err.Error())
		return
	}

	helper.ResponseSuccessPayload(ctx, "user:deleted", user)
}
