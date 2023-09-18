package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Cookie-Byte-Software/ProgFlow-Backend/types"
	"github.com/gin-gonic/gin"
)

type ProjectStore interface {
	AddProject(c context.Context, project *types.ProjectAdd) error
}

type PGProjectStore struct {
	db *sql.DB
}

func NewPGProjectStore(db *sql.DB) *PGProjectStore {
	return &PGProjectStore{
		db: db,
	}
}

func (s *PGProjectStore) AddProject(ctx *gin.Context, project *types.ProjectAdd) error {
	var count int
	query := `
		SELECT COUNT(*) FROM projects WHERE project_name = $1 AND username = $2
	`

	username, _ := ctx.Get("username")

	err := s.db.QueryRowContext(ctx, query, project.Name, username).Scan(&count)
	if err != nil {
		ctx.Abort()
		return err
	}

	if count > 0 {
		ctx.Abort()
		return fmt.Errorf("project already exists")
	}

	query = `
		INSERT INTO projects (project_name, username, created_at) 
		VALUES 
		($1, $2, $3)
	`

	currentTime := time.Now()
	formattedDate := currentTime.Format("2006-01-02")

	_, err = s.db.Exec(query, project.Name, username, formattedDate)
	if err != nil {
		ctx.Abort()
		return fmt.Errorf(err.Error())
	}
	return nil
}
