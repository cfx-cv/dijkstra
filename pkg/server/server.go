package server

import (
	"log"
	"net/http"
	"time"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"

	"github.com/cfx-cv/dijkstra/pkg/dijkstra"
	dredis "github.com/cfx-cv/dijkstra/pkg/redis"
	"github.com/cfx-cv/herald/pkg/common"
)

const (
	distanceURL string = "/distance"
)

type Server struct {
	store dijkstra.Store
}

func NewServer(client *redis.Client, expiration time.Duration) *Server {
	store := dredis.NewStore(client, expiration)
	return &Server{store: store}
}

func (s *Server) Start() {
	router := mux.NewRouter()
	router.HandleFunc(distanceURL, s.distance).Methods("GET")

	err := http.ListenAndServe(":80", router)
	if err != nil {
		log.Fatal(err)
		common.Publish(common.DijkstraErrors, err.Error())
	}
}
