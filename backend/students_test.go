package service

import (
	"context"
	"fmt"
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
func testGetStudent(t *testing.T) {
	connectDatabse()
	mainService.db.
		Create(&Student{
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
		})

	student, err := GetStudent(context.Background(), 1)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(student.Exams[1].Comment)
}
