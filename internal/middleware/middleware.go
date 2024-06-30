package middleware

import (
	"CRM-Service/internal/auth"
	"context"
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
		claims, err := auth.ValidateJWT(token)
		if err != nil {
			http.Error(w, "Invalid Token", http.StatusUnauthorized)
			return
		}

		// Set the user email in the request context
		ctx := context.WithValue(r.Context(), "email", claims.Email)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
