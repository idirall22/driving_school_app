package service

import (
	"encoding/json"
	"fmt"
)

// GetStudent return a single student by provaiding id or lastName
func (s *Service) GetStudent(studentID uint,
	lastName string) (*GetStudentInfos, error) {
	getStudentInfos := &GetStudentInfos{
		Student: &Student{},
		Exams:   []*Exam{},
	}

	if studentID != 0 && lastName != "" {
		lastName = ""
	}

	tx := MainService.db.Begin()
	tx.Find(&getStudentInfos.Student, "id=? OR last_name=?", studentID, lastName)
	if err := tx.Find(&getStudentInfos.Exams, "id=?", studentID).
		Error; err != nil {
		return nil, err
	}
	tx.Commit()
	return getStudentInfos, nil
}

// StudentsListOut model
type StudentsListOut struct {
	Students []*Student `json:"students"`
	Count    uint       `json:"count"`
}

// GetStudents return a list of students,
func (s *Service) GetStudents(lastName, ordering string,
	limit, offset int) (*StudentsListOut, error) {

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
