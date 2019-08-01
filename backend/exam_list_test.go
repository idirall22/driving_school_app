package service

import (
	"log"
	"testing"
	"time"
)

// Test create an exam list
func testCreateExamList(t *testing.T) {
	exams := []*Exam{}

	for i := 0; i < len(students); i++ {
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

	//create an exam list
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

	if len(examList.StudentsExams) != len(students) {
		t.Errorf("There is an error, the length should be %d but got %d",
			len(students), len(examList.StudentsExams))
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
	if err := MainService.db.Find(&exams).Limit(5).Error; err != nil {
		log.Fatalf("Error find exams: %s", err)
	}

	examList, _ := MainService.GetExamList(1)
	studentDeletedFromExamList := make([]interface{}, 0)
	examList.StudentsExams = []*Exam{}
	for i := 0; i < len(exams); i++ {
		if i%2 == 0 {
			examList.StudentsExams = append(examList.StudentsExams, exams[i])
		} else {
			studentDeletedFromExamList = append(studentDeletedFromExamList,
				exams[i].ID)
		}
	}
	m, e := getExamListMapFromModel(*examList)
	if e != nil {
		t.Fatal(e)
	}

	ex, err := MainService.UpdateExamList(m, studentDeletedFromExamList)
	if err != nil {
		t.Fatal(err)
	}
	if len(ex.StudentsExams) != len(exams)-len(studentDeletedFromExamList) {
		t.Errorf("There is an error the length should be %d, but got %d",
			len(exams)-len(studentDeletedFromExamList), len(ex.StudentsExams))
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
		t.Errorf("There is an error length should be zero but got %d",
			len(exams))
	}
}
