package middleware

import (
	"net/http"
	"os"
)

func BasicAuth(h http.Handler) http.Handler {
	UserID := os.Getenv("BASIC_AUTH_USER_ID")
	Password := os.Getenv("BASIC_AUTH_PASSWORD")

	fn := func(w http.ResponseWriter, r *http.Request) {
		userID, password, ok := r.BasicAuth()
		if !ok || userID != UserID || password != Password {
			w.Header().Set("WWW-Authenticate", `Basic realm="auth area"`)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
