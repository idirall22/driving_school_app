package service

import (
	"encoding/json"
	"log"
	"time"
)

// CreateExamList create an exam list
func (s *Service) CreateExamList(date time.Time, students []*Student) (*ExamList, error) {

	examList := &ExamList{DateExam: date}
	MainService.db.Create(&examList).Association("Students").Append(students)

	return examList, nil
}

// CreateExamListMap create an exam list
func (s *Service) CreateExamListMap(date, examiner string,

	studentsList []interface{}) *ExamList {

	students := []*Student{}
	data, err := json.Marshal(studentsList)
	if err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(data, &students); err != nil {
		log.Fatal(err)
	}
	if examiner == "" {
		examiner = "No Name"
	}

	examList := &ExamList{DateExam: time.Now(), Examiner: examiner}
	MainService.db.Create(&examList).Association("Students").Append(students)

	return examList
}

// GetExamList get an exam list by date
func (s *Service) GetExamList(id uint) (*ExamList, error) {

	examList := &ExamList{}
	MainService.db.Select("*").Find(&examList, "id=?", id).
		Related(&examList.Students, "Students")
	return examList, nil
}

// GetExamLists get a list of exam list
func (s *Service) GetExamLists(limit, offset uint) []*ExamList {
	examLists := []*ExamList{}
	MainService.db.Limit(limit).Offset(offset).Order("date_exam desc").Find(&examLists)
	return examLists
}

// UpdateExamList update an exam list
func (s *Service) UpdateExamList(examList *ExamList) error {

	MainService.db.Save(&examList).Association("Students").
		Replace(examList.Students)

	return nil
}

// DeleteExamList update an exam list
func (s *Service) DeleteExamList(id uint) error {
	MainService.db.Unscoped().Where("id=?", id).
		Delete(ExamList{}).Association("Students").Clear()
	return nil
}
