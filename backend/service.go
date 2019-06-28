package service

import "github.com/jinzhu/gorm"

// service struct
type service struct {
	db *gorm.DB
}

var mainService = &service{}

// InitService init the service
func InitService(db *gorm.DB) {
	mainService.db = db
}
