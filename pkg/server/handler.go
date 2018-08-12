package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/cfx-cv/trail/pkg/trail"
)

type DistanceRequest struct {
	Origin      trail.Place `json:"origin"`
	Destination trail.Place `json:"destination"`

	APIKey string `json:"api_key"`
}

func (s *Server) distance(w http.ResponseWriter, r *http.Request) {
	var request DistanceRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Print(err)
		return
	}

	key := fmt.Sprintf("%v:%v", request.Origin, request.Destination)
	if value, err := s.client.Get(key).Result(); err == nil {
		if err := json.NewEncoder(w).Encode(value); err != nil {
			return
		}

		log.Print(err)
	}

	result, err := trail.FindDistanceAndDuration(request.Origin, request.Destination, request.APIKey)
	if err != nil {
		log.Print(err)
		return
	}

	s.client.Set(key, result, s.expiration).Err()
	if err := json.NewEncoder(w).Encode(result); err != nil {
		log.Print(err)
		return
	}
}
