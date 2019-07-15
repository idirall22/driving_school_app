package service

import (
	"encoding/json"
	"log"
	"time"
)

var timeFormat = time.RFC3339

// CreateExamList create an exam list
func (s *Service) CreateExamList(date, examiner string,
	studentsList []interface{}) (*ExamList, error) {

	// Marshal an unmarshal the list interface into a list of students
	students := []*Student{}
	data, err := json.Marshal(studentsList)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(data, &students); err != nil {
		return nil, err
	}

	// Chek if there is a name for examiner
	if examiner == "" {
		examiner = "No Name"
	}

	// Parse exam date
	dateTime, err := time.Parse(timeFormat, date)
	if err != nil {
		return nil, err
	}

	// Create an exam instance
	examList := &ExamList{
		DateExam: dateTime,
		Examiner: examiner,
	}

	if err := MainService.db.Create(&examList).Error; err != nil {
		return nil, err
	}

	query := `INSERT INTO exams (
			exam_name,
			date_exam,
			status,
			student_id,
			exam_list_id
			)Values(
				?,?,?,?,?
				);`

	tx := MainService.db.Begin()
	for i := 0; i < len(students); i++ {
		if err := tx.Exec(query,
			students[i].NextExam,
			examList.DateExam,
			false,
			students[i].ID,
			examList.ID,
		).Error; err != nil {
			return nil, err
		}
	}
	tx.Commit()
	return examList, nil
}

// GetExamList get an exam list by id
func (s *Service) GetExamList(id uint) (*ExamList, error) {

	examList := &ExamList{}
	tx := MainService.db.Begin()

	if err := tx.Find(&examList, "id=?", id).
		Related(&examList.StudentsExams, "StudentsExams").
		Error; err != nil {
		return nil, err
	}
	for i := 0; i < len(examList.StudentsExams); i++ {
		if err := tx.Find(&examList.StudentsExams[i].Student, "id=?",
			examList.StudentsExams[i].StudentID).Error; err != nil {
			return nil, err
		}
	}
	tx.Commit()
	// if err := MainService.db.Find(&examList, "id=?", id).
	// 	Related(&examList.StudentsExams, "StudentsExams").
	// 	Error; err != nil {
	// 	return nil, err
	// }

	return examList, nil
}

// ExamListsOut model
type ExamListsOut struct {
	ExamLists []*ExamList `json:"examLists"`
	Count     uint        `json:"count"`
}

// GetExamLists get a list of exam list
func (s *Service) GetExamLists(limit, offset uint) (*ExamListsOut, error) {
	examListsOut := &ExamListsOut{}

	if err := MainService.db.
		Model(&ExamList{}).
		Order("date_exam desc").
		Count(&examListsOut.Count).
		Limit(limit).Offset(offset).
		Find(&examListsOut.ExamLists).
		Error; err != nil {
		return nil, err
	}

	return examListsOut, nil
}

// UpdateExamList update an exam list
func (s *Service) UpdateExamList(examListID uint, date, examiner string,
	exams []interface{}, studentsList []interface{}) (*ExamList, error) {

	// Parse exam date
	dateParsed, errParse := time.Parse(timeFormat, date)
	if errParse != nil {
		return nil, errParse
	}

	newStudentsExams := []*Exam{}
	if len(studentsList) > 0 {
		students := []*Student{}
		data, err := json.Marshal(studentsList)
		if err != nil {
			return nil, err
		}
		if err := json.Unmarshal(data, &students); err != nil {
			return nil, err
		}
		for i := 0; i < len(students); i++ {
			tx := MainService.db.Begin()
			exam := &Exam{
				ExamName:   students[i].NextExam,
				DateExam:   dateParsed,
				StudentID:  students[i].ID,
				ExamListID: examListID,
			}
			if err := tx.Create(&exam).Error; err != nil {
				return nil, err
			}
			newStudentsExams = append(newStudentsExams, exam)
			tx.Commit()
		}
	}

	examList := &ExamList{
		ID:       examListID,
		Examiner: examiner,
		DateExam: dateParsed,
	}

	if len(exams) > 0 {
		data, err := json.Marshal(exams)
		if err != nil {
			return nil, err
		}
		if err := json.Unmarshal(data, &examList.StudentsExams); err != nil {
			return nil, err
		}
	}
	examList.StudentsExams = append(examList.StudentsExams, newStudentsExams...)

	if err := MainService.db.Save(&examList).Association("StudentsExams").
		Replace(examList.StudentsExams).Error; err != nil {
		return nil, err
	}

	return examList, nil
}

// DeleteExamList update an exam list
func (s *Service) DeleteExamList(id uint) error {
	examList := ExamList{ID: id}

	tx := MainService.db.Begin()
	tx.Find(&examList, "id=?", id).Related(&examList.StudentsExams).
		Unscoped().Delete(&examList).
		Unscoped().Delete(&examList.StudentsExams)
	err := tx.Commit().Error
	return err
}
