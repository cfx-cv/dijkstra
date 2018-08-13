package dijkstra

type Dijkstra struct {
	store Store
}

func NewDijkstra(store Store) *Dijkstra {
	return &Dijkstra{store: store}
}
