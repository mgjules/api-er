package middleware

import (
	"github.com/JulesMike/api-er/service"

	"github.com/JulesMike/api-er/entity"
	"github.com/JulesMike/api-er/helper"
	"github.com/JulesMike/api-er/repository"

	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"
)

// Auth is the auth middleware
func Auth(e *casbin.Enforcer, securitySvc *service.Security, userRepo *repository.User) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userRole := entity.GuestUserRole

		userID, err := securitySvc.UserIDFromSessionContext(ctx)
		if err == nil {
			user := &entity.User{}
			user.ID = userID
			if user, err := userRepo.Get(user); err == nil {
				userRole = user.Role

				ctx.Set(entity.UserContextKey, user)
			}
		}

		r := ctx.Request

		if !e.Enforce(userRole, r.RequestURI, r.Method) {
			helper.ResponseUnauthorized(ctx, "auth:notallowed")
			return
		}

		ctx.Next()
	}
}
