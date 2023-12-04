package middleware

import (
	"database/sql"
	"net/http"

	"github.com/Cookie-Byte-Software/ProgFlow-Backend/db"
	"github.com/gin-gonic/gin"
)

func CheckAPIKey(pgConn *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("x-api-key")

		if apiKey == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "API Key is missing",
			})
			c.Abort()
			return
		}

		uStore := db.NewPGUserStore(pgConn)

		user, err := uStore.GetMeViaAPI(c, apiKey)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.Set("username", user.Username)
		c.Set("isPremium", user.IsPremium)
	}
}
