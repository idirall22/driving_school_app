package service

import (
	"encoding/json"
	"errors"
)

// Return a student model from a map
func getStudentModelFromMap(studentMap map[string]interface{}) (*Student, error) {
	student := &Student{}
	data, err := json.Marshal(studentMap)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &student); err != nil {
		return nil, err
	}
	if !checkIfStudentIsValid(student) {
		return nil, errors.New("Error student not valid")
	}
	return student, nil
}

// Check if student model is valid
func checkIfStudentIsValid(student *Student) bool {
	if student.FirstName == "" ||
		student.LastName == "" ||
		student.Gender == "" ||
		student.Country == "" ||
		student.Department == "" ||
		student.City == "" ||
		student.AddressStreet == "" {
		return false
	}
	return true
}

// return a map of a student from a struct provaided
func getStudentMapFromModel(student Student) (map[string]interface{}, error) {
	data, err := json.Marshal(&student)
	if err != nil {
		return nil, err
	}
	m := make(map[string]interface{})
	if err := json.Unmarshal(data, &m); err != nil {
		return nil, err
	}
	return m, nil
}
