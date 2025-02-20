package utils

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var secretKey = []byte(os.Getenv("JWT_SECRET"))

type Claims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateToken(userID uint, username string) (string, error) {
	claims := Claims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		},
	)

	if err != nil {
		log.Printf("Token parsing error: %v", err)
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		log.Printf("Failed to parse claims")
		return nil, fmt.Errorf("failed to parse claims")
	}

	if claims.ExpiresAt.Time.Before(time.Now()) {
		log.Printf("Token expired")
		return nil, fmt.Errorf("token expired")
	}

	log.Printf("Token validated, claims: %+v", claims)
	return claims, nil
}
