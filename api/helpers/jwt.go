package helpers

import (
	"errors"
	"fmt"
	api "las_api"
	"las_api/models"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var key = []byte(os.Getenv("JWT_PRIVATE_KEY"))

func GenerateJWT(a models.Admin) (string, error) {
	ttl, _ := strconv.Atoi(os.Getenv("TOKEN_TTL"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  a.ID,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Second * time.Duration(ttl)).Unix(),
	})
	return token.SignedString(key)
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

func CurrentAdmin(ctx *gin.Context, s api.Store) (models.Admin, error) {
	if err := ValidateJWT(ctx); err != nil {
		return models.Admin{}, err
	}

	token, _ := getToken(ctx)
	claims, _ := token.Claims.(jwt.MapClaims)
	id := claims["id"]

	admin, err := s.AdminById(id.(string))
	if err != nil {
		return models.Admin{}, err
	}
	return admin, nil
}

func getToken(ctx *gin.Context) (*jwt.Token, error) {
	token, err := jwt.Parse(getTokenFromRequest(ctx), func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return key, nil
	})
	return token, err
}

func getTokenFromRequest(ctx *gin.Context) string {
	bearer := ctx.Request.Header.Get("Authorization")
	parts := strings.Split(bearer, " ")
	if len(parts) == 2 {
		return parts[1]
	}
	return ""
}
