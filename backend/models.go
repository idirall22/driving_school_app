package service

import (
	"time"
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
	ID             uint       `gorm:"primary_key" json:"id,omitempty"`
	CreatedAt      time.Time  `json:"created_at,omitempty"`
	UpdatedAt      time.Time  `json:"updated_at,omitempty"`
	DeletedAt      *time.Time `sql:"index"`
	FileNumber     string     `gorm:"unique" json:"file_number,omitempty"`
	FirstName      string     `gorm:"first_name" json:"first_name,omitempty"`
	LastName       string     `gorm:"index:last_name" json:"last_name,omitempty"`
	MaidenName     string     `json:"maiden_name,omitempty"`
	PhoneNumber    string     `json:"phone_number,omitempty"`
	Job            string     `json:"job,omitempty"`
	BirthDay       time.Time  `json:"birthday,omitempty"`
	Gender         string     `json:"gender,omitempty"`
	Country        string     `json:"country,omitempty"`
	City           string     `json:"city,omitempty"`
	Department     string     `json:"department,omitempty"`
	AddressStreet  string     `json:"address_street,omitempty"`
	RegistredDate  time.Time  `json:"registred_date,omitempty"`
	Image          string     `gorm:"default:'imageURL'" json:"image,omitempty"`
	NextExam       string     `gorm:"default:'Highway code'" json:"next_exam,omitempty"`
	LastExamDate   *time.Time `gorm:"index:last_exam_date" json:"last_exam_date,omitempty"`
	LastExamStatus bool       `gorm:"default:false" json:"last_exam_status,omitempty"`
}

// TableName :Database table name
func (s Student) TableName() string {
	return "students"
}

// Exam an exam passed by a student
type Exam struct {
	ID         uint       `gorm:"unique_index" json:"id,omitempty"`
	CreatedAt  time.Time  `json:"created_at,omitempty"`
	UpdatedAt  time.Time  `json:"updated_at,omitempty"`
	DeletedAt  *time.Time `sql:"index"`
	ExamName   string     `json:"exam,omitempty"`
	DateExam   time.Time  `json:"date_exam,omitempty"`
	Status     bool       `json:"status,omitempty"`
	StudentID  uint       `json:"students_id,omitempty"`
	ExamListID uint       `json:"exam_list_id,omitempty"`
}

// TableName :Database table name
func (e Exam) TableName() string {
	return "exams"
}

//ExamList a list of students who are added to pass an exam
type ExamList struct {
	ID            uint       `gorm:"unique_index" json:"id,omitempty"`
	CreatedAt     time.Time  `json:"created_at,omitempty"`
	UpdatedAt     time.Time  `json:"updated_at,omitempty"`
	DeletedAt     *time.Time `sql:"index"`
	DateExam      time.Time  `json:"date_exam,omitempty"`
	Examiner      string     `json:"examiner,omitempty"`
	StudentsExams []*Exam    `json:"students_exams,omitempty"`
}

// TableName :Database table name
func (s ExamList) TableName() string {
	return "exam_lists"
}
