package service

import (
	"context"
	"time"
)

// CreateExamList create an exam list
func CreateExamList(ctx context.Context, date time.Time,
	students []*Student) (*ExamList, error) {

	return nil, nil
}

// GetExamList get an exam list by date
func GetExamList(ctx context.Context, date time.Time) (*ExamList, error) {

	return nil, nil
}

// GetExamLists get a list of exam list
func GetExamLists(ctx context.Context, limit, offset string) ([]*ExamList, error) {

	return nil, nil
}

// UpdateExamList update an exam list
func UpdateExamList(ctx context.Context, examList *ExamList) error {

	return nil
}

// DeleteExamList update an exam list
func DeleteExamList(ctx context.Context, id uint) error {

	return nil
}
