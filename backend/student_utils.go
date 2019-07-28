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
