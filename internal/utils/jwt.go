package utils

import (
	"errors"
	"go-minimal/internal/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(config.GetJwtSecretKey())


func GenerateToken(userID,roleID int) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"role_id": roleID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return  token.SignedString(jwtSecret)
}

func GetUserIDFromToken(tokenString string) (int, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return jwtSecret, nil
    })
    if err != nil {
        return 0, err
    }

    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok || !token.Valid {
        return 0, errors.New("invalid token")
    }

    // ✅ MapClaims stores numbers as float64
    userIDFloat, ok := claims["user_id"].(float64)
    if !ok {
        return 0, errors.New("invalid user_id in token")
    }

    return int(userIDFloat), nil // ✅ convert float64 → int
}
func GetRoleIDFromToken(tokenString string) (int, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return jwtSecret, nil
    })
    if err != nil {
        return 0, err
    }

    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok || !token.Valid {
        return 0, errors.New("invalid token")
    }

    // ✅ MapClaims stores numbers as float64
    roleIDFloat, ok := claims["role_id"].(float64)
    if !ok {
        return 0, errors.New("invalid role in token")
    }

    return int(roleIDFloat), nil // ✅ convert float64 → int
}
