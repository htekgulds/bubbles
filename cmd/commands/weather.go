package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// WeatherData represents the weather information
type WeatherData struct {
	Location    string
	Temperature string
	Condition   string
	Humidity    string
	WindSpeed   string
}

// WeatherResponse represents the JSON response from wttr.in API
type WeatherResponse struct {
	CurrentCondition []struct {
		TempC       string `json:"temp_C"`
		WeatherDesc []struct {
			Value string `json:"value"`
		} `json:"weatherDesc"`
		Humidity  string `json:"humidity"`
		Windspeed string `json:"windspeedKmph"`
	} `json:"current_condition"`
	NearestArea []struct {
		AreaName []struct {
			Value string `json:"value"`
		} `json:"areaName"`
		Country []struct {
			Value string `json:"value"`
		} `json:"country"`
	} `json:"nearest_area"`
}

// FetchWeather fetches weather data from a public API
// location: city name or coordinates (e.g., "London", "Istanbul", "40.7128,-74.0060")
func FetchWeather(location string) (*WeatherData, error) {
	if location == "" {
		location = "Istanbul" // default location
	}

	// Using wttr.in API - a free weather API that doesn't require an API key
	url := fmt.Sprintf("https://wttr.in/%s?format=j1", location)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch weather: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("weather API returned status: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	var weatherResp WeatherResponse
	if err := json.Unmarshal(body, &weatherResp); err != nil {
		return nil, fmt.Errorf("failed to parse weather data: %w", err)
	}

	// Extract weather information
	weather := &WeatherData{}

	if len(weatherResp.CurrentCondition) > 0 {
		current := weatherResp.CurrentCondition[0]
		weather.Temperature = current.TempC + "°C"
		weather.Humidity = current.Humidity + "%"
		weather.WindSpeed = current.Windspeed + " km/h"

		if len(current.WeatherDesc) > 0 {
			weather.Condition = current.WeatherDesc[0].Value
		}
	}

	if len(weatherResp.NearestArea) > 0 {
		area := weatherResp.NearestArea[0]
		if len(area.AreaName) > 0 && len(area.Country) > 0 {
			weather.Location = area.AreaName[0].Value + ", " + area.Country[0].Value
		}
	}

	return weather, nil
}
