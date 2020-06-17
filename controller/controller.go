package controller

import "github.com/jinzhu/gorm"

var _db *gorm.DB

// Init inits the controller
func Init(db *gorm.DB) {
	_db = db
}
