package service

import (
	"testing"
	"time"
)

// Test Analytics
func testAnalytics(t *testing.T) {
	year := 2011
	dates, err := MainService.Analytics(
		"students",
		"win_licence_date",
		year,
	)

	if err != nil {
		t.Fatal(err)
	}
	students := openJSONStudentFile("")
	count := 0
	for _, std := range students {
		if val, ok := std["win_licence_date"]; ok {
			date, _ := time.Parse(time.RFC3339, val.(string))
			if year == date.Year() {
				count++
			}
		}
	}
	if count != len(dates) {
		t.Errorf("There is an error this should be equal but found %d != %d",
			count, len(dates))
	}
}
