package api

import (
	"net/http"
	"strconv"

	"github.com/Cookie-Byte-Software/ProgFlow-Backend/db"
	"github.com/gin-gonic/gin"
)

type AnalyticsHandler struct {
	store *db.PGAnalyticsStore
}

func NewAnalyticsHandler(analyticsStore *db.PGAnalyticsStore) *AnalyticsHandler {
	return &AnalyticsHandler{
		store: analyticsStore,
	}
}

func (h *AnalyticsHandler) HandleGetCodingTime(ctx *gin.Context) {
	username := ctx.Query("username")
	days := ctx.Query("days")
	if days == "" || username == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "missing username or days query",
		})
		ctx.Abort()
		return
	}

	daysInt, error := strconv.Atoi(days)
	if error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid days query data",
		})
		ctx.Abort()
		return
	}

	result, err := h.store.GetCodingTime(ctx, username, daysInt)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"username": username,
		"data":     result,
	})
}

// func (h *AnalyticsHandler) HandleGetCodingTimeRangePerProject(ctx *gin.Context) {
// 	result, err := h.store.GetCodingTimeRangePerProject(ctx)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadGateway, gin.H{
// 			"error": err,
// 		})
// 		ctx.Abort()
// 		return
// 	}

// 	user, _ := ctx.Get("username")

// 	ctx.JSON(http.StatusOK, gin.H{
// 		"username": user,
// 		"data":     result,
// 	})
// }

func (h *AnalyticsHandler) HandleGetCodingStatistics(ctx *gin.Context) {
	username := ctx.Query("username")

	if username == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "missing username query",
		})
		ctx.Abort()
		return
	}

	resp, err := h.store.GetCodingStatistics(ctx, username)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": resp,
	})
}

func (h *AnalyticsHandler) HandleGetLanguageTime(ctx *gin.Context) {
	username := ctx.Query("username")
	days := ctx.Query("days")

	if username == "" || days == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "missing username or days query",
		})
		ctx.Abort()
		return
	}

	daysInt, error := strconv.Atoi(days)
	if error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid days query data",
		})
		ctx.Abort()
		return
	}

	resp, err := h.store.GetLanguageTime(ctx, username, daysInt)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": resp,
	})
}

func (h *AnalyticsHandler) HandleGetAverageTime(ctx *gin.Context) {
	username := ctx.Query("username")

	if username == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "missing username query",
		})
		ctx.Abort()
		return
	}

	resp, err := h.store.GetAverageTime(username)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": resp,
	})
}

func (h *AnalyticsHandler) HandleGetProjectTime(ctx *gin.Context) {
	username := ctx.Query("username")

	if username == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "missing username query",
		})
		ctx.Abort()
		return
	}

	resp, err := h.store.GetProjectTime(username)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": resp,
	})
}

func (h *AnalyticsHandler) HandleGetLeaderboard(ctx *gin.Context) {
	resp, err := h.store.GetLeaderBoard()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": resp,
	})
}
