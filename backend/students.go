package service

import (
	"context"
)

// GetStudent return a single student
func GetStudent(ctx context.Context, id int64) (*Student, error) {
	student := &Student{}
	rows, err := mainService.db.Table(student.TableName()).Select("*").
		Joins("LEFT JOIN exams on exams.student_id = students.id").
		Where("students.id=?", id).Rows()

	if err != nil {
		return nil, err
	}
	exams := []*Exam{}
	for rows.Next() {
		exam := &Exam{}
		if err := rows.Scan(
			&student.ID,
			&student.CreatedAt,
			&student.UpdatedAt,
			&student.DeletedAt,
			&student.FileNumber,
			&student.FirstName,
			&student.LastName,
			&student.MaidenName,
			&student.PhoneNumber,
			&student.Job,
			&student.BirthDay,
			&student.Gender,
			&student.City,
			&student.AddressStreet,
			&student.RegistredDate,
			&student.Image,

			&exam.ID,
			&exam.CreatedAt,
			&exam.UpdatedAt,
			&exam.DeletedAt,
			&exam.ExamName,
			&exam.Examiner,
			&exam.Comment,
			&exam.DateExam,
			&exam.Status,
			&exam.StudentID,
		); err != nil {
			return nil, err
		}
		exams = append(exams, exam)
	}
	student.Exams = exams
	return student, nil
}

// GetStudents return a list of students
func GetStudents(ctx context.Context, limit, offset int) ([]*Student, error) {
	rows, err := mainService.db.Select("first_name, last_name, registred_date").
		Find(&Student{}).Limit(limit).Offset(offset).Rows()

	if err != nil {
		return nil, err
	}
	var students = []*Student{}
	for rows.Next() {
		student := &Student{}
		if err := mainService.db.ScanRows(rows, &student); err != nil {
			return nil, err
		}
		students = append(students, student)
	}
	return students, nil
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
