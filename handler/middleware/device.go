package middleware

import (
	"context"
	"log"
	"net/http"

	"github.com/mileusna/useragent"
)

func Device(h http.Handler) http.Handler {
	type osContextKey string
	k := osContextKey("os")

	fn := func(w http.ResponseWriter, r *http.Request) {
		ua := useragent.Parse(r.UserAgent())
		ctx := context.WithValue(r.Context(), k, ua.OS)
		log.Println(ctx.Value(k))
		h.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}
