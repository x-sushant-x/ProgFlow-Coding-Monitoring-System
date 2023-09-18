/*
	Naming Conventions
		1. GetCodingTime -> Return array of coding time with given days.

		2. GetCodingStatistics -> Returns Today, Yesterday, This Month and All Time Coding Time
			int the form of 4 Hours 30 Minutes.
*/

package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/Cookie-Byte-Software/ProgFlow-Backend/types"
	"github.com/Cookie-Byte-Software/ProgFlow-Backend/utils"
	"github.com/gin-gonic/gin"
)

type AnalyticsStore interface {
	GetCodingTime(ctx *gin.Context) (*types.CodingTimeInDays, error)
	GetCodingStatistics(ctx *gin.Context, username string) (*types.CodingStatisticsResponse, error)
	GetLanguageTime(ctx *gin.Context, username string, days int) (*types.LanguageTimeResponse, error)
	GetAverageTime(username string) (*types.AverageTime, error)
	GetProjectTime(username string) ([]*types.ProjectTime, error)
	GetLeaderBoard() ([]types.Leaderboard, error)
}

type PGAnalyticsStore struct {
	db *sql.DB
}

func NewPGAnalyticsStore(db *sql.DB) *PGAnalyticsStore {
	return &PGAnalyticsStore{
		db: db,
	}
}

func (s *PGAnalyticsStore) GetCodingTime(ctx *gin.Context, username string, days int) ([]types.CodingTimeInDays, error) {
	var result []types.CodingTimeInDays

	query := `
		SELECT username, created_at, SUM(duration) FROM coding_activities WHERE username = $1 GROUP BY username, created_at ORDER BY created_at DESC LIMIT($2);
	`

	rows, err := s.db.Query(query, username, days)
	if err != nil {
		ctx.Abort()
		return nil, fmt.Errorf("no data found")
	}

	for rows.Next() {
		var row types.CodingTimeInDays

		err := rows.Scan(&row.Username, &row.Date, &row.TotalTime)
		if err != nil {
			ctx.Abort()
			return nil, fmt.Errorf(err.Error())
		}

		row.Duration = utils.GetHoursAndMins(row.TotalTime)
		result = append(result, row)
	}

	return result, nil
}

// func (s *PGAnalyticsStore) GetCodingTimeRangePerProject(ctx *gin.Context) ([]types.CodingTimeRangePerProject, error) {
// 	var result []types.CodingTimeRangePerProject
// 	username, _ := ctx.Get("username")

// 	query := `
// 		SELECT DISTINCT project_name, start_time, end_time, created_at
// 		FROM coding_activities
// 		WHERE username = $1
// 		ORDER BY created_at DESC
// 		LIMIT 30;
// 	`

// 	rows, err := s.db.Query(query, username)
// 	if err != nil {
// 		ctx.Abort()
// 		return nil, err
// 	}

// 	for rows.Next() {
// 		var row types.CodingTimeRangePerProject
// 		err := rows.Scan(&row.ProjectName, &row.StartTime, &row.EndTime, &row.CreatedAt)
// 		if err != nil {
// 			ctx.Abort()
// 			return nil, err
// 		}
// 		result = append(result, row)
// 	}

// 	return result, nil
// }

func (s *PGAnalyticsStore) GetCodingStatistics(ctx *gin.Context, username string) (*types.CodingStatisticsResponse, error) {
	today := utils.GetFormattedDate()

	query := `
	SELECT
    COALESCE(today.total_time, 0) AS today,
    COALESCE(this_week.total_time, 0) AS this_week,
    COALESCE(this_month.total_time, 0) AS this_month,
    COALESCE(all_time.total_time, 0) AS all_time
	FROM
    (SELECT SUM(duration) AS total_time FROM coding_activities WHERE username = $1 AND created_at = $2) AS today
	CROSS JOIN
    (SELECT SUM(duration) AS total_time FROM coding_activities WHERE username = $1 AND created_at BETWEEN $3 AND $2) AS this_week
	CROSS JOIN
    (SELECT SUM(duration) AS total_time FROM coding_activities WHERE username = $1 AND created_at BETWEEN $4 AND $5) AS this_month
	CROSS JOIN
    (SELECT SUM(duration) AS total_time FROM coding_activities WHERE username = $1) AS all_time;
	`

	month := utils.GetMonthStartAndEnd()
	var qResp types.CodingStatisticsResponse

	sevenDaysBackDate := utils.SubtractDate(today, 6)

	err := s.db.QueryRow(query, username, today, sevenDaysBackDate, month.StartDate, month.EndDate).Scan(&qResp.Today, &qResp.ThisWeek, &qResp.ThisMonth, &qResp.AllTime)
	if err != nil {
		ctx.Abort()
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no data found")
		}
		return nil, err
	}

	resp := types.ConvertCodingStatistics(&qResp)

	return resp, nil
}

func (s *PGAnalyticsStore) GetLanguageTime(ctx *gin.Context, username string, days int) ([]types.LanguageTimeResponse, error) {
	startDate := time.Now().AddDate(0, 0, -days).Format("2006-01-02")
	endDate := time.Now().Format("2006-01-02")

	query := `
	SELECT
    language_name,
    SUM(duration) AS total_duration
	FROM
    language_activities
	WHERE
    created_at::date BETWEEN $2 AND $3
    AND username = $1
	GROUP BY
    language_name;
	`

	rows, err := s.db.Query(query, username, startDate, endDate)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no data found")
		}
		return nil, err
	}

	var response []types.LanguageTimeResponse

	for rows.Next() {
		var singleResp types.LanguageTimeResponse
		err := rows.Scan(&singleResp.LanguageName, &singleResp.TotalDuration)
		if err != nil {
			ctx.Abort()
			return nil, err
		}

		singleResp.DurationText = utils.GetHoursAndMins(singleResp.TotalDuration)
		response = append(response, singleResp)
	}

	return response, nil
}

func (s *PGAnalyticsStore) GetAverageTime(username string) (*types.AverageTime, error) {
	query := `SELECT ROUND(COALESCE(AVG(duration), 0)) FROM coding_activities WHERE username = $1 AND created_at BETWEEN $2 AND $3`

	var resp types.AverageTime
	err := s.db.QueryRow(query, username, utils.SubtractDate(utils.GetFormattedDate(), -7), utils.SubtractDate(utils.GetFormattedDate(), -1)).Scan(&resp.AverageTime)

	if err != nil {
		fmt.Println("Error 1")
		return nil, fmt.Errorf(err.Error())
	}

	query = `SELECT ROUND(COALESCE(duration, 0)) FROM coding_activities WHERE username = $1 AND created_at = $2`

	// No need to handle error here because it will always return 0 if there is not coding time today
	s.db.QueryRow(query, username, utils.GetFormattedDate()).Scan(&resp.TodayTime)

	resp.AverageDurationText = utils.GetHoursAndMins(resp.AverageTime)
	resp.TodayDurationText = utils.GetHoursAndMins(resp.TodayTime)

	return &resp, nil
}

func (s *PGAnalyticsStore) GetProjectTime(username string) ([]types.ProjectTime, error) {
	query := `SELECT projects.project_name, SUM(duration), projects.created_at 
	FROM projects 
	INNER JOIN coding_activities ON projects.project_name = coding_activities.project_name 
	WHERE projects.username = $1 
	AND projects.created_at 
	BETWEEN $2 AND $3 GROUP BY projects.project_name, projects.created_at`

	var resp []types.ProjectTime

	rows, err := s.db.Query(query, username, utils.SubtractDate(utils.GetFormattedDate(), 6), utils.GetFormattedDate())
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no data found")
		}
		return nil, err
	}

	for rows.Next() {
		var singleResp types.ProjectTime
		rows.Scan(&singleResp.ProjectName, &singleResp.TotalTime, &singleResp.CreatedAt)
		singleResp.TotalTimeDuration = utils.GetHoursAndMins(singleResp.TotalTime)

		resp = append(resp, singleResp)

		fmt.Println(singleResp)
	}

	return resp, nil
}

func (s *PGAnalyticsStore) GetLeaderBoard() ([]types.Leaderboard, error) {
	query := `
	SELECT 
	RANK() OVER (ORDER BY SUM(duration) DESC) AS rank,
	username, 
	STRING_AGG(DISTINCT language_name, ', ' ORDER BY language_name) AS languages, 
	SUM(duration) AS total_duration 
  	FROM language_activities
  	WHERE 
	TO_DATE(created_at, 'YYYY-MM-DD') >= DATE_TRUNC('week', CURRENT_DATE)
	AND TO_DATE(created_at, 'YYYY-MM-DD') < DATE_TRUNC('week', CURRENT_DATE) + INTERVAL '1 week'
  	GROUP BY username
  	ORDER BY SUM(duration) DESC
	`

	rows, err := s.db.Query(query)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no data found")
		}
		return nil, err
	}

	var leaderboard []types.Leaderboard

	for rows.Next() {
		var singleResp types.Leaderboard
		rows.Scan(&singleResp.Rank, &singleResp.Username, &singleResp.Languages, &singleResp.TotalDuration)

		singleResp.WeekDurationText = utils.GetHoursAndMins(singleResp.TotalDuration)
		singleResp.DailyDuration = utils.GetHoursAndMins(singleResp.TotalDuration / 7)

		leaderboard = append(leaderboard, singleResp)
	}

	return leaderboard, nil
}
