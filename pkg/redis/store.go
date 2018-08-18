package redis

import (
	"encoding/json"
	"log"
	"time"

	"github.com/cfx-cv/dijkstra/pkg/dijkstra"
	"github.com/cfx-cv/herald/pkg/common"
	"github.com/go-redis/redis"
)

type Store struct {
	client     *redis.Client
	expiration time.Duration
}

func NewStore(client *redis.Client, expiration time.Duration) *Store {
	return &Store{client: client, expiration: expiration}
}

func (s *Store) Get(key string) (*dijkstra.DistanceDuration, bool) {
	if value, err := s.client.Get(key).Result(); err == nil {
		var result *dijkstra.DistanceDuration
		json.Unmarshal([]byte(value), &result)
		return result, true
	}
	return nil, false
}

func (s *Store) Set(key string, value *dijkstra.DistanceDuration) {
	if j, err := json.Marshal(value); err == nil {
		s.client.Set(key, j, s.expiration)
	} else {
		log.Print(err)
		common.Publish(common.DijkstraErrors, err.Error())
	}
}
