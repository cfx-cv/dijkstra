package server

import (
	"log"
	"net/http"
	"time"

	"github.com/cfx-cv/dijkstra/pkg/dijkstra"
	dredis "github.com/cfx-cv/dijkstra/pkg/redis"
	"github.com/go-redis/redis"
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
	http.HandleFunc(distanceURL, s.distance)

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
