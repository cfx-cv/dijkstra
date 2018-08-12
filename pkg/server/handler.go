package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/cfx-cv/trail/pkg/trail"
)

type Env struct{}

type DistanceRequest struct {
	Origin      trail.Place `json:"origin"`
	Destination trail.Place `json:"destination"`

	APIKey string `json:"api_key"`
}

func (e *Env) distance(w http.ResponseWriter, r *http.Request) {
	var request DistanceRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Print(err)
		return
	}

	distance := trail.CalculateDistance(request.Origin, request.Destination, request.APIKey)
	fmt.Println(distance)
	err := json.NewEncoder(w).Encode(*distance)
	if err != nil {
		log.Println(err)
	}
}
