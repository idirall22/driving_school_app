package service

import (
	"encoding/json"
	"log"
	"testing"
	"time"
)

// Number of students to add when create examList test
var studentsNumber = 3

// Test create an exam list
func testCreateExamList(t *testing.T) {
	exams := []*Exam{}

	for i := 0; i < studentsNumber; i++ {
		exam := &Exam{
			Exam:      1,
			DateExam:  time.Now(),
			Status:    false,
			StudentID: uint(i + 1),
		}
		exams = append(exams, exam)
	}
	examList := ExamList{
		DateExam:      time.Now(),
		Examiner:      "no name",
		Archived:      false,
		StudentsExams: exams,
	}

	// Get an examList map from a model
	m, err := getExamListMapFromModel(examList)
	if err != nil {
		log.Fatal(err)
	}

	id, errEL := MainService.CreateExamList(m)

	if errEL != nil {
		t.Error(errEL)
	}
	if id != 1 {
		t.Errorf("There is an error the id should be 1 found %d", id)
	}
}

// Test get a single exam list
func testGetExamList(t *testing.T) {
	id := uint(1)
	examList, err := MainService.GetExamList(id)
	if err != nil {
		log.Fatal(err)
	}

	if examList.ID != id {
		t.Errorf("There is an error, the id should be %d but got %d",
			id, examList.ID)
	}

	if len(examList.StudentsExams) != studentsNumber {
		t.Errorf("There is an error, the length should be %d but got %d",
			studentsNumber, len(examList.StudentsExams))
	}
}

// Test get a list of exam list
func testGetExamLists(t *testing.T) {
	limit := uint(1)
	examLists, err := MainService.GetExamLists(limit, 0)
	if err != nil {
		log.Fatal(err)
	}

	if len(examLists.ExamLists) != int(limit) && examLists.Count != limit {
		t.Errorf("There is an error, the length should be %d but got %d",
			limit, len(examLists.ExamLists))
	}
}

// Test update an examList
func testUpdateExamList(t *testing.T) {

	// Get exams from db and convert them to map
	exams := []*Exam{}
	if err := MainService.db.Find(&exams).Error; err != nil {
		log.Fatalf("Error find exams: %s", err)
	}

	examsMap := make(map[string]interface{})
	data, err := json.Marshal(examsMap)
	if err != nil {
		log.Fatalf("Error marshal exams: %s", err)
	}
	if err := json.Unmarshal(data, &examsMap); err != nil {
		log.Fatalf("Error unmarshal exams: %s", err)
	}
	exams[0].Status = true

	// Data test struct
	dataTest := []struct {
		exams       []*Exam
		examiner    string
		archived    bool
		date        time.Time
		studentsIDS []uint
	}{
		// Test 0 archive examList:
		// this return an error it say you can not archive an
		// exam list if it time is not equal to today or more
		{exams, "no name", true, time.Now().Add(time.Hour * 48), []uint{}},

		// Test 1 archive examList:
		// this test should successed
		{exams[:2], "no name", true, time.Now().Add(time.Hour), []uint{}},

		// Test 2 archive examList:
		// this test should successed the student of exams01[0]
		// and return student.NextExam = 2 and LastExamStatus true
		{exams[:2], "no name", true, time.Now().Add(time.Hour), []uint{}},

		// Test 3 archive examList:
		// this test should successed the student of exams01[0]
		// and return student.NextExam = 1 and LastExamStatus false
		{exams[:2], "no name", true, time.Now().Add(time.Hour), []uint{}},

		// Test 4 archive examList:
		// here we are going to add a new exam list fro student id 3
		{exams, "no name", false, time.Now().Add(time.Hour), []uint{}},

		// Test 5 archive examList:
		// this test should successed and return only two studentsExams
		// because we provaided an array with ids we want to delete
		{exams[1:], "no name", false, time.Now().Add(time.Hour), []uint{1}},

		// Test 6 examiner is empty
		{exams, "", false, time.Now(), []uint{}},
	}

	for i := 0; i < len(dataTest); i++ {
		// Exam list to test with
		m := make(map[string]interface{})
		m["id"] = 1
		m["date_exam"] = dataTest[i].date.Format(timeFormat)
		m["examiner"] = dataTest[i].examiner
		m["archived"] = dataTest[i].archived
		m["students_exams"] = dataTest[i].exams

		if i == 3 {
			exams[0].Status = false
		}

		examList, errU := MainService.UpdateExamList(m, dataTest[i].studentsIDS)

		if i == 0 {
			if errU.Error() != "Exam List not Valid" {
				t.Error("There is an error this need should be not valid")
			}
		} else if i == 1 {
			if examList.Archived != true {
				t.Error("There is an error this archived should be true")
			}
		} else if i == 2 {
			student := &Student{}
			MainService.db.Find(&student, "id=1")
			if student.NextExam != 2 {
				t.Errorf("There is an error this nextExam should be 2 found %d",
					student.NextExam)
			}
		} else if i == 3 {
			student := &Student{}
			MainService.db.Find(&student, "id=1")
			if student.NextExam != 1 {
				t.Errorf("There is an error this nextExam should be 1 found %d",
					student.NextExam)
			}
		} else if i == 4 {
			examsDB := []*Exam{}
			MainService.db.Find(&examsDB)

			if len(examList.StudentsExams) != len(examsDB) {
				t.Errorf("There is an error this length should be %d found %d",
					len(examList.StudentsExams), len(examsDB))
			}
		} else if i == 5 {
			examsDB := []*Exam{}
			MainService.db.Find(&examsDB)
			if len(examList.StudentsExams) != len(examsDB) {
				t.Errorf("There is an error this length should be %d found %d",
					len(examList.StudentsExams), len(examsDB))
			}
		}
	}
}

// Test delete exam list
func testDeleteExamList(t *testing.T) {
	if err := MainService.DeleteExamList(1); err != nil {
		t.Error(err)
	}
}

func testCleanedDatabase(t *testing.T) {
	student := []Student{}
	examList := []ExamList{}
	exams := []Exam{}

	MainService.db.Find(&student)
	MainService.db.Find(&examList)
	MainService.db.Find(&exams)

	if len(student) != 0 {
		t.Error("There is an error length should be zero")
	}

	if len(examList) != 0 {
		t.Error("There is an error length should be zero")
	}

	if len(exams) != 0 {
		t.Error("There is an error length should be zero")
	}
}
