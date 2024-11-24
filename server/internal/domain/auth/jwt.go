package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var appName = os.Getenv("APP_NAME")
var jwtSecretKey = os.Getenv("JWT_SECRET")
var jwtSecret = []byte(jwtSecretKey)

func GenerateToken(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
		"iat":      time.Now().Unix(),
		"iss":      appName,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
