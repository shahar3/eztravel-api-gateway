package client

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"time"
)

var httpClient = &http.Client{
	Timeout: 10 * time.Second,
}

// CreateTrip forwards the trip creation request to the Trip Planning Service.
func CreateTrip(payload map[string]interface{}) (*http.Response, error) {
	// Read the endpoint from an environment variable (or you could pass from config)
	endpoint := os.Getenv("TRIP_SERVICE_ENDPOINT")
	if endpoint == "" {
		endpoint = "http://localhost:8082/api/trip" // fallback default
	}

	data, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	return httpClient.Do(req)
}
