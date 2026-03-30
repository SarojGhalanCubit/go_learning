package middleware

import (
	"context"
	"go-minimal/internal/utils"
	"net/http"
	"strings"
)

type contextKey string

const UserIDKey contextKey = "user_id"
const RoleIDKey contextKey = "role_id"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			utils.WriteError(w, http.StatusUnauthorized, "No token Provided", "Unauthorized")
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		userID, err := utils.GetUserIDFromToken(tokenString)
		if err != nil {
			utils.WriteError(w, http.StatusUnauthorized, "Invalid token", "Unauthorized")
			return
		}

		roleID, err := utils.GetRoleIDFromToken(tokenString)
		if err != nil {
			utils.WriteError(w, http.StatusUnauthorized, "Invalid token", "Unauthorized")
			return
		}
		ctx := context.WithValue(r.Context(), UserIDKey, userID)
		ctx = context.WithValue(ctx, RoleIDKey, roleID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
