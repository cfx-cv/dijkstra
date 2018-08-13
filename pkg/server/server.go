package server

import (
	"log"
	"net/http"
	"time"

	"github.com/go-redis/redis"
)

const (
	distanceURL string = "/distance"
)

type Server struct {
	// redis
	client     *redis.Client
	expiration time.Duration
}

func NewServer(client *redis.Client, expiration time.Duration) *Server {
	return &Server{client: client, expiration: expiration}
}

func (s *Server) Start() {
	http.HandleFunc(distanceURL, s.distance)

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
