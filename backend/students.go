package service

import (
	"fmt"
)

// GetStudent return a single student by id, lastName or phoneNumber
func (s *Service) GetStudent(studentID uint,
	lastName, phoneNumber string) (*GetStudentInfos, error) {

	// If multiple parametters provaided the id will be the main
	// parametter for searching
	// searching priority: 1: id, 2: lastName, 3: phoneNumber
	if studentID == 0 {
		if lastName != "" && phoneNumber != "" {
			phoneNumber = ""
		}

	} else {
		phoneNumber = ""
		lastName = ""
	}

	getStudentInfos := &GetStudentInfos{
		Student: &Student{},
		Exams:   []*Exam{},
	}

	// Begin tx
	tx := MainService.db.Begin()
	if err := tx.Find(&getStudentInfos.Student,
		"id=? OR last_name Like ? OR phone_number=?",
		studentID, "%"+lastName+"%", phoneNumber).
		Related(&getStudentInfos.Exams).
		Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return getStudentInfos, nil
}

// GetStudents return a list of students,
func (s *Service) GetStudents(lastName, ordering string,
	limit, offset uint) (*StudentsListOut, error) {

	defaultOrdering := "desc"
	if ordering != "" {
		if ordering != "asc" || ordering != defaultOrdering {
			ordering = defaultOrdering
		}
	} else {
		ordering = defaultOrdering
	}

	orderingString := fmt.Sprintf("registred_date %s", ordering)
	var studentsListOut = &StudentsListOut{}
	if err := MainService.db.
		Model(&Student{}).
		Order(orderingString).
		Count(&studentsListOut.Count).
		Limit(limit).Offset(offset).
		Find(&studentsListOut.Students, `last_name LIKE ?`, "%"+lastName+"%").
		Error; err != nil {
		return nil, err
	}
	return studentsListOut, nil
}

//CreateStudent create a student from a map
func (s *Service) CreateStudent(studentMap map[string]interface{}) (uint, error) {

	student, err := getStudentModelFromMap(studentMap)
	if err != nil {
		return 0, nil
	}
	if err := MainService.db.Create(&student).Error; err != nil {
		return 0, err
	}
	return student.ID, nil
}

// UpdateStudent update a student
func (s *Service) UpdateStudent(studentMap map[string]interface{}) (*Student, error) {
	student, err := getStudentModelFromMap(studentMap)
	if err != nil {
		return nil, err
	}

	if err := MainService.db.
		Model(&Student{}).
		Updates(&student).Error; err != nil {
		return nil, err
	}
	return student, nil
}

// DeleteStudent update a student
func (s *Service) DeleteStudent(id uint) error {
	student := Student{ID: id}
	tx := MainService.db.Begin()
	tx.Delete(&student)
	exams := []*Exam{}
	err := tx.
		Table("exams").
		Where("student_id=?", id).
		Unscoped().
		Delete(&exams).
		Error
	tx.Commit()
	return err
}
