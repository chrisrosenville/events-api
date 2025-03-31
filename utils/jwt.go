package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "supersecretkey"

func GenerateToken(id int64, name, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    id,
		"name":  name,
		"email": email,
		"exp":   time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}