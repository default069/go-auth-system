package helpers

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userID uint, role, secret string) (string, error) {
	claims := jwt.MapClaims{
		"sub":  userID,
		"role": role,
		"exp":  time.Now().Add(24 * time.Hour).Unix(),
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secret))
}