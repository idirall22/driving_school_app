package service

import (
	"encoding/json"
	"time"
)

// GetStudent return a single student
func (s *Service) GetStudent(id int64) *Student {
	student := &Student{}
	MainService.db.Find(&student, "id=?", id).Related(&student.Exams)
	return student
}

// GetStudentByName return a single student by last_name
func (s *Service) GetStudentByName(lastName string) *Student {
	student := &Student{}
	MainService.db.Find(&student, "last_name=?", lastName).Related(&student.Exams)
	return student
}

// GetStudents return a list of students
func (s *Service) GetStudents(limit, offset int) []*Student {

	var students = []*Student{}
	MainService.db.Limit(limit).Offset(offset).
		Order("registred_date desc").Find(&students)
	return students
}

//CreateStudentMap create a student from a map
func (s *Service) CreateStudentMap(studentInfo map[string]interface{}) (uint, error) {

	student := &Student{}
	studentInfo["birthday"] = time.Now()
	studentInfo["registred_date"] = time.Now()
	data, err := json.Marshal(studentInfo)
	if err != nil {
		return 0, err
	}
	if err := json.Unmarshal(data, &student); err != nil {
		return 0, err
	}
	MainService.db.Create(&student)
	return student.ID, nil
}

// CreateStudent create a student
func (s *Service) CreateStudent(student *Student) {
	student.BirthDay = time.Now()
	student.RegistredDate = time.Now()
	MainService.db.Create(&student)
}

// UpdateStudent update a student
func (s *Service) UpdateStudent(student *Student) {
	MainService.db.Save(&student)
}

// DeleteStudent update a student
func (s *Service) DeleteStudent(id uint) {
	MainService.db.Unscoped().Where("id=?", id).Delete(Student{})
}
