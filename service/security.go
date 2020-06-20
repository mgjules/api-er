package service

import (
	"errors"

	"github.com/JulesMike/api-er/entity"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	csrf "github.com/utrack/gin-csrf"
	"golang.org/x/crypto/bcrypt"
)

// Security represents the security service
type Security struct {
	pwdHash string
}

// NewSecurity returns a new Security
func NewSecurity(passwordHash string) *Security {
	return &Security{pwdHash: passwordHash}
}

// HashPassword is a wrapper over bcrypt.GenerateFromPassword
func (s *Security) HashPassword(password []byte) ([]byte, error) {
	saltedPassword := append(password, s.pwdHash...)
	return bcrypt.GenerateFromPassword(saltedPassword, 10)
}

// ComparePassword is a wrapper over bcrypt.CompareHashAndPassword
func (s *Security) ComparePassword(hashedPassword, password []byte) error {
	saltedPassword := append(password, s.pwdHash...)
	return bcrypt.CompareHashAndPassword(hashedPassword, saltedPassword)
}

// IsPasswordHashed checks if password is hashed
func (s *Security) IsPasswordHashed(hashedPassword []byte) bool {
	_, err := bcrypt.Cost(hashedPassword)
	return err == nil
}

// Token returns a token from gin context
func (s *Security) Token(ctx *gin.Context) string {
	return csrf.GetToken(ctx)
}

// UserFromContext returns the user entity from gin.Context
func (s *Security) UserFromContext(ctx *gin.Context) (*entity.User, bool) {
	v, ok := ctx.Get(entity.UserContextKey)
	if !ok {
		return nil, false
	}

	user, ok := v.(*entity.User)
	if !ok {
		return nil, false
	}

	return user, true
}

// UserIDFromSessionContext returns the user ID from gin.Context
func (s *Security) UserIDFromSessionContext(ctx *gin.Context) (uuid.UUID, error) {
	session := sessions.Default(ctx)

	rawUserID := session.Get(entity.UserSessionKey)

	strUserID, ok := rawUserID.(string)
	if !ok {
		return uuid.Nil, errors.New("can't convert session userID to string")
	}

	return uuid.FromString(strUserID)
}

// SetUserIDSessionContext sets user ID in gin.Context
func (s *Security) SetUserIDSessionContext(ctx *gin.Context, userID string) error {
	session := sessions.Default(ctx)

	session.Set(entity.UserSessionKey, userID)

	return session.Save()
}

// DeleteUserIDSessionContext sets user ID in gin.Context
func (s *Security) DeleteUserIDSessionContext(ctx *gin.Context) error {
	session := sessions.Default(ctx)

	session.Delete(entity.UserSessionKey)

	return session.Save()
}
