package main

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// GenerateJWT генерирует JWT для заданного username
func GenerateJWT(username string) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "", errors.New("jwt secret not set in environment (JWT_SECRET)")
	}

	claims := jwt.MapClaims{
		"user": username,
		"exp":  jwt.NewNumericDate(time.Now().Add(15 * time.Minute)), // 15 минут
		"iat":  jwt.NewNumericDate(time.Now()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
