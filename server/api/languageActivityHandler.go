package api

import (
	"net/http"

	"github.com/Cookie-Byte-Software/ProgFlow-Backend/db"
	"github.com/Cookie-Byte-Software/ProgFlow-Backend/types"
	"github.com/gin-gonic/gin"
)

type LanguageActivityHandler struct {
	store *db.PGLanguageStore
}

func NewLanguageActivityHandler(store *db.PGLanguageStore) *LanguageActivityHandler {
	return &LanguageActivityHandler{
		store: store,
	}
}

func (h *LanguageActivityHandler) UpdateLanguageActivity(ctx *gin.Context) {
	var params *types.LanguageActivity

	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid post data",
		})
		ctx.Abort()
		return
	}

	if err := h.store.UpdateLanguageActivity(ctx, params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Language activity updated successfully",
	})
}
