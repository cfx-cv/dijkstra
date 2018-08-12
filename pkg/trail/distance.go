package trail

import (
	"encoding/json"
	"errors"
	"fmt"
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

func FindDistanceAndDuration(origin, destination Place, apiKey string) (*DistanceDuration, error) {
	url := buildDistanceMatrixURL(origin, destination, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	distance, duration, err := parseDistanceMatrixResponse(resp)
	if err != nil {
		return nil, err
	}

	return &DistanceDuration{*distance, *duration}, nil
}

func buildDistanceMatrixURL(origin, destination Place, apiKey string) string {
	u := url.Values{}
	u.Add("origins", fmt.Sprintf("%f,%f", origin.Latitude, origin.Longitude))
	u.Add("destinations", fmt.Sprintf("%f,%f", destination.Latitude, destination.Longitude))
	u.Add("key", apiKey)
	return distanceMatrixURL + u.Encode()
}

func parseDistanceMatrixResponse(resp *http.Response) (*Distance, *Duration, error) {
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
		return nil, nil, err
	}
	if status := body.Status; status != "OK" {
		return nil, nil, errors.New(status)
	}
	if status := body.Rows[0].Elements[0].Status; status != "OK" {
		return nil, nil, errors.New(status)
	}

	distance := &body.Rows[0].Elements[0].Distance
	duration := &body.Rows[0].Elements[0].Duration
	return distance, duration, nil
}
