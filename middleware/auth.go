package middleware

import (
	uuid "github.com/satori/go.uuid"

	"github.com/JulesMike/api-er/entity"
	"github.com/JulesMike/api-er/helper"
	"github.com/JulesMike/api-er/repository"

	"github.com/casbin/casbin"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Auth is the auth middleware
func Auth(e *casbin.Enforcer, userRepo *repository.User) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)

		userRole := entity.GuestUserRole

		if userID := session.Get(entity.UserSessionKey); userID != nil {
			user := &entity.User{}
			user.ID = uuid.FromStringOrNil(userID.(string))
			if user.ID != uuid.Nil {
				if user, err := userRepo.Get(user); err == nil {
					userRole = user.Role

					ctx.Set(entity.UserContextKey, user)
				}
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
