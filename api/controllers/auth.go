package controllers

import (
	"las_api/helpers"
	"las_api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthLoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthRegisterInput struct {
	AuthLoginInput
	FirstName string `json:"first_name" binding:"required"`
}

func Register(ctx *gin.Context) {
	var input AuthRegisterInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	admin := models.Admin{
		FirstName: input.FirstName,
		Username:  input.Username,
		Password:  input.Password,
	}

	savedAdmin, err := admin.Save()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"admin": savedAdmin})
}

func Login(ctx *gin.Context) {
	var input AuthLoginInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	admin, err := models.FindAdminByUsername(input.Username)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	err = admin.ValidatePassword(input.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jwt, err := helpers.GenerateJWT(admin)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"jwt": jwt})
}
