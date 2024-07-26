package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"os"
)

var secretKey []byte = []byte(os.Getenv("JWT_SECRET_KEY"))

func CreateToken(claims jwt.MapClaims) (string, error) {
	// Create new token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token using sign method and secret key
	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(jwtString string) error {
	token, err := jwt.Parse(jwtString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return err
}
