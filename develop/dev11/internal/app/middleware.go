package app

import (
	"log"
	"net/http"
)

// logEvent func - middleware for http routes, which will log incoming request method and URI to the console
func logEvent(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("[Info] request method: "+r.Method, "| URI: "+r.RequestURI)
		next.ServeHTTP(w, r)
	}
}
