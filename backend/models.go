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

// GetStudentInfos model returned when request single student infos
type GetStudentInfos struct {
	Student *Student `json:"student"`
	Exams   []*Exam  `json:"exams"`
}

// MultiLanguageField model
type MultiLanguageField struct {
	FR string `json:"fr"`
	AR string `json:"ar"`
}

//Student model
type Student struct {
	ID            uint       `gorm:"primary_key" json:"id,omitempty"`
	CreatedAt     time.Time  `sql:"DEFAULT:current_timestamp" json:"created_at,omitempty"`
	UpdatedAt     time.Time  `json:"updated_at,omitempty"`
	DeletedAt     *time.Time `sql:"index"`
	FileNumber    string     `gorm:"unique" json:"file_number,omitempty"`
	FirstName     string     `gorm:"type:varchar(512)" json:"first_name,omitempty"`
	LastName      string     `gorm:"type:varchar(512);index:last_name" json:"last_name,omitempty"`
	MaidenName    string     `gorm:"type:varchar(512)" json:"maiden_name,omitempty"`
	PhoneNumber   string     `json:"phone_number,omitempty"`
	Job           string     `json:"job,omitempty"`
	BirthDay      time.Time  `json:"birthday,omitempty"`
	Gender        string     `json:"gender,omitempty"`
	Country       string     `json:"country,omitempty"`
	City          string     `json:"city,omitempty"`
	Department    string     `json:"department,omitempty"`
	AddressStreet string     `gorm:"type:varchar(512)" json:"address_street,omitempty"`
	RegistredDate time.Time  `json:"registred_date,omitempty"`
	Image         string     `gorm:"default:'imageURL'" json:"image,omitempty"`
	NextExam      string     `gorm:"default:'Highway code'" json:"next_exam,omitempty"`
	LastExamDate  *time.Time `gorm:"index:last_exam_date" json:"last_exam_date,omitempty"`
	Archived      bool       `json:"archived"`
}

// TableName :Database table name
func (s Student) TableName() string {
	return "students"
}

// Exam an exam passed by a student
type Exam struct {
	ID         uint       `gorm:"unique_index" json:"id,omitempty"`
	CreatedAt  time.Time  `sql:"DEFAULT:current_timestamp" json:"created_at,omitempty"`
	UpdatedAt  time.Time  `json:"updated_at,omitempty"`
	DeletedAt  *time.Time `sql:"index"`
	ExamName   string     `json:"exam,omitempty"`
	DateExam   time.Time  `json:"date_exam,omitempty"`
	Status     bool       `gorm:"default:false" json:"status,omitempty"`
	StudentID  uint       `json:"student_id,omitempty"`
	Student    Student    `json:"student,omitempty"`
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
	Archived      bool       `json:"archived"`
	StudentsExams []*Exam    `json:"students_exams,omitempty"`
}

// TableName :Database table name
func (s ExamList) TableName() string {
	return "exam_lists"
}
