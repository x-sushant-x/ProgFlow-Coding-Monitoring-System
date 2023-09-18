package types

import (
	"fmt"
	"math/rand"
	"regexp"
	"time"

	"github.com/golang-jwt/jwt"
)

type User struct {
	ID        int
	Name      string `json:"name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	JoinDate  string `json:"joinDate"`
	IsPremium bool   `json:"isPremium"`
	APIKey    string `json:"apiKey"`
	Photo     string `json:"photo"`
}

func NewUserFromParams(params CreateUser) (*User, error) {
	currentTime := time.Now()
	formattedDate := currentTime.Format("2006-01-02")

	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	apiKey := make([]byte, 40)
	for i := range apiKey {
		apiKey[i] = charset[seededRand.Intn(len(charset))]
	}

	return &User{
		Name:      params.Name,
		Username:  params.Username,
		Email:     params.Email,
		JoinDate:  formattedDate,
		IsPremium: false,
		APIKey:    string(apiKey),
		Photo:     params.Photo,
	}, nil
}

type CreateUser struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	JoinDate string `json:"joinDate"`
	Photo    string `json:"photo" binding:"required"`
}

type LoginUser struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(e)
}

func (params CreateUser) Validate() error {
	if len(params.Name) < 8 {
		return fmt.Errorf("name must be at least 3 characters")
	}

	if !isEmailValid(params.Email) {
		return fmt.Errorf("email not valid")
	}

	return nil
}

type CustomClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type LoginResponse struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	JoinDate  string `json:"join_date"`
	IsPremium bool   `json:"is_premium"`
	APIKey    string `json:"api_key"`
	Photo     string `json:"photo"`
}

func MapToLoginResponse(user User) LoginResponse {
	return LoginResponse{
		ID:        user.ID,
		Name:      user.Name,
		Username:  user.Username,
		Email:     user.Email,
		JoinDate:  user.JoinDate,
		IsPremium: user.IsPremium,
		APIKey:    user.APIKey,
		Photo:     user.Photo,
	}
}
