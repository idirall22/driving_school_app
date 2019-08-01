package service

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"testing"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type m map[string]interface{}
type arrayMap []m

var students = arrayMap{}

func openJSONStudentFile() arrayMap {
	fileName := "students_test.json"
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Printf("Could not open %s file.\n", fileName)
		os.Exit(1)
	}
	a := arrayMap{}
	if err := json.Unmarshal(data, &a); err != nil {
		log.Fatal(err)
	}
	return a
}

func connectDatabse() {
	students = openJSONStudentFile()

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
	LastName:      `"fr":"lastname", "ar": "lastname"`,
	MaidenName:    "none",
	PhoneNumber:   "0000000000",
	Job:           "devloper",
	BirthDay:      time.Now(),
	Country:       "Algeria",
	Department:    "Oran",
	Gender:        "Male",
	City:          "Oran",
	AddressStreet: "1273 oran street",
	RegistredDate: time.Now(),
}

////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////
// Test GetStudent /////////////////////////////////////////
func testGetStudent(t *testing.T) {

	dataTest := []struct {
		id          uint
		lastName    string
		phoneNumber string
	}{
		// Test 0 when we provaid id only
		{1, "", ""},
		// Test 1 when we provaid lastName
		{0, "polo", ""},
		// Test 2 when we provaid lastName ans phoneNumber the priority is
		// for lastName
		{0, "polo", "0000000000"},
		// Test 3 when we provaid phoneNumber
		{0, "", "0000000000"},
	}
	for i := 0; i < len(dataTest); i++ {
		getStudentInfos, err := MainService.
			GetStudent(dataTest[i].id, dataTest[i].lastName, dataTest[i].phoneNumber)
		if err != nil {
			log.Fatal(err)
			switch i {
			case 0:
				if getStudentInfos.Student.ID != dataTest[i].id {
					t.Errorf("There is an error, the ids does not match should found %d but got %d",
						dataTest[i].id, getStudentInfos.Student.ID)
				}
				break
			case 1:
				if getStudentInfos.Student.LastName != dataTest[i].lastName {
					t.Error("There is an error the last name does not match")
				}
				break
			case 2:
				if getStudentInfos.Student.LastName != dataTest[i].lastName {
					t.Error("There is an error the last name does not match")
				}
				break
			case 3:
				if getStudentInfos.Student.LastName != dataTest[i].phoneNumber {
					t.Error("There is an error the phone number does not match")
				}
				break
			}
		}
	}
}

////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////
// Test GetStudents ////////////////////////////////////////
func testGetStudents(t *testing.T) {
	dataTest := []struct {
		lastName string
		ordering string
		limit    uint
		offset   uint
	}{
		// Test 0 get a list of students
		{"", "", 10, 0},
		// Test 1 get a list of students who should match lastName
		// provaided
		{"polo", "", 10, 0},
		// Test 2 this should return one student with id 2
		{"", "", 10, 1},
	}

	for i := 0; i < len(dataTest); i++ {
		studentsListOut, err := MainService.
			GetStudents(dataTest[i].lastName, dataTest[i].ordering,
				dataTest[i].limit, dataTest[i].offset)
		if err != nil {
			t.Error(err)
		}
		switch i {
		case 0:
			if len(studentsListOut.Students) != int(dataTest[i].limit) {
				t.Errorf("There is an error count does not match should be %d but got %d",
					dataTest[i].limit, len(studentsListOut.Students))
			}
		case 1:
			if studentsListOut.Students[0].LastName != students[0]["last_name"] {
				t.Error("There is an error the first name does not match")
			}
			break
		case 2:
			if len(studentsListOut.Students) != int(dataTest[i].limit) {
				t.Errorf("There is an error count does not match should be %d but got %d",
					dataTest[i].limit-dataTest[i].offset, len(studentsListOut.Students))
			}
		}
	}
}

////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////
// Test create student /////////////////////////////////////
func testCreateStudent(t *testing.T) {
	// Connect to database test
	connectDatabse()

	for i := 0; i < len(students); i++ {
		_, err := MainService.CreateStudent(students[i])
		if err != nil {
			t.Error("There is an error student was not created")
		}
	}
}

////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////
// Test update student /////////////////////////////////////
func testUpdateStudent(t *testing.T) {
	var studentUpdated = make(map[string]interface{})
	for key, value := range students[0] {
		studentUpdated[key] = value
	}
	studentUpdated["id"] = 1
	studentUpdated["last_name"] = "super polo"
	studentUpdated["file_number"] = "200"

	student, err := MainService.UpdateStudent(studentUpdated)
	if err != nil {
		t.Fatal("There is an error the student was not updated")
	}
	if student.LastName != studentUpdated["last_name"] {
		t.Error("There is an error the student last name was not updated")
	}
	if student.FileNumber != studentUpdated["file_number"] {
		t.Error("There is an error the student file number was not updated")
	}
}

////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////
// Test delete student /////////////////////////////////////
func testDeleteStudent(t *testing.T) {
	for i := 0; i < len(students); i++ {
		if err := MainService.DeleteStudent(uint(i)); err != nil {
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
