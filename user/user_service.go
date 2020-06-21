package user

import (
	"errors"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// Service represents a user service
type Service struct {
	pwdHash string
}

// NewService returns a new user service
func NewService(passwordHash string) *Service {
	return &Service{pwdHash: passwordHash}
}

// HashPassword is a wrapper over bcrypt.GenerateFromPassword
func (s *Service) HashPassword(password []byte) ([]byte, error) {
	saltedPassword := append(password, s.pwdHash...)
	return bcrypt.GenerateFromPassword(saltedPassword, 10)
}

// ComparePassword is a wrapper over bcrypt.CompareHashAndPassword
func (s *Service) ComparePassword(hashedPassword, password []byte) error {
	saltedPassword := append(password, s.pwdHash...)
	return bcrypt.CompareHashAndPassword(hashedPassword, saltedPassword)
}

// IsPasswordHashed checks if password is hashed
func (s *Service) IsPasswordHashed(hashedPassword []byte) bool {
	_, err := bcrypt.Cost(hashedPassword)
	return err == nil
}

// UserFromContext returns the user entity from gin.Context
func (s *Service) UserFromContext(ctx *gin.Context) (*Model, bool) {
	v, ok := ctx.Get(ContextKey)
	if !ok {
		return nil, false
	}

	user, ok := v.(*Model)
	if !ok {
		return nil, false
	}

	return user, true
}

// UserIDFromSessionContext returns the user ID from gin.Context
func (s *Service) UserIDFromSessionContext(ctx *gin.Context) (primitive.ObjectID, error) {
	session := sessions.Default(ctx)

	strUserID, ok := session.Get(SessionKey).(string)
	if !ok {
		return primitive.NilObjectID, errors.New("can't convert session userIDs to string")
	}

	return primitive.ObjectIDFromHex(strUserID)
}

// SetUserIDSessionContext sets user ID in gin.Context
func (s *Service) SetUserIDSessionContext(ctx *gin.Context, userID string) error {
	session := sessions.Default(ctx)

	session.Set(SessionKey, userID)

	return session.Save()
}

// DeleteUserIDSessionContext sets user ID in gin.Context
func (s *Service) DeleteUserIDSessionContext(ctx *gin.Context) error {
	session := sessions.Default(ctx)

	session.Delete(SessionKey)

	return session.Save()
}
