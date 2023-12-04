package db

import (
	"database/sql"

	"github.com/Cookie-Byte-Software/ProgFlow-Backend/types"
	"github.com/Cookie-Byte-Software/ProgFlow-Backend/utils"
	"github.com/gin-gonic/gin"
)

type LanguageStore interface {
	UpdateLanguageActivity(ctx *gin.Context, activity *types.LanguageActivity) error
}

type PGLanguageStore struct {
	db *sql.DB
}

func NewPGLanguageStore(db *sql.DB) *PGLanguageStore {
	return &PGLanguageStore{
		db: db,
	}
}

func (s *PGLanguageStore) UpdateLanguageActivity(ctx *gin.Context, activity *types.LanguageActivity) error {
	var count int
	query := `
		SELECT COUNT(*) FROM language_activities 
		WHERE username = $1 AND project_name = $2 AND language_name = $3 AND start_time = $4 AND created_at = $5
	`

	username, _ := ctx.Get("username")

	err := s.db.QueryRowContext(ctx, query, username, activity.ProjectName, activity.LanguageName, activity.StartTime, utils.GetFormattedDate()).Scan(&count)
	if err != nil {
		return err
	}

	if count == 0 {
		query = `
			INSERT INTO language_activities (username, project_name, language_name, start_time, end_time, duration, created_at)
		 	VALUES($1, $2, $3, $4, $5, $6, $7)`

		_, err := s.db.Exec(query, username, activity.ProjectName, activity.LanguageName, activity.StartTime, utils.GetFormattedTime(), 0, utils.GetFormattedDate())
		if err != nil {
			return err
		}
	} else {
		var existingActivity types.LanguageActivityResponse

		query = `
			SELECT * FROM language_activities WHERE 
			username = $1 AND project_name = $2 AND language_name = $3 AND start_time = $4 AND created_at = $5`

		err := s.db.QueryRow(query, username, activity.ProjectName, activity.LanguageName, activity.StartTime, utils.GetFormattedDate()).Scan(&existingActivity.ID, &existingActivity.Username, &existingActivity.ProjectName, &existingActivity.LanguageName, &existingActivity.StartTime, &existingActivity.EndTime, &existingActivity.Duration, &existingActivity.CreatedAt)

		if err != nil {
			return err
		}

		duration, err := utils.CalculateDuration(existingActivity.StartTime, utils.GetFormattedTime())
		if err != nil {
			return err
		}

		query = `
			UPDATE language_activities SET end_time = $1, duration = $2 
			WHERE username = $3 AND project_name = $4 AND start_time = $5 AND created_at = $6`

		_, err = s.db.Exec(query, utils.GetFormattedTime(), duration, username, existingActivity.ProjectName, existingActivity.StartTime, utils.GetFormattedDate())
		if err != nil {
			return err
		}
	}

	return nil
}
