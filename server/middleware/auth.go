package middleware

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Cookie-Byte-Software/ProgFlow-Backend/types"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

func CheckLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtToken := c.GetHeader("Authorization")

		if jwtToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "You are not logged in.",
			})
			c.Abort()
			return
		}

		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		secretKey := []byte(os.Getenv("JWT_SECRET"))

		token, err := jwt.ParseWithClaims(jwtToken, &types.CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(*types.CustomClaims)

		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		c.Set("username", claims.Username)
		c.Next()
	}
}

func GenerateJWTToken(username string) (string, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 48).Unix()

	secretKey := []byte(os.Getenv("JWT_SECRET"))

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
