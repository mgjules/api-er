package entity

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
