package jwt

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

const Secret string = "secret"

func NewJWT() (string, error) {
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{}).SignedString([]byte(Secret))
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return token, nil
}

func CheckJWT(token string) (bool, error) {
	parsed, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secret), nil
	})
	if err != nil {
		return false, fmt.Errorf("failed to parse token: %w", err)
	}

	if !parsed.Valid {
		return false, fmt.Errorf("invalid token")
	}

	return true, nil
}
