package helpers

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var secret = "itsveryverysecret"

func CreateToken(id uint) string {

	// jwtExpired := time.Now().Local().Add(time.Minute * time.Duration(60))

	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Local().Add(time.Minute * time.Duration(60)).Unix(),
		Issuer:    "bagusnurhuda-mygram",
		Subject:   strconv.Itoa(int(id)),
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := parseToken.SignedString([]byte(secret))

	return signedToken
}

func VerifyToken(c *gin.Context) (interface{}, error) {
	errResponse := errors.New("Sign in to proceed")
	headerToken := c.Request.Header.Get("Authorization")
	bearer := strings.HasPrefix(headerToken, "Bearer")

	if !bearer {
		return nil, errResponse
	}

	stringToken := strings.Split(headerToken, " ")[1]

	token, _ := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Invalid token signing method")
		}
		return []byte(secret), nil
	})
	data, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return nil, errors.New("invalid token")
	}

	now := time.Now().Unix()
	if exp, ok := data["exp"].(float64); ok {
		if now > int64(exp) {
			return nil, errors.New("token is expired, please re-login to refresh token")
		}
	}

	id, err := strconv.Atoi(data["sub"].(string))
	fmt.Println(data["exp"])

	if err != nil {
		return 0, err
	}

	return uint(id), nil
}
