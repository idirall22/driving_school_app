package service

import (
	"context"
	"time"
)

// CreateExamList create an exam list
func CreateExamList(ctx context.Context, date time.Time,
	students []*Student) (*ExamList, error) {

	examList := &ExamList{DateExam: date}
	mainService.db.Create(&examList).Association("Students").Append(students)

	return examList, nil
}

// GetExamList get an exam list by date
func GetExamList(ctx context.Context, id uint) (*ExamList, error) {

	examList := &ExamList{}
	mainService.db.Select("*").Find(&examList, "id=?", id).
		Related(&examList.Students, "Students")
	return examList, nil
}

// GetExamLists get a list of exam list
func GetExamLists(ctx context.Context, limit, offset uint) ([]*ExamList, error) {
	examLists := []*ExamList{}
	mainService.db.Limit(limit).Offset(offset).Find(&examLists)
	return examLists, nil
}

// UpdateExamList update an exam list
func UpdateExamList(ctx context.Context, examList *ExamList) error {

	mainService.db.Save(&examList).Association("Students").
		Replace(examList.Students)

	return nil
}

// DeleteExamList update an exam list
func DeleteExamList(ctx context.Context, id uint) error {
	mainService.db.Unscoped().Where("id=?", id).
		Delete(ExamList{}).Association("Students").Clear()
	return nil
}
