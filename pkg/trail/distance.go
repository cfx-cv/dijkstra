package trail

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

type Distance struct {
	Distance int64 `json:"distance"`
	Duration int64 `json:"duration"`
}

const (
	distanceURL = "https://maps.googleapis.com/maps/api/distancematrix/json?"
)

func FindDistance(origin, destination Place, apiKey string) (*Distance, error) {
	url := buildDistanceURL(origin, destination, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	distance, duration, err := parseDistanceResponse(resp)
	if err != nil {
		return nil, err
	}

	return &Distance{distance, duration}, nil
}

func buildDistanceURL(origin, destination Place, apiKey string) string {
	u := url.Values{}
	u.Add("origins", fmt.Sprintf("%f,%f", origin.Latitude, origin.Longitude))
	u.Add("destinations", fmt.Sprintf("%f,%f", destination.Latitude, destination.Longitude))
	u.Add("key", apiKey)
	return distanceURL + u.Encode()
}

func parseDistanceResponse(resp *http.Response) (int64, int64, error) {
	var body struct {
		Rows []struct {
			Elements []struct {
				Distance struct {
					Value int64 `json:"value"`
				} `json:"distance"`
				Duration struct {
					Value int64 `json:"value"`
				} `json:"duration"`

				Status string `json:"status"`
			} `json:"elements"`
		} `json:"rows"`

		Status string `json:"status"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return 0, 0, err
	}
	if status := body.Status; status != "OK" {
		return 0, 0, errors.New(status)
	}
	if status := body.Rows[0].Elements[0].Status; status != "OK" {
		return 0, 0, errors.New(status)
	}

	distance := body.Rows[0].Elements[0].Distance.Value
	duration := body.Rows[0].Elements[0].Duration.Value
	return distance, duration, nil
}
