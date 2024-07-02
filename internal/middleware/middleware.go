package middleware

import (
	"CRM-Service/config"
	"CRM-Service/internal/auth"
	"net/http"
	"strings"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing Authorization Header", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Invalid Authorization Header", http.StatusUnauthorized)
			return
		}

		token := parts[1]

		conf := config.GetGlobalConfig()

		if _, err := auth.ValidateJWT(conf, token); err != nil {
			http.Error(w, "Invalid Token", http.StatusUnauthorized)
			return
		}

	})
}
