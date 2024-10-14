package middleware

import (
	"net/http"
	"os"
)

func BasicAuth(h http.Handler) http.Handler {
	UserID := os.Getenv("BASIC_AUTH_USER_ID")
	Password := os.Getenv("BASIC_AUTH_PASSWORD")

	fn := func(w http.ResponseWriter, r *http.Request) {
		userID, password, _ := r.BasicAuth()
		if userID != UserID || password != Password {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
