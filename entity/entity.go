package entity

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// Entity contains common columns for all tables.
type Entity struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (e *Entity) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("ID", uuid.NewV4().String())
}
