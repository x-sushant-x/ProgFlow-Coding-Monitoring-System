package api

import (
	"net/http"

	"github.com/Cookie-Byte-Software/ProgFlow-Backend/db"
	"github.com/Cookie-Byte-Software/ProgFlow-Backend/types"
	"github.com/gin-gonic/gin"
)

type ProjectHandler struct {
	projectStore *db.PGProjectStore
}

func NewProjectHandler(u *db.PGProjectStore) *ProjectHandler {
	return &ProjectHandler{projectStore: u}
}

func (h *ProjectHandler) HandleAddProject(ctx *gin.Context) {
	var params *types.ProjectAdd

	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid input data",
		})
		ctx.Abort()
		return
	}

	if err := h.projectStore.AddProject(ctx, params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "project added successfully",
	})
}
