package service

// GetStudent return a single student
func (s *Service) GetStudent(id int64) *Student {
	student := &Student{}
	MainService.db.Find(&student, "id=?", id).Related(&student.Exams)
	return student
}

// GetStudents return a list of students
func (s *Service) GetStudents(limit, offset int) []*Student {

	var students = []*Student{}
	MainService.db.Limit(limit).Offset(offset).Find(&students)
	return students
}

// CreateStudent create a student
func (s *Service) CreateStudent(student *Student) {
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
