package utils

import (
	"time"
	"github.com/golang-jwt/jwt/v5"
	"go-minimal/internal/config"
)

var jwtSecret = []byte(config.GetJwtSecretKey())


func GenerateToken(userID int) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return  token.SignedString(jwtSecret)
}
