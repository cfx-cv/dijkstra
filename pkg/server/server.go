package server

import (
	"time"

	"github.com/go-redis/redis"
)

type Server struct {
	// redis
	client     *redis.Client
	expiration time.Duration
}

func NewServer(client *redis.Client, expiration time.Duration) *Server {
	return &Server{client: client, expiration: time.Duration(expiration)}
}
