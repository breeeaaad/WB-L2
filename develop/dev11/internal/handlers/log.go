package handlers

import (
	"log"
	"net/http"
	"time"
)

func Log(h http.Handler) http.Handler {
	return http.HandlerFunc(func(c http.ResponseWriter, w *http.Request) {
		log.Printf("method: %s | uri: %s | time: %s", w.Method, w.RequestURI, time.Now())
		h.ServeHTTP(c, w)
	})
}
