package service

import (
	"time"

	"github.com/jinzhu/gorm"
)

var (
	//ExamOne Highway code
	ExamOne = "Highway code"

	//ExamTwo Niche
	ExamTwo = "Niche"

	//ExamThree Circuit
	ExamThree = "Circuit"
)

//Student model
type Student struct {
	gorm.Model
	FileNumber    string    `gorm:"unique_index" json:"file_number,omitempty"`
	FirstName     string    `gorm:"index:first_name" json:"first_name,omitempty"`
	LastName      string    `gorm:"index:last_name" json:"last_name,omitempty"`
	MaidenName    string    `json:"maiden_name,omitempty"`
	PhoneNumber   string    `json:"phone_number,omitempty"`
	Job           string    `json:"job,omitempty"`
	BirthDay      time.Time `json:"birthday,omitempty"`
	Gender        string    `json:"gender,omitempty"`
	City          string    `json:"city,omitempty"`
	AddressStreet string    `json:"address_street,omitempty"`
	RegistredDate time.Time `json:"registred_date,omitempty"`
	Image         string    `json:"image,omitempty"`
	Exams         []*Exam   `json:"exams,omitempty"`
}

// TableName :Database table name
func (s Student) TableName() string {
	return "students"
}

// Exam an exam passed by a student
type Exam struct {
	gorm.Model
	ExamName  string    `json:"exam,omitempty"`
	Examiner  string    `json:"examiner,omitempty"`
	Comment   string    `json:"comment,omitempty"`
	DateExam  time.Time `json:"date_exam,omitempty"`
	Status    bool      `json:"status,omitempty"`
	StudentID int64     `json:"students_id,omitempty"`
}

// TableName :Database table name
func (e Exam) TableName() string {
	return "exams"
}

//ExamList a list of students who are added to pass an exam
type ExamList struct {
	gorm.Model
	DateExam time.Time  `json:"date_exam,omitempty"`
	Students []*Student `gorm:"many2many:examlists_students;" json:"students,omitempty"`
}

// TableName :Database table name
func (s ExamList) TableName() string {
	return "exam_lists"
}
