package service

import (
	"encoding/json"
	"testing"
	"time"
)

// Test create an exam list
func testCreateExamList(t *testing.T) {
	students, err := MainService.GetStudents("", "", 3, 0)

	if err != nil {
		t.Error(err)
	}
	data, err := json.Marshal(students)
	if err != nil {
		t.Error(err)
	}
	var studentsMap = make([]interface{}, 30)
	if err := json.Unmarshal(data, &studentsMap); err != nil {
		t.Error(err)
	}

	examDate := time.Now().Format(timeFormat)
	examiner := "examiner name"

	examList, errEL := MainService.CreateExamList(
		examDate, examiner, studentsMap)

	if errEL != nil {
		t.Error(errEL)
	}

	if examList != nil {
		if examList.ID != 1 {
			t.Error("There is an error when exam list was created")
		}
	}
}

// Test get a single exam list
func testGetExamList(t *testing.T) {
	examList, err := MainService.GetExamList(1)
	if err != nil {
		t.Error("There is an error, to get exam list")
	}

	if examList != nil {
		if examList.ID != 1 {
			t.Error("There is an error, the id should be 1")
		}

		if len(examList.StudentsExams) != 3 {
			t.Error("There is an error, the length should be 1")
		}
	}
}

// Test get a list of exam list
func testGetExamLists(t *testing.T) {
	examLists, err := MainService.GetExamLists(10, 0)

	if examLists != nil {
		if err != nil {
			t.Error("There is an error, to get exam lists")
		}
		if len(examLists.ExamLists) != 1 && examLists.Count != 1 {
			t.Error("There is an error, the length should be 1")
		}
	}
}

// Test Update exam list
func testUpdateExamList(t *testing.T) {
	examList, err := MainService.GetExamList(1)
	if err != nil {
		t.Error(err)
	}

	examDate := time.Now().Format(timeFormat)
	newExaminer := "new Examiner name"
	m := make([]interface{}, 30)

	data, err := json.Marshal(examList.StudentsExams)
	if err != nil {
		t.Error("Error to marshal students exams")
	}

	if err := json.Unmarshal(data, &m); err != nil {
		t.Error("Error to unmarshal students exams bytes")
	}

	studentsToAdd, err := MainService.GetStudents("", "", 2, 3)
	if err != nil {
		t.Error(err)
	}

	data, err = json.Marshal(studentsToAdd)
	if err != nil {
		t.Error("Error to marshal students")
	}
	studentsToAddMap := make([]interface{}, 30)
	if err := json.Unmarshal(data, &studentsToAddMap); err != nil {
		t.Error("Error to unmarshal students exams bytes")
	}

	examListUpdated, errUpdate := MainService.UpdateExamList(1, examDate, newExaminer,
		m, studentsToAddMap)

	if errUpdate != nil {
		t.Error(errUpdate)
	}
	if examListUpdated != nil {
		if len(examListUpdated.StudentsExams) != testStudentsCount {
			t.Errorf("There is an error the length should be %d", testStudentsCount)
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
