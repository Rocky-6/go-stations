package middleware

import (
	"context"
	"log"
	"net/http"

	"github.com/mileusna/useragent"
)

type osContextKey string

const osKey = osContextKey("os")

func Device(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ua := useragent.Parse(r.UserAgent())
		ctx := context.WithValue(r.Context(), osKey, ua.OS)
		log.Println(ctx.Value(osKey))
		h.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}
