package utils

import (
	"time"
	"github.com/dgrijalva/jwt-go"
)

var JwtSecretKey = []byte("your_secret_key") // Sesuaikan dengan SECRET_KEY di .env

func GenerateToken(email string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtSecretKey)
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return JwtSecretKey, nil
	})
}
