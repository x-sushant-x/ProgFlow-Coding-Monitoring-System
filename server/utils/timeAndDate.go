package utils

import (
	"fmt"
	"time"
)

func GetFormattedDate() string {
	currentTime := time.Now()
	return currentTime.Format("2006-01-02")
}

func GetFormattedTime() string {
	currentTime := time.Now()
	return currentTime.Format("15:04")
}

func CalculateDuration(sTime, eTime string) (int, error) {
	layout := "15:04"

	startTime, err := time.Parse(layout, sTime)
	if err != nil {
		return 0, fmt.Errorf("error parsing start time: %v", err)

	}

	endTime, err := time.Parse(layout, eTime)
	if err != nil {
		return 0, fmt.Errorf("error parsing end time: %v", err)

	}

	duration := endTime.Sub(startTime).Minutes()

	return int(duration), nil
}

func GetHoursAndMins(minutes int) string {
	hours := minutes / 60
	mins := minutes % 60

	return fmt.Sprintf("%d Hours %d Minutes", hours, mins)
}

func SubtractDate(date string, days int) string {
	parsedDate, _ := time.Parse("2006-01-02", date)
	return parsedDate.AddDate(0, 0, -days).Format("2006-01-02")
}

type Month struct {
	StartDate string
	EndDate   string
}

func GetMonthStartAndEnd() Month {
	monthToDays := map[string]int{
		"January":   31,
		"February":  28,
		"March":     31,
		"April":     30,
		"May":       31,
		"June":      30,
		"July":      31,
		"August":    31,
		"September": 30,
		"October":   31,
		"November":  30,
		"December":  31,
	}

	month := time.Now().Month()

	year := time.Now().Year()

	startDate := fmt.Sprintf("%d-%02d-01", year, month)
	endDate := fmt.Sprintf("%d-%02d-%d", year, month, monthToDays[month.String()])

	var result Month

	result.StartDate = startDate
	result.EndDate = endDate
	return result
}
