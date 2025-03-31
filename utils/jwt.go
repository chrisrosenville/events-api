package utils

import (
	"errors"
	"log"
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

func VerifyToken(tokenString string) error {
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			log.Println("Unexpected signing method")
			return nil, errors.New("unexpected signing method")
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		log.Println("Error parsing token:", err)
		return errors.New("could not parse token")
	}

	tokenIsValid := parsedToken.Valid
	if !tokenIsValid {
		log.Println("Token is invalid")
		return errors.New("invalid token")
	}
	

	// claims, ok :=parsedToken.Claims.(jwt.MapClaims)
	// if !ok {
	// 	return errors.New("invalid token claims")
	// }

	// id := int64(claims["id"].(float64))
	// name := claims["name"].(string)
	// email := claims["email"].(string)

	return nil
}