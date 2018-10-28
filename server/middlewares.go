package server

import (
	"net/http"
	log "github.com/sirupsen/logrus"
)

func withTracing(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(log.Fields{
			"addr": r.RemoteAddr,
			"method": r.Method,
			"uri": r.URL.Path,
		}).Info("Tracing request")
		next.ServeHTTP(w, r)
	}
}
