package helpers

import (
	"errors"
	"fmt"
	"las_api/models"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var privateKey = []byte(os.Getenv("JWT_PRIVATE_KEY"))

func GenerateJWT(admin models.Admin) (string, error) {
	ttl, _ := strconv.Atoi(os.Getenv("TOKEN_TTL"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  admin.ID,
		"iat": time.Now().Unix(),
		"eat": time.Now().Add(time.Second * time.Duration(ttl)).Unix(),
	})
	return token.SignedString(privateKey)
}

func ValidateJWT(ctx *gin.Context) error {
	token, err := getToken(ctx)
	if err != nil {
		return err
	}

	_, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return nil
	}
	return errors.New("invalid token provided")
}

func CurrentAdmin(ctx *gin.Context) (models.Admin, error) {
	err := ValidateJWT(ctx)
	if err != nil {
		return models.Admin{}, err
	}

	token, _ := getToken(ctx)
	claims, _ := token.Claims.(jwt.MapClaims)
	adminId := uint(claims["id"].(float64))

	admin, err := models.FindAdminById(adminId)
	if err != nil {
		return models.Admin{}, err
	}
	return admin, nil
}

func getToken(ctx *gin.Context) (*jwt.Token, error) {
	tokenStr := getTokenFromRequest(ctx)
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return privateKey, nil
	})
	return token, err
}

func getTokenFromRequest(ctx *gin.Context) string {
	bearerToken := ctx.Request.Header.Get("Authorization")
	splitToken := strings.Split(bearerToken, " ")
	if len(splitToken) == 2 {
		return splitToken[1]
	}
	return ""
}
