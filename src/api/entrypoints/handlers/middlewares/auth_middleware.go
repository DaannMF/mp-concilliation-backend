package middlewares

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/proethics/mp-conciliation/src/api/core/entities"
	"gorm.io/gorm"
)

type Auth struct {
	StoreClient *gorm.DB
}

func (middleware Auth) Handle(c *gin.Context) {

	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		c.AbortWithError(http.StatusUnauthorized, errors.New("authorization header is missing"))
		return
	}

	authToken := strings.Split(authHeader, " ")
	if len(authToken) != 2 || authToken[0] != "Bearer" {
		c.AbortWithError(http.StatusUnauthorized, errors.New("invalid token format"))
		return
	}

	tokenString := authToken[1]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil || !token.Valid {
		c.AbortWithError(http.StatusUnauthorized, errors.New("invalid or expired token"))
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.AbortWithError(http.StatusUnauthorized, errors.New("invalid token"))
		return
	}

	if float64(time.Now().Unix()) > claims["exp"].(float64) {
		c.AbortWithError(http.StatusUnauthorized, errors.New("token expired"))
		return
	}

	var user entities.User
	middleware.StoreClient.First(&user, claims["id"])

	if user.ID == 0 {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Set("currentUser", user)

	c.Next()
}
