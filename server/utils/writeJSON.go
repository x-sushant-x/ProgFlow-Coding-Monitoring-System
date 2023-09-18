package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func WriteJSON(c *gin.Context, statusCode int, body any) {
	c.JSON(http.StatusBadRequest, body)
	return
}
