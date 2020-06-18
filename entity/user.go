package entity

import (
	"errors"

	"github.com/JulesMike/api-er/security"
	"github.com/jinzhu/gorm"
)

// User Roles
const (
	AdminUserRole  = ":admin:"
	NormalUserRole = ":normal:"
	GuestUserRole  = ":guest:"
)

// UserSessionKey is the user session key
const UserSessionKey = "user"

// UserContextKey is the user context key
const UserContextKey = "userctx"

// User entity holds user information
type User struct {
	Entity
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Email    string `gorm:"not null"`
	Verified bool   `gorm:"not null"`
	Role     string `gorm:"default:':normal:'"`
}

// IsAdmin checks if user belongs to AdminUserRole
func (u *User) IsAdmin() bool {
	return u.Role == AdminUserRole
}

// IsNormal checks if user belongs to NormalUserRole
func (u *User) IsNormal() bool {
	return u.Role == NormalUserRole
}

// BeforeSave gorm hook
func (u *User) BeforeSave(scope *gorm.Scope) (err error) {
	if security.IsPasswordHashed([]byte(u.Password)) {
		return
	}

	hashedPassword, err := security.HashPassword([]byte(u.Password))
	if err != nil {
		err = errors.New("Can't hash password")
		return
	}

	scope.SetColumn("Password", string(hashedPassword))

	return
}

// AfterFind gorm hook
func (u *User) AfterFind() (err error) {
	if u.Role == "" {
		u.Role = NormalUserRole
	}

	return
}
