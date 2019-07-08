package service

import (
	"encoding/json"
	"fmt"
)

// GetStudent return a single student by provaiding id or lastName
func (s *Service) GetStudent(studentID uint, lastName string) (*Student, error) {
	student := &Student{}

	if studentID != 0 && lastName != "" {
		lastName = ""
	}

	if err := MainService.db.Find(&student,
		"id=? OR last_name=?", studentID, lastName).Error; err != nil {
		return nil, err
	}
	return student, nil
}

// GetStudents return a list of students,
func (s *Service) GetStudents(lastName, ordering string, limit, offset int) ([]*Student, error) {

	defaultOrdering := "desc"
	if ordering != "" {
		if ordering != "asc" || ordering != defaultOrdering {
			ordering = defaultOrdering
		}
	} else {
		ordering = defaultOrdering
	}

	orderingString := fmt.Sprintf("registred_date %s", ordering)
	var students = []*Student{}
	if err := MainService.db.Limit(limit).Offset(offset).
		Order(orderingString).
		Find(&students, "last_name LIKE ?", lastName+"%").Error; err != nil {
		return nil, err
	}
	return students, nil
}

//CreateStudent create a student from a map
func (s *Service) CreateStudent(studentInfo map[string]interface{}) (uint, error) {

	student := &Student{}
	data, err := json.Marshal(studentInfo)
	if err != nil {
		return 0, err
	}
	if err := json.Unmarshal(data, &student); err != nil {
		return 0, err
	}
	if err := MainService.db.Create(&student).Error; err != nil {
		return 0, err
	}
	return student.ID, nil
}

// UpdateStudent update a student
func (s *Service) UpdateStudent(studentUpdate map[string]interface{}) (uint, error) {
	data, err := json.Marshal(studentUpdate)
	if err != nil {
		return 0, err
	}
	student := &Student{}
	if err := json.Unmarshal(data, &student); err != nil {
		return 0, err
	}
	dbRes := MainService.db.Save(&student)

	if dbRes.Error != nil {
		return 0, dbRes.Error
	}
	return uint(dbRes.RowsAffected), nil
}

// DeleteStudent update a student
func (s *Service) DeleteStudent(id uint) error {
	student := Student{ID: id}

	err := MainService.db.Delete(&student).Error
	return err
}
