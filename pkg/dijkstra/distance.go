package dijkstra

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
)

type DistanceDuration struct {
	Distance `json:"distance"`
	Duration `json:"duration"`
}

type Distance struct {
	Text  string `json:"text"`
	Value int64  `json:"value"`
}

type Duration struct {
	Text  string `json:"text"`
	Value int64  `json:"value"`
}

const (
	distanceMatrixURL = "https://maps.googleapis.com/maps/api/distancematrix/json?"
)

func (d *Dijkstra) FindDistanceAndDuration(origin, destination, apiKey string) (*DistanceDuration, error) {
	key := generateDistanceKey(origin, destination)
	if value, ok := d.store.Get(key); ok {
		return value, nil
	}

	url := buildDistanceMatrixURL(origin, destination, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	result, err := parseDistanceMatrixResponse(resp)
	if err != nil {
		return nil, err
	}
	defer d.store.Set(key, result)
	return result, nil
}

func buildDistanceMatrixURL(origin, destination, apiKey string) string {
	u := url.Values{}
	u.Add("origins", origin)
	u.Add("destinations", destination)
	u.Add("key", apiKey)
	return distanceMatrixURL + u.Encode()
}

func parseDistanceMatrixResponse(resp *http.Response) (*DistanceDuration, error) {
	var body struct {
		Rows []struct {
			Elements []struct {
				Distance `json:"distance"`
				Duration `json:"duration"`

				Status string `json:"status"`
			} `json:"elements"`
		} `json:"rows"`

		Status string `json:"status"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return nil, err
	}
	if status := body.Status; status != "OK" {
		return nil, errors.New(status)
	}
	if status := body.Rows[0].Elements[0].Status; status != "OK" {
		return nil, errors.New(status)
	}

	distance := body.Rows[0].Elements[0].Distance
	duration := body.Rows[0].Elements[0].Duration
	return &DistanceDuration{distance, duration}, nil
}
