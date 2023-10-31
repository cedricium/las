package middleware

import (
	"las_api/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := helpers.ValidateJWT(ctx)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
