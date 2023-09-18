package db

import (
	"database/sql"

	"github.com/Cookie-Byte-Software/ProgFlow-Backend/types"
	"github.com/Cookie-Byte-Software/ProgFlow-Backend/utils"
	"github.com/gin-gonic/gin"
)

type CodingActivityStore interface {
	UpdateCodingActivity(c *gin.Context, codingActivity *types.CodingActivity) error
}

type PGCodingActivityStore struct {
	db *sql.DB
}

func NewPGCodingActivityStore(db *sql.DB) *PGCodingActivityStore {
	return &PGCodingActivityStore{
		db: db,
	}
}

func (s *PGCodingActivityStore) UpdateCodingActivity(c *gin.Context, codingActivity *types.CodingActivity) error {
	var count int
	query := `
		SELECT COUNT(*) FROM coding_activities WHERE username = $1 AND project_name = $2 AND start_time = $3 AND created_at = $4
	`

	username, _ := c.Get("username")

	err := s.db.QueryRowContext(c, query, username, codingActivity.ProjectName, codingActivity.StartTime, utils.GetFormattedDate()).Scan(&count)
	if err != nil {
		return err
	}

	if count == 0 {
		query = `INSERT INTO coding_activities (username, project_name, start_time, end_time, duration, created_at) VALUES($1, $2, $3, $4, $5, $6)`

		_, err := s.db.Exec(query, username, codingActivity.ProjectName, codingActivity.StartTime, utils.GetFormattedTime(), 0, utils.GetFormattedDate())
		if err != nil {
			return err
		}
	} else {
		var existingActivity types.CodingActivityResponse

		query = `SELECT * FROM coding_activities WHERE username = $1 AND project_name = $2 AND start_time = $3 AND created_at = $4`

		s.db.QueryRow(query, username, codingActivity.ProjectName, codingActivity.StartTime, utils.GetFormattedDate()).Scan(&existingActivity.ID, &existingActivity.Username, &existingActivity.ProjectName, &existingActivity.StartTime, &existingActivity.EndTime, &existingActivity.Duration, &existingActivity.CreatedAt)

		duration, err := utils.CalculateDuration(existingActivity.StartTime, utils.GetFormattedTime())
		if err != nil {
			return err
		}

		query = `UPDATE coding_activities SET end_time = $1, duration = $2 WHERE username = $3 AND project_name = $4 AND start_time = $5 AND created_at = $6`

		_, err = s.db.Exec(query, utils.GetFormattedTime(), duration, username, codingActivity.ProjectName, codingActivity.StartTime, utils.GetFormattedDate())
		if err != nil {
			return err
		}
	}

	return nil
}
