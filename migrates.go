package main

import (
	"github.com/JulesMike/api-er/entity"
	"github.com/jinzhu/gorm"
)

func autoMigrate(db *gorm.DB) {
	db.AutoMigrate(&entity.User{})
}
