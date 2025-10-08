package main

import (
	"github.com/dgrijalva/jwt-go"
)

// Генерация JWT
func GenerateJWT() (string, error) {
	claims := jwt.MapClaims{
		"user": "john_doe",
		"exp":  15000, // Время жизни токена
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("secret"))
}
