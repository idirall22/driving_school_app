package service

import (
	"context"
)

// GetStudent return a single student
func GetStudent(ctx context.Context, id int64) (*Student, error) {
	student := &Student{}
	mainService.db.Find(&student, "id=?", id).Related(&student.Exams)
	return student, nil
}

// GetStudents return a list of students
func GetStudents(ctx context.Context, limit, offset int) ([]*Student, error) {

	var students = []*Student{}
	mainService.db.Limit(limit).Offset(offset).Find(&students)
	return students, nil
}

// CreateStudent create a student
func CreateStudent(ctx context.Context, student *Student) {
	mainService.db.Create(&student)
}

// UpdateStudent update a student
func UpdateStudent(ctx context.Context, student *Student) {
	mainService.db.Save(&student)
}

// DeleteStudent update a student
func DeleteStudent(ctx context.Context, id uint) {
	mainService.db.Unscoped().Where("id=?", id).Delete(Student{})
}
