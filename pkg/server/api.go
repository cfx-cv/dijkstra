package server

import (
	"log"
	"net/http"
)

const (
	distanceURL string = "/distance"
)

// Start initializes API endpoints
func Start(port string, apiKey string) {
	env := Env{}
	http.HandleFunc(distanceURL, env.distance)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
