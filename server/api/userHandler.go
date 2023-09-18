package api

import (
	"context"
	"net/http"

	"github.com/Cookie-Byte-Software/ProgFlow-Backend/db"
	"github.com/Cookie-Byte-Software/ProgFlow-Backend/types"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userStore *db.PGUserStore
}

func NewUserHandler(u *db.PGUserStore) *UserHandler {
	return &UserHandler{userStore: u}
}

func (h *UserHandler) HandleCreateUser(c *gin.Context) {
	var params *types.CreateUser

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	req, err := types.NewUserFromParams(*params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := h.userStore.InsertUser(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func (h *UserHandler) HandleUserLogin(c *gin.Context) {
	var params types.LoginUser

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid input data",
		})
		return
	}

	user, err := h.userStore.GetMe(context.Background(), params.Username, params.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
