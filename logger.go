package main

import (
	log "github.com/sirupsen/logrus"
	"net/http"
)

func Logger(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.SetFormatter(&log.JSONFormatter{})

		log.WithFields(log.Fields{
			"method": r.Method,
			"uri":  r.RequestURI,
		}).Info("Request Handled")

		fn(w, r)
	}
}