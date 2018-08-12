package trail

type Distance struct {
	Distance int64 `json:"distance"`
	Duration int64 `json:"duration"`
}

func CalculateDistance(origin, destination Place, apiKey string) *Distance {
	return &Distance{} // TODO
}
