package service

import (
	"log"

	"github.com/jinzhu/gorm"
)

// Service struct
type Service struct {
	db *gorm.DB
}

// test mode
var testMode = true

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
	if testMode == true {
		MainService.db.DropTableIfExists(&Student{}, &Exam{}, &ExamList{})
	}
	MainService.db.AutoMigrate(&Student{}, &Exam{}, &ExamList{})

	if testMode == true {
		students := openJSONStudentFile("backend/students_test.json")

		for i := 0; i < len(students); i++ {
			_, err := MainService.CreateStudent(students[i])
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

//CloseService close database connection
func CloseService() {
	MainService.db.Close()

}
