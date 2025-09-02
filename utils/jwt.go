package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(getEnv("JWT_SECRET", "j4N9anSh@rE_t0K3NkEy!N1!123456789"))

func getEnv(key, fallback string) string {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	return v
}

// Ekspos jwtKey supaya bisa dipakai middleware
func JwtKey() []byte {
	return jwtKey
}

func GenerateToken(userID uint, role string, userName string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"user_name":userName
		"role":    role,
		"exp":     time.Now().Add(72 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
