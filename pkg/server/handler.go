package server

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"

	"github.com/cfx-cv/dijkstra/pkg/dijkstra"
)

func (s *Server) distance(w http.ResponseWriter, r *http.Request) {
	d := dijkstra.NewDijkstra(s.store)
	origin, destination, apiKey := parseURL(r.URL.Query())

	result, err := d.FindDistanceAndDuration(origin, destination, apiKey)
	if err != nil {
		log.Print(err)
		return
	}

	if err = json.NewEncoder(w).Encode(result); err != nil {
		log.Print(err)
		return
	}
}

func parseURL(u url.Values) (origin, destination, apiKey string) {
	origin = u.Get("origin")
	destination = u.Get("destination")
	apiKey = u.Get("api_key")
	return
}
