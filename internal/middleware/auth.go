package middleware

import (
	"context"
	"go-minimal/internal/config"
	"go-minimal/internal/utils"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(config.GetJwtSecretKey())// later move to env variable
type contextKey string

const UserIDKey contextKey = "user_id"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			utils.WriteError(w,http.StatusUnauthorized,"No token Provided","Unauthorized")
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			utils.WriteError(w,http.StatusUnauthorized,"Invalid token","Unauthorized")
			return
		}

		// Extract user_id and put in context
		claims := token.Claims.(jwt.MapClaims)
		userID := claims["user_id"]

		ctx := context.WithValue(r.Context(), UserIDKey, userID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
