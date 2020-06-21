package user

import (
	"log"

	"github.com/JulesMike/api-er/helper"

	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"
)

// Auth is the auth middleware
func Auth(e *casbin.Enforcer, userSvc *Service, userRepo *Repository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userRole := GuestRole

		if userID, err := userSvc.UserIDFromSessionContext(ctx); err == nil {
			user := &Model{}
			user.ID = userID
			if user, err := userRepo.Get(ctx, user); err == nil {
				userRole = user.Role

				ctx.Set(ContextKey, user)
			} else {
				log.Println(err)
			}
		} else {
			log.Println(err)
		}

		r := ctx.Request

		if !e.Enforce(userRole, r.RequestURI, r.Method) {
			helper.ResponseUnauthorized(ctx, "auth:notallowed")
			return
		}

		ctx.Next()
	}
}
