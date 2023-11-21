package handlers

import (
	api "las_api"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UtilHandler struct {
	Store api.Store
}

func (h UtilHandler) CheckHealth(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{})
}
