package service

import (
	"github.com/jinzhu/gorm"
)

// Service struct
type Service struct {
	db *gorm.DB
}

// MainService model
var MainService = &Service{}

// CreateDatabaseDirFile create db files
func CreateDatabaseDirFile(directory, fileName string) error {
	if err := createDatabaseDir(directory, 0777); err != nil {
		return err
	}
	if err := createDatabaseFile(directory, fileName, 0777); err != nil {
		return err
	}
	return nil
}

// InitService init the service
func InitService(db *gorm.DB) {
	MainService.db = db
	MainService.db.AutoMigrate(&Student{}, &Exam{}, &ExamList{})
}

//CloseService close database connection
func CloseService() {
	MainService.db.Close()

}
