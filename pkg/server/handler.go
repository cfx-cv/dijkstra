package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/cfx-cv/dijkstra/pkg/dijkstra"
)

type DistanceRequest struct {
	Origin      dijkstra.Place `json:"origin"`
	Destination dijkstra.Place `json:"destination"`

	APIKey string `json:"api_key"`
}

func (s *Server) distance(w http.ResponseWriter, r *http.Request) {
	var request DistanceRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Print(err)
		return
	}

	d := dijkstra.NewDijkstra(s.store)
	result, err := d.FindDistanceAndDuration(request.Origin, request.Destination, request.APIKey)
	if err != nil {
		log.Print(err)
		return
	}

	if err = json.NewEncoder(w).Encode(result); err != nil {
		log.Print(err)
		return
	}
}
