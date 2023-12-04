package types

import (
	"strconv"

	"github.com/Cookie-Byte-Software/ProgFlow-Backend/utils"
)

type CodingTimeInDays struct {
	Username  string `json:"-"`
	TotalTime int    `json:"time"`
	Date      string `json:"date"`
	Duration  string `json:"duration"`
}

type CodingTimeRangePerProject struct {
	ID          int    `json:"-"`
	ProjectName string `json:"projectName"`
	StartTime   string `json:"startTime"`
	EndTime     string `json:"endTime"`
	CreatedAt   string `json:"createdAt"`
}
type CodingStatisticsResponse struct {
	Today     string `json:"today"`
	ThisWeek  string `json:"thisWeek"`
	ThisMonth string `json:"thisMonth"`
	AllTime   string `json:"allTime"`
}

func ConvertCodingStatistics(stats *CodingStatisticsResponse) *CodingStatisticsResponse {
	today, _ := strconv.Atoi(stats.Today)
	finalToday := utils.GetHoursAndMins(today)

	thisWeek, _ := strconv.Atoi(stats.ThisWeek)
	finalThisWeek := utils.GetHoursAndMins(thisWeek)

	thisMonth, _ := strconv.Atoi(stats.ThisMonth)
	finalThisMonth := utils.GetHoursAndMins(thisMonth)

	allTime, _ := strconv.Atoi(stats.AllTime)
	finalAllTime := utils.GetHoursAndMins(allTime)

	return &CodingStatisticsResponse{
		Today:     finalToday,
		ThisWeek:  finalThisWeek,
		ThisMonth: finalThisMonth,
		AllTime:   finalAllTime,
	}
}

type LanguageTimeResponse struct {
	LanguageName  string `json:"languageName"`
	TotalDuration int    `json:"totalDuration"`
	DurationText  string `json:"durationText"`
}

type AverageTime struct {
	AverageTime         int    `json:"averageTime"`
	TodayTime           int    `json:"todayTime"`
	AverageDurationText string `json:"averageDurationText"`
	TodayDurationText   string `json:"todayDurationText"`
}

type ProjectTime struct {
	ProjectName       string `json:"projectName"`
	TotalTime         int    `json:"totalTime"`
	CreatedAt         string `json:"createdAt"`
	TotalTimeDuration string `json:"totalTimeDuration"`
}
