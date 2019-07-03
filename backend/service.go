package service

import "github.com/jinzhu/gorm"

// service struct
type service struct {
	db *gorm.DB
}

// MainService model
var MainService = &service{}

// InitService init the service
func InitService(db *gorm.DB) {
	MainService.db = db
}

//CloseService close database connection
func CloseService() {
	MainService.db.Close()
}
