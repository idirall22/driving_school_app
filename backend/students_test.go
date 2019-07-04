package service

import (
	"log"
	"testing"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func connectDatabse() {
	db, err := gorm.Open("sqlite3", "db/db.sqlite3")
	if err != nil {
		log.Println("Error to connect database")
		log.Fatal(err)
	}
	MainService.db = db
	MainService.db.DropTableIfExists(&Student{}, &Exam{}, &ExamList{})
	MainService.db.AutoMigrate(&Student{}, &Exam{}, &ExamList{})
}

var dummyStudent = &Student{
	FileNumber:    "0001",
	FirstName:     "idir",
	LastName:      "makhlouf",
	MaidenName:    " ",
	PhoneNumber:   "0557083719",
	Job:           "devloper",
	BirthDay:      time.Now(),
	Gender:        "Male",
	City:          "Oran",
	AddressStreet: "1273 oran street",
	RegistredDate: time.Now(),
	Image:         "image.png",
	Exams: []*Exam{{
		ExamName: "01",
		Examiner: "examiner",
		Comment:  "no comment",
		DateExam: time.Now(),
		Status:   false},
		{
			ExamName: "02",
			Examiner: "examiner2",
			Comment:  "no comment 2",
			DateExam: time.Now(),
			Status:   true},
	},
}

// Test GetStudent
func testGetStudent(t *testing.T) {
	// MainService.db.Create(&dummyStudent)
	student := MainService.GetStudent(1)
	if student.FileNumber != dummyStudent.FileNumber {
		t.Error("There is an error the file number does not match")
	}
}

// Test GetStudents
func testGetStudents(t *testing.T) {
	students := MainService.GetStudents(10, 0)
	if students[0].FirstName != dummyStudent.FirstName {
		t.Error("There is an error the first name does not match")
	}
}

func testCreateStudent(t *testing.T) {
	connectDatabse()
	MainService.CreateStudent(dummyStudent)
	if dummyStudent.ID == 0 {
		t.Error("There is an error with created student")
	}
}

func testUpdateStudent(t *testing.T) {
	name := "updated"
	dummyStudent.FirstName = name
	MainService.UpdateStudent(dummyStudent)

	student := MainService.GetStudent(int64(dummyStudent.ID))

	if student.FirstName != name {
		t.Error("There is an error with first name")
	}
}

func testDeleteStudent(t *testing.T) {

	MainService.DeleteStudent(1)

	s := MainService.GetStudents(10, 0)

	if len(s) != 0 {
		t.Error("There is an error, list should to be equal to 0")
	}
}
