package dijkstra

import (
	"fmt"
)

type Store interface {
	Get(key string) (*DistanceDuration, bool)
	Set(key string, value *DistanceDuration)
}

func generateKey(origin, destination Place) string {
	return fmt.Sprintf("%v:%v", origin, destination)
}
