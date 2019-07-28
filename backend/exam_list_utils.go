package service

import (
	"encoding/json"
	"errors"
	"time"
)

// Check if exam list is valid
func checkIfExamListIsValid(examList *ExamList) bool {
	// Check if examiner is not empty
	if examList.Examiner == "" {
		examList.Examiner = "No Name"
	}

	if examList.Archived {
		if examList.DateExam.Day() <= time.Now().Day() {
			return true
		}
		return false
	}
	return true
}

// Set next exam level
func setNextExam(dateExam time.Time, lastExamDate *time.Time,
	LastExamStatus, status bool, nextExam uint8) uint8 {

	// When the examList is archived first time
	if lastExamDate == nil {
		if status {
			return nextExam + 1
		}
	} else {
		if dateExam.Unix() == lastExamDate.Unix() {
			if LastExamStatus {
				nextExam--
			}
			if status {
				nextExam++
			}
			return nextExam
		}
	}
	return nextExam
}

// Return an examList model from a map
func getExamListModelFromMap(examListMap map[string]interface{}) (*ExamList, error) {
	data, err := json.Marshal(examListMap)
	if err != nil {
		return nil, err
	}
	examList := &ExamList{}
	if err := json.Unmarshal(data, &examList); err != nil {
		return nil, err
	}
	if !checkIfExamListIsValid(examList) {
		return nil, errors.New("Exam List not Valid")
	}
	return examList, nil
}

func getExamListMapFromModel(examList ExamList) (map[string]interface{}, error) {
	data, err := json.Marshal(&examList)
	if err != nil {
		return nil, err
	}
	m := make(map[string]interface{})
	if err := json.Unmarshal(data, &m); err != nil {
		return nil, err
	}
	return m, nil
}
