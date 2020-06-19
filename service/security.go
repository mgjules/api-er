package service

import (
	"github.com/JulesMike/api-er/entity"
	"github.com/gin-gonic/gin"
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
