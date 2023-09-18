package api

import (
	"net/http"

	"github.com/Cookie-Byte-Software/ProgFlow-Backend/db"
	"github.com/Cookie-Byte-Software/ProgFlow-Backend/types"
	"github.com/gin-gonic/gin"
)

type CodingActivityHandler struct {
	db *db.PGCodingActivityStore
}

func NewCodingActivityHandler(db *db.PGCodingActivityStore) *CodingActivityHandler {
	return &CodingActivityHandler{
		db: db,
	}
}

func (h *CodingActivityHandler) HandleUpdateCodingActivity(c *gin.Context) {
	var params *types.CodingActivity

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		c.Abort()
		return
	}

	if err := h.db.UpdateCodingActivity(c, params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "coding activity updated",
	})
}
