package service

import (
	"context"
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
	mainService.db = db
	mainService.db.DropTableIfExists(&Student{}, &Exam{}, &ExamList{})
	mainService.db.AutoMigrate(&Student{}, &Exam{}, &ExamList{})
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
	// mainService.db.Create(&dummyStudent)
	student, err := GetStudent(context.Background(), 1)
	if err != nil {
		t.Error(err)
	}
	if student.FileNumber != dummyStudent.FileNumber {
		t.Error("There is an error the file number does not match")
	}
}

// Test GetStudents
func testGetStudents(t *testing.T) {
	students, err := GetStudents(context.Background(), 10, 0)
	if err != nil {
		t.Error(err)
	}
	if students[0].FirstName != dummyStudent.FirstName {
		t.Error("There is an error the first name does not match")
	}
}

func testCreateStudent(t *testing.T) {
	connectDatabse()
	CreateStudent(context.Background(), dummyStudent)
	if dummyStudent.ID == 0 {
		t.Error("There is an error with created student")
	}
}

func testUpdateStudent(t *testing.T) {
	name := "updated"
	dummyStudent.FirstName = name
	UpdateStudent(context.Background(), dummyStudent)

	student, err := GetStudent(context.Background(), int64(dummyStudent.ID))

	if err != nil {
		t.Error(err)
	}
	if student.FirstName != name {
		t.Error("There is an error with first name")
	}
}

func testDeleteStudent(t *testing.T) {

	DeleteStudent(context.Background(), 1)
	student, err := GetStudent(context.Background(), int64(dummyStudent.ID))

	if err != nil {
		t.Error(err)
	}
	if student.DeletedAt == nil {
		t.Error("There is an error this sudents was deleted")
	}
}
