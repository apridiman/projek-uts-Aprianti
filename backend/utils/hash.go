package utils

import (
	"library-app/models"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("your_jwt_secret_key")

func GenerateJWT(user models.User) (string, error) {
	claims := &jwt.StandardClaims{
		Subject:   string(user.ID),
		ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		log.Fatal("Error generating token:", err)
		return "", err
	}

	return tokenString, nil
}
