package service

import "context"

// GetStudent return a single student
func GetStudent(ctx context.Context, id int64) (*Student, error) {
	return nil, nil
}

// GetStudents return a list of students
func GetStudents(ctx context.Context, limit, offset int) ([]*Student, error) {
	return nil, nil
}

// CreateStudent create a student
func CreateStudent(ctx context.Context, student *Student) error {
	return nil
}

// UpdateStudent update a student
func UpdateStudent(ctx context.Context, student *Student) error {
	return nil
}

// DeleteStudent update a student
func DeleteStudent(ctx context.Context, id int64) error {
	return nil
}
