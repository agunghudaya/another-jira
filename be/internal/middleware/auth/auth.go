package middleware

import (
	"net/http"
)

// AuthMiddleware is a placeholder for authentication middleware
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Placeholder for authentication logic

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
