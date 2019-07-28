package service

import (
	"encoding/json"
	"log"
	"strconv"
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
	MainService.db.Model(&Student{})
}

// Number of student created on the tests
var testStudentsCount = 5

var dummyStudent = &Student{
	FileNumber:    "0001",
	FirstName:     "idir",
	LastName:      `"fr":"makhlouf", "ar": "مخلوف"`,
	MaidenName:    "none",
	PhoneNumber:   "05-57-08-37-19",
	Job:           "devloper",
	BirthDay:      time.Now(),
	Country:       "Algeria",
	Department:    "Oran",
	Gender:        "Male",
	City:          "Oran",
	AddressStreet: "1273 oran street",
	RegistredDate: time.Now(),
}

// Test GetStudent
func testGetStudent(t *testing.T) {
	getStudentInfos, err := MainService.GetStudent(1, "", "")
	if err != nil {
		log.Fatal(err)
	}
	if getStudentInfos.Student.ID != 1 {
		t.Error("There is an error the ids does not match")
	}
}

// Test GetStudents
func testGetStudents(t *testing.T) {
	studentsListOut, _ := MainService.GetStudents("", "", 10, 0)
	if studentsListOut.Students[0].FirstName != dummyStudent.FirstName {
		t.Error("There is an error the first name does not match")
	}

	studentsListOut, _ = MainService.GetStudents(dummyStudent.LastName, "", 10, 0)
	if len(studentsListOut.Students) != testStudentsCount {
		t.Error("There is an error the first name does not match")
	}
}

// Test create student
func testCreateStudent(t *testing.T) {
	// Connect to database test
	connectDatabse()

	m := make(map[string]interface{})
	data, err := json.Marshal(dummyStudent)
	if err != nil {
		t.Error(err)
	}

	if err := json.Unmarshal(data, &m); err != nil {
		t.Error(err)
	}

	for i := 1; i <= testStudentsCount; i++ {
		m["file_number"] = "00" + strconv.Itoa(i)
		id, err := MainService.CreateStudent(m)
		if err != nil {
			t.Error("There is an error student was not created")
		}
		if dummyStudent != nil {
			dummyStudent.ID = id
			if id != uint(i) {
				t.Error("There is an error, ids does not match")
			}
		}
	}
}

// Test update student
func testUpdateStudent(t *testing.T) {
	m := make(map[string]interface{})
	data, err := json.Marshal(dummyStudent)
	if err != nil {
		t.Error(err)
	}
	if err := json.Unmarshal(data, &m); err != nil {
		t.Error(err)
	}
	m["id"] = 1
	m["first_name"] = "updated"
	id, err := MainService.UpdateStudent(m)
	if err != nil {
		t.Error("There is an error the student was not updated")
	}
	getStudentInfos, err := MainService.GetStudent(1, "", "")
	if err != nil {
		t.Error("There is an error can not get student")
	}

	if getStudentInfos.Student != nil {
		if getStudentInfos.Student.FirstName != m["first_name"] && id != 1 {
			t.Error("There is an error with the student who was updated")
		}
	}
}

// Test delete student
func testDeleteStudent(t *testing.T) {
	for i := 1; i <= testStudentsCount; i++ {
		err := MainService.DeleteStudent(uint(i))
		if err != nil {
			t.Error(err)
		}
	}
	studentsListOut, err := MainService.GetStudents("", "", 10, 0)

	if err != nil {
		t.Error(err)
	}
	if len(studentsListOut.Students) != 0 {
		t.Errorf("There is an error the length should be %d", testStudentsCount-1)
	}
}
