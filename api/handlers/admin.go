package handlers

import (
	api "las_api"
	"las_api/helpers"
	"las_api/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AdminHandler struct {
	Store api.Store
}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *AdminHandler) Login(ctx *gin.Context) {
	var json LoginInput
	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	a, err := h.Store.AdminByUsername(json.Username)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	if err := a.ValidatePassword(json.Password); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error(), "success": false})
		return
	}

	jwt, err := helpers.GenerateJWT(a)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": true, "data": jwt})
}

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *AdminHandler) Register(ctx *gin.Context) {
	var json RegisterInput
	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	if _, err := h.Store.AdminByUsername(json.Username); err == nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": "username taken", "success": false})
		return
	}

	password, err := bcrypt.GenerateFromPassword([]byte(json.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	if err := h.Store.CreateAdmin(&models.Admin{
		ID:        "admin_" + json.Username,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Username:  json.Username,
		Password:  string(password),
	}); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"success": true, "data": nil})
}

func (h *AdminHandler) List(ctx *gin.Context) {
	aa, err := h.Store.Admins()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": true, "data": aa})
}

func (h *AdminHandler) Show(ctx *gin.Context) {
	a, err := h.Store.AdminById(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": true, "data": a})
}
