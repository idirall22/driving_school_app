package service

import (
	"context"
	"testing"
	"time"
)

// Test create an exam list
func testCreateExamList(t *testing.T) {
	student, _ := GetStudent(context.Background(), 1)

	examList, errEL := CreateExamList(context.Background(),
		time.Now(), []*Student{student})
	if errEL != nil {
		t.Error(errEL)
	}
	if examList.Students[0].FirstName != student.FirstName {
		t.Error("There is an error, student does not match")
	}
}

// Test get a single exam list
func testGetExamList(t *testing.T) {
	examList, _ := GetExamList(context.Background(), 1)
	if len(examList.Students) != 1 {
		t.Error("There is an error, the length should be 1")
	}
}

// Test get a list of exam list
func testGetExamLists(t *testing.T) {
	examLists, _ := GetExamLists(context.Background(), 10, 0)
	if len(examLists) != 1 {
		t.Error("There is an error, the length should be 1")
	}
}

// Test Update exam list
func testUpdateExamList(t *testing.T) {

	examList, _ := GetExamList(context.Background(), 1)
	dateExam := time.Now().AddDate(0, 0, 1)
	examList.DateExam = dateExam

	// Update Date
	UpdateExamList(context.Background(), examList)
	examListUpdated, _ := GetExamList(context.Background(), 1)

	if examListUpdated.DateExam.Day() != dateExam.Day() {
		t.Error("There is an error the dateExam was not updated")
	}

	// Update students list
	examList.Students = []*Student{}
	UpdateExamList(context.Background(), examList)

	examListUpdated, _ = GetExamList(context.Background(), 1)

	if len(examListUpdated.Students) != 0 {
		t.Error("There is an error the students list was not updated")
	}

}

func testDeleteExamList(t *testing.T) {
	DeleteExamList(context.Background(), 1)
	e, _ := GetExamLists(context.Background(), 10, 0)

	if len(e) != 0 {
		t.Error("There is an error the lenght of examlists should be equal to 0")
	}
}
