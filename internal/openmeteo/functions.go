package openmeteo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func createRequestUrl(lat, lon string) (string, error) {
	u, err := url.Parse("https://api.open-meteo.com/v1/forecast")
	if err != nil {
		return "", fmt.Errorf("createRequestUrl(lat, lon string): %v", err)
	}

	params := url.Values{}
	params.Add("latitude", lat)
	params.Add("longitude", lon)
	// params.Add("timeformat", "unixtime")
	// params.Add("current_weather", "true")
	params.Add("hourly", "temperature_2m")
	params.Add("wind_speed_unit", "ms")

	u.RawQuery = params.Encode()

	return u.String(), nil
}

func getWeather(lat, lon string) (Forecast, error) {
	errorPrefix := "getWeather(lat, lon string): "

	u, err := createRequestUrl(lat, lon)
	if err != nil {
		return Forecast{}, fmt.Errorf("%s%v", errorPrefix, err)
	}

	resp, err := http.Get(u)
	if err != nil {
		return Forecast{}, fmt.Errorf("%s%v", errorPrefix, err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Forecast{}, fmt.Errorf("%s%v", errorPrefix, err)
	}

	if resp.StatusCode != http.StatusOK {
		errorResp := ErrorResp{}
		err := json.Unmarshal(data, &errorResp)
		if err != nil {
			return Forecast{}, fmt.Errorf("%s%v", errorPrefix, err)
		}
		return Forecast{}, fmt.Errorf("%s%s", errorPrefix, errorResp.Reason)
	}

	forecastResp := Forecast{}
	err = json.Unmarshal(data, &forecastResp)

	if err != nil {
		return Forecast{}, fmt.Errorf("%s%v", errorPrefix, err)
	}
	return forecastResp, nil
}
