package main

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"io/ioutil"
)

func Logger(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.SetFormatter(&log.JSONFormatter{})

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("Error reading body: %v", err)
			http.Error(w, "can't read body", http.StatusBadRequest)
			return
		}

		log.WithFields(log.Fields{
			"method": r.Method,
			"uri":  r.RequestURI,
			"data": string(body),
		}).Info("Request Handled")

		fn(w, r)
	}
}
