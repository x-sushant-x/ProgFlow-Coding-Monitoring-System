package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/Cookie-Byte-Software/ProgFlow-Backend/types"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type UserStore interface {
	IsUserPresent(username, email string) bool
	InsertUser(ctx context.Context, user *types.User) (*types.User, error)
	GetMe(ctx context.Context, username string) (*types.LoginResponse, error)
	GetMeViaAPI(ctx *gin.Context, apiKey string) (*types.LoginResponse, error)
}

type PGUserStore struct {
	db *sql.DB
}

func NewPGUserStore(db *sql.DB) *PGUserStore {
	return &PGUserStore{
		db: db,
	}
}

func (s *PGUserStore) IsUserPresent(username, email string) bool {
	var result int

	query := `SELECT COUNT(*) FROM users WHERE username = $1 AND email = $2`

	err := s.db.QueryRow(query, username, email).Scan(&result)
	if err != nil {
		log.Fatal(err.Error())
		return false
	}

	if result == 0 {
		return false
	}
	return true
}

func (s *PGUserStore) InsertUser(ctx context.Context, user *types.User) (*types.User, error) {
	isPresent := s.IsUserPresent(user.Username, user.Email)
	fmt.Println("Value of IsPresent: ", isPresent)

	if !isPresent {
		fmt.Println("User not found Inseting")
		query := `
		INSERT INTO users (name, username, email, api_key, is_premium, photo)
		VALUES
		($1, $2, $3, $4, $5, $6)`

		_, err := s.db.Exec(query, user.Name, user.Username, user.Email, user.APIKey, user.IsPremium, user.Photo)
		if err != nil {
			return nil, fmt.Errorf(err.Error())
		}
		return user, nil
	}

	// Returning existing user
	query := `SELECT * FROM users WHERE username = $1 AND email = $2`
	err := s.db.QueryRow(query, user.Username, user.Email).Scan(&user.ID, &user.Name, &user.Username, &user.Email, &user.JoinDate, &user.IsPremium, &user.APIKey, &user.Photo)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *PGUserStore) GetMe(ctx context.Context, username, email string) (*types.LoginResponse, error) {
	var query = `SELECT * FROM users WHERE username = $1 AND email = $2`

	var user types.User

	err := s.db.QueryRow(query, username, email).Scan(&user.ID, &user.Name, &user.Username, &user.Email, &user.JoinDate, &user.IsPremium, &user.APIKey)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		} else {
			return nil, fmt.Errorf(err.Error())
		}
	}

	loggedInUser := types.MapToLoginResponse(user)

	return &loggedInUser, nil
}

func (s *PGUserStore) GetMeViaAPI(ctx *gin.Context, apiKey string) (*types.LoginResponse, error) {

	var query = `SELECT * FROM users WHERE api_key = $1`

	var user types.User

	err := s.db.QueryRow(query, apiKey).Scan(&user.ID, &user.Name, &user.Username, &user.Email, &user.JoinDate, &user.IsPremium, &user.APIKey, &user.Photo)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.Abort()
			return nil, fmt.Errorf("user not found")
		} else {
			ctx.Abort()
			return nil, fmt.Errorf(err.Error())
		}
	}
	loggedInUser := types.MapToLoginResponse(user)

	return &loggedInUser, nil
}
