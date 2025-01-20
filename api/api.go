package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Response struct {
	Address           string `json:"address"`
	Timezone          string `json:"timezone"`
	Description       string `json:"description"`
	CurrentConditions struct {
		Temp float64 `json:"temp"`
	} `json:"currentConditions"`
}

func url(location string, dates []string) (string, error) {

	err := godotenv.Load()
	if err != nil {
		return "", fmt.Errorf("error occurred with loading environment variables: %w", err)
	}

	if location == "" {
		return "", fmt.Errorf("error: location for weather forecast wasn't specified")
	}

	token := os.Getenv("WEATHER_API_TOKEN")

	url := fmt.Sprintf("https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/%s/%s/%s?key=%s ", location, dates[0], dates[1], token)

	return url, nil
}

func GetForecast(location string, dates []string) (*Response, error) {
	var respObj Response
	var resp *http.Response

	if ExistsInCache(location) {
		respObj, err := GetCachedResponse(location)
		if err != nil {
			return nil, err
		}

		fmt.Printf("Temperature in %s (timezone: %s) is %.1f°C. %s\n", respObj.Address, respObj.Timezone, celsiusConverter(respObj.CurrentConditions.Temp), respObj.Description)

		return respObj, nil
	}

	url, err := url(location, dates)
	if err != nil {
		return nil, err
	}

	resp, err = http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error when fetching data from API occurred: %w", err)
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case 400:
		return nil, fmt.Errorf("invalid location name")
	case 401:
		return nil, fmt.Errorf("invalid API token")
	case 503:
		return nil, fmt.Errorf("server is unavailable")
	}

	dec := json.NewDecoder(resp.Body)
	if err = dec.Decode(&respObj); err != nil {
		return nil, fmt.Errorf("error decoding API response: %w", err)
	}

	fmt.Printf("Temperature in %s (timezone: %s) is %.1f°C. %s\n", respObj.Address, respObj.Timezone, celsiusConverter(respObj.CurrentConditions.Temp), respObj.Description)

	if err = CacheResponse(location, &respObj); err != nil {
		return nil, err
	}

	return &respObj, nil
}

func celsiusConverter(temp float64) float64 {
	return (temp - 32) / 1.8
}
