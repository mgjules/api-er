package middleware

import (
	uuid "github.com/satori/go.uuid"

	"github.com/JulesMike/api-er/entity"
	"github.com/JulesMike/api-er/helper"

	"github.com/casbin/casbin"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Auth is the auth middleware
func Auth(e *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)

		userRole := entity.GuestUserRole

		if userID := session.Get(entity.UserSessionKey); userID != nil {
			var user entity.User
			user.ID = uuid.FromStringOrNil(userID.(string))
			if user.ID != uuid.Nil {
				if err := _db.Select("role").First(&user).Error; err == nil {
					userRole = user.Role
				}
			}
		}

		r := c.Request

		if !e.Enforce(userRole, r.RequestURI, r.Method) {
			helper.ResponseUnauthorized(c, "You are not allowed!")
			return
		}

		c.Next()
	}
}