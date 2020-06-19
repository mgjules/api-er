package controller

import (
	uuid "github.com/satori/go.uuid"

	"github.com/JulesMike/api-er/entity"
	"github.com/JulesMike/api-er/helper"
	"github.com/JulesMike/api-er/repository"
	"github.com/gin-gonic/gin"
)

// TODO: add proper validation
type userJSON struct {
	Username string `json:"username" binding:"required,min=8"`
	Password string `json:"password" binding:"required,min=8"`
	Email    string `json:"email" binding:"required,email"`
	Verified bool   `json:"verified" binding:""`
}

// User represents the user controller
type User struct {
	userRepo *repository.User
}

// NewUser returns a new User
func NewUser(userRepo *repository.User) *User {
	return &User{userRepo: userRepo}
}

// Create creates a new user
func (c *User) Create(ctx *gin.Context) {
	var json userJSON

	if err := ctx.ShouldBindJSON(&json); err != nil {
		helper.ResponseBadRequest(ctx, err.Error())
		return
	}

	user := &entity.User{
		Username: json.Username,
		Password: json.Password,
		Email:    json.Email,
		Verified: json.Verified,
	}

	user, err := c.userRepo.Create(user)
	if err != nil {
		helper.ResponseInternalServerError(ctx, err.Error())
		return
	}

	helper.ResponseSuccessPayload(ctx, "user:created", user)
}

// Get retrieves a single user
func (c *User) Get(ctx *gin.Context) {
	id, err := uuid.FromString(ctx.Param("id"))
	if err != nil {
		helper.ResponseBadRequest(ctx, err.Error())
		return
	}

	user := &entity.User{}
	user.ID = id

	user, err = c.userRepo.Get(user)
	if err != nil {
		if err == repository.ErrRecordNotFound {
			helper.ResponseNotFound(ctx, "user:notfound")
			return
		}

		helper.ResponseInternalServerError(ctx, "user:internalerror")
		return
	}

	helper.ResponseSuccessPayload(ctx, "user:retrieved", user)
}

// List retrieves a list of user
func (c *User) List(ctx *gin.Context) {
	users, err := c.userRepo.List()
	if err != nil {
		helper.ResponseInternalServerError(ctx, err.Error())
		return
	}

	helper.ResponseSuccessPayload(ctx, "users:retrieved", users)
}

// Update updates a user
func (c *User) Update(ctx *gin.Context) {
	id, err := uuid.FromString(ctx.Param("id"))
	if err != nil {
		helper.ResponseBadRequest(ctx, err.Error())
		return
	}

	var json userJSON

	if err := ctx.ShouldBindJSON(&json); err != nil {
		helper.ResponseBadRequest(ctx, err.Error())
		return
	}

	user := &entity.User{}
	user.ID = id

	user, err = c.userRepo.Update(user, json)
	if err != nil {
		if err == repository.ErrRecordNotFound {
			helper.ResponseNotFound(ctx, "user:notfound")
			return
		}

		helper.ResponseInternalServerError(ctx, err.Error())
		return
	}

	helper.ResponseSuccessPayload(ctx, "user:updated", user)
}

// Delete deletes a user
func (c *User) Delete(ctx *gin.Context) {
	id, err := uuid.FromString(ctx.Param("id"))
	if err != nil {
		helper.ResponseBadRequest(ctx, err.Error())
		return
	}

	user := &entity.User{}
	user.ID = id

	user, err = c.userRepo.Delete(user)
	if err != nil {
		if err == repository.ErrRecordNotFound {
			helper.ResponseNotFound(ctx, "user:notfound")
			return
		}

		helper.ResponseInternalServerError(ctx, err.Error())
		return
	}

	helper.ResponseSuccessPayload(ctx, "user:deleted", user)
}
