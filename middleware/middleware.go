package middleware

import "github.com/jinzhu/gorm"

var _db *gorm.DB

// Init inits the middleware
func Init(db *gorm.DB) {
	_db = db
}
