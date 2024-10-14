package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type AccessLog struct {
	TimeStamp time.Time `json:"time_stamp"`
	Latency   int64     `json:"latency"`
	Path      string    `json:"path"`
	OS        string    `json:"os"`
}

func Logging(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		h.ServeHTTP(w, r)
		finish := time.Now()

		urlPath := r.URL.Path

		osValue, ok := r.Context().Value(osKey).(string)
		if !ok {
			log.Println(osValue)
		}

		accessLog := &AccessLog{
			TimeStamp: start,
			Latency:   finish.Unix() - start.Unix(),
			Path:      urlPath,
			OS:        osValue,
		}

		jsonData, err := json.Marshal(accessLog)
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("%s\n", jsonData)
	}

	return http.HandlerFunc(fn)
}
