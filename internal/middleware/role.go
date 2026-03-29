package middleware

import (
	"go-minimal/internal/utils"
	"log"
	"net/http"
)

func RequireRole(allowedRoles ...int) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			roleID, ok := r.Context().Value(RoleIDKey).(int)
			if !ok {
				utils.WriteError(w, http.StatusForbidden, "Request Failed", "Role not found")
				return
			}

			log.Println("ROLE ID :: ", roleID)

			for _, role := range allowedRoles {
				if roleID == role {
					next.ServeHTTP(w, r)
					return
				}
			}

			utils.WriteError(w, http.StatusForbidden, "your current role cannot access this feature", "Permission denied")
		})
	}
}
