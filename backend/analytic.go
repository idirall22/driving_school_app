package service

import (
	"fmt"
	"strconv"
	"time"
)

// Check if year is valide an return a year
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

// Analytics get analytics from a table, year and column provaided
// example: table: "students", column: "registred_date", and year: 2011
func (s *Service) Analytics(table, column string,
	year int) ([]string, error) {

	tx := MainService.db.Begin()
	query := fmt.Sprintf("%s like %s", column, "'"+getYear(year)+"%"+"'")
	rows, err := tx.
		Table(table).
		Select(column).
		Where(query).
		Rows()

	if err != nil {
		tx.Rollback()
		return nil, err
	}
	defer rows.Close()

	dates := []string{}
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
