package user

import "go.mongodb.org/mongo-driver/bson/primitive"

// User Roles
const (
	AdminRole  = ":admin:"
	NormalRole = ":normal:"
	GuestRole  = ":guest:"
)

// SessionKey is the user session key
const SessionKey = ":user:key:session"

// ContextKey is the user context key
const ContextKey = ":user:key:context"

// Model represents a single user date structure
type Model struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username,omitempty"`
	Password string             `bson:"password,omitempty"`
	Email    string             `bson:"email,omitempty"`
	Verified bool               `bson:"verified,omitempty"`
	Role     string             `bson:"role,omitempty"`
}

// SetDefaults sets default values
func (m *Model) SetDefaults() {
	if m.Role == "" {
		m.Role = NormalRole
	}
}
