package dijkstra

import (
	"fmt"
)

type Store interface {
	Get(key string) (*DistanceDuration, bool)
	Set(key string, value *DistanceDuration)
}

func generateDistanceKey(origin, destination string) string {
	return fmt.Sprintf("distance:%s:%s", origin, destination)
}
