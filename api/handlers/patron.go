package handlers

import (
	api "las_api"

	"github.com/gin-gonic/gin"
)

type PatronHandler struct {
	Store api.Store
}

func (h *PatronHandler) List(ctx *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (h *PatronHandler) Show(ctx *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (h *PatronHandler) Register(ctx *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (h *PatronHandler) Update(ctx *gin.Context) {
	panic("not implemented") // TODO: Implement
}
