package service

import (
	"testing"
	"time"
)

func testGetWinLicencePerYear(t *testing.T) {
	year := 2011
	dates, err := MainService.GetWinLicencePerYear(year)
	if err != nil {
		t.Fatal(err)
	}
	students := openJSONStudentFile("")
	count := 0
	for _, std := range students {
		if val, ok := std["win_date"]; ok {
			date, _ := time.Parse(time.RFC3339, val.(string))
			if year == date.Year() {
				count++
			}
		}
	}
	if count != len(dates) {
		t.Error("There is an error this should be equal")
	}
}

func testGetExamsResults(t *testing.T) {
	_, err := MainService.GetExamsResults(2019)

	if err != nil {
		t.Fatal(err)
	}
}
