package types

type LanguageActivity struct {
	ProjectName  string `json:"projectName"`
	LanguageName string `json:"languageName"`
	StartTime    string `json:"startTime"`
}

type LanguageActivityResponse struct {
	ID           int
	Username     string `json:"username"`
	ProjectName  string `json:"projectName"`
	LanguageName string `json:"languageName"`
	StartTime    string `json:"startTime"`
	EndTime      string `json:"endTime"`
	Duration     int    `json:"duration"`
	CreatedAt    string `json:"createdAt"`
}
