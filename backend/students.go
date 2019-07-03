package service

// GetStudent return a single student
func (s *service) GetStudent(id int64) (*Student, error) {
	student := &Student{}
	MainService.db.Find(&student, "id=?", id).Related(&student.Exams)
	return student, nil
}

// GetStudents return a list of students
func (s *service) GetStudents(limit, offset int) ([]*Student, error) {

	var students = []*Student{}
	MainService.db.Limit(limit).Offset(offset).Find(&students)
	return students, nil
}

// CreateStudent create a student
func (s *service) CreateStudent(student *Student) {
	MainService.db.Create(&student)
}

// UpdateStudent update a student
func (s *service) UpdateStudent(student *Student) {
	MainService.db.Save(&student)
}

// DeleteStudent update a student
func (s *service) DeleteStudent(id uint) {
	MainService.db.Unscoped().Where("id=?", id).Delete(Student{})
}
