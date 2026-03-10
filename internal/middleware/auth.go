package middleware

import (
	"go-minimal/internal/config"
	"go-minimal/internal/utils"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(config.GetJwtSecretKey())// later move to env variable

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

		// Optional: Extract user_id and put in context
		next.ServeHTTP(w, r)
	})
}
