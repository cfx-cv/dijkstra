package dijkstra

import (
	"fmt"
)

type Store interface {
	Get(key string) (*DistanceDuration, bool)
	Set(key string, value *DistanceDuration)
}

func generateKey(origin, destination string) string {
	return fmt.Sprintf("%s:%s", origin, destination)
}
