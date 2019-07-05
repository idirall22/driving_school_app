package service

import (
	"github.com/jinzhu/gorm"
)

// Service struct
type Service struct {
	db *gorm.DB
}

// func (s *service) WailsInit(runtime *wails.Runtime) error {
// 	return nil
// }

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
	// 	var dummyStudent = &Student{
	// 		FileNumber:    "0001",
	// 		FirstName:     "idir",
	// 		LastName:      "makhlouf",
	// 		MaidenName:    " ",
	// 		PhoneNumber:   "0557083719",
	// 		Job:           "devloper",
	// 		BirthDay:      time.Now(),
	// 		Gender:        "Male",
	// 		City:          "Oran",
	// 		AddressStreet: "1273 oran street",
	// 		RegistredDate: time.Now(),
	// 		Image:         "image.png",
	// 		Exams: []*Exam{{
	// 			ExamName: "01",
	// 			Comment:  "no comment",
	// 			DateExam: time.Now(),
	// 			Status:   false},
	// 			{
	// 				ExamName: "02",
	// 				Comment:  "no comment 2",
	// 				DateExam: time.Now(),
	// 				Status:   true},
	// 		},
	// 	}
	// 	MainService.CreateStudent(dummyStudent)
}

//CloseService close database connection
func CloseService() {
	MainService.db.Close()

}
