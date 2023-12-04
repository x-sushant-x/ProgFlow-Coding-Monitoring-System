package types

type Leaderboard struct {
	Rank             int    `json:"rank"`
	Username         string `json:"username"`
	Languages        string `json:"languages"`
	TotalDuration    int    `json:"total_duration"`
	WeekDurationText string `json:"week_duration_text"`
	DailyDuration    string `json:"daily_duration_text"`
}
