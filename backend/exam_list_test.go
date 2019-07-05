package service

import (
	"encoding/json"
	"testing"
	"time"
)

// Test create an exam list
func testCreateExamList(t *testing.T) {
	student := MainService.GetStudent(1)

	examList, errEL := MainService.CreateExamList(
		time.Now(), []*Student{student})
	if errEL != nil {
		t.Error(errEL)
	}
	if examList.Students[0].FirstName != student.FirstName {
		t.Error("There is an error, student does not match")
	}
}

// Test CreateExamListMap
func testCreateExamListMap(t *testing.T) {
	students := []*Student{}
	student := MainService.GetStudent(1)
	students = append(students, student)

	var m []interface{}

	data, _ := json.Marshal(students)
	json.Unmarshal(data, &m)
	examinerName := "examiner Name"
	examList := MainService.CreateExamListMap("time.Now()", examinerName, m)

	if len(examList.Students) != 1 {
		t.Error("There is an error, with examlist created")
	}
	if examList.Examiner != examinerName {
		t.Error("There is an error, with examlist created")
	}

}

// Test get a single exam list
func testGetExamList(t *testing.T) {
	examList, _ := MainService.GetExamList(1)

	if examList.ID != 1 {
		t.Error("There is an error, the length should be 1")
	}
}

// Test get a list of exam list
func testGetExamLists(t *testing.T) {
	examLists := MainService.GetExamLists(10, 0)
	if len(examLists) != 1 {
		t.Error("There is an error, the length should be 1")
	}
}

// Test Update exam list
func testUpdateExamList(t *testing.T) {

	examList, _ := MainService.GetExamList(1)
	dateExam := time.Now().AddDate(0, 0, 1)
	examList.DateExam = dateExam

	// Update Date
	MainService.UpdateExamList(examList)
	examListUpdated, _ := MainService.GetExamList(1)

	if examListUpdated.DateExam.Day() != dateExam.Day() {
		t.Error("There is an error the dateExam was not updated")
	}

	// Update students list
	examList.Students = []*Student{}
	MainService.UpdateExamList(examList)

	examListUpdated, _ = MainService.GetExamList(1)

	if len(examListUpdated.Students) != 0 {
		t.Error("There is an error the students list was not updated")
	}

}

func testDeleteExamList(t *testing.T) {
	MainService.DeleteExamList(1)
	e := MainService.GetExamLists(10, 0)

	if len(e) != 0 {
		t.Error("There is an error the lenght of examlists should be equal to 0")
	}
}
