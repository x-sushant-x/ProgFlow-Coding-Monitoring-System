package types

type CodingActivity struct {
	ProjectName string `json:"projectName"`
	StartTime   string `json:"startTime"`
	EndTime     string `json:"endTime"`
}

type CodingActivityResponse struct {
	ID          int
	Username    string `json:"username"`
	ProjectName string `json:"projectName"`
	Duration    int    `json:"duration"`
	StartTime   string `json:"startTime"`
	EndTime     string `json:"endTime"`
	CreatedAt   string `json:"createdAt"`
}
