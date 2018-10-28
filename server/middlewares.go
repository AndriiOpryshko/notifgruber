package server

import (
	"net/http"
	"log"
)

func withTracing(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Tracing request from %s for %s %s", r.RemoteAddr, r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	}
}
