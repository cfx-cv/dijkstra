package server

import (
	"log"
	"net/http"
)

const (
	distanceURL string = "/distance"
)

func (s *Server) Start() {
	http.HandleFunc(distanceURL, s.distance)

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
