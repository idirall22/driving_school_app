package service

import (
	"strconv"
	"time"
)

func getYear(year int) string {
	t := time.Now()

	if year == 0 {
		year = 0
	} else {
		if year < 2010 || year > time.Now().Year() {
			year = t.Year()
		}
	}
	return strconv.Itoa(year)
}

// GetWinLicencePerYear get students win their licence per date
func (s *Service) GetWinLicencePerYear(year int) ([]string, error) {

	tx := MainService.db.Begin()
	dates := []string{}

	rows, err := tx.Table("students").Select("win_date").
		Where("win_date LIKE ?", getYear(year)+"%").Rows()

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		date := ""
		if err := rows.Scan(&date); err != nil {
			tx.Rollback()
			return nil, err
		}
		dates = append(dates, date)
	}
	tx.Commit()
	return dates, nil
}

func (s *Service) GetExamsResults(year int) (arrayMap, error) {
	tx := MainService.db.Begin()
	rows, err := tx.Table("exams").Select("status, date_exam").
		Where("date_exam LIKE ?", getYear(year)+"%").Rows()

	if err != nil {
		tx.Rollback()
		return nil, err
	}
	examsData := arrayMap{}
	m := m{}
	for rows.Next() {
		status := false
		date := ""
		if err := rows.Scan(&status, &date); err != nil {
			tx.Rollback()
			return nil, err
		}
		m["status"] = status
		m["date"] = date

		examsData = append(examsData, m)
	}
	tx.Commit()
	return examsData, nil
}
