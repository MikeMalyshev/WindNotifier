package openmeteo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/MikeMalyshev/WindNotifier/internal/weatheragent"
)

type OpenMeteo struct{}

func New() *OpenMeteo {
	return &OpenMeteo{}
}

func (op OpenMeteo) createRequestUrl(lat, lon string) (string, error) {
	u, err := url.Parse("https://api.open-meteo.com/v1/forecast")
	if err != nil {
		return "", fmt.Errorf("createRequestUrl(lat, lon string): %v", err)
	}

	params := url.Values{}
	params.Add("latitude", lat)
	params.Add("longitude", lon)
	params.Add("wind_speed_unit", "ms")
	params.Add("temperature_unit", "celsius")

	hourlyParams := []string{
		"temperature_2m",
		"wind_speed_10m",
		"wind_direction_10m",
		"wind_gusts_10m",
		"cloud_cover_low",
		"cloud_cover_mid",
		"cloud_cover_high",
	}

	params.Add("hourly", strings.Join(hourlyParams, ","))

	u.RawQuery = params.Encode()

	return u.String(), nil
}

func (op OpenMeteo) findNearestTime(time time.Time, timeList []TimeISO8601) (int, error) {
	for idx, t := range timeList {
		if t.After(time) {
			return idx, nil
		}
	}
	return 0, fmt.Errorf("findNearestTime(timeList []time.Time): no time found")
}

func (op OpenMeteo) doRequest(lat, lon string) (Forecast, error) {
	errorPrefix := "getWeather(lat, lon string): "

	u, err := op.createRequestUrl(lat, lon)
	if err != nil {
		return Forecast{}, fmt.Errorf("%s%v", errorPrefix, err)
	}
	fmt.Println(u)
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

func (op OpenMeteo) GetForecast(lat, lon string, tm time.Time) (weatheragent.Weather, error) {
	forecast, err := op.doRequest(lat, lon)
	if err != nil {
		return weatheragent.Weather{}, fmt.Errorf("getForecast(lat, lon string): %v", err)
	}

	if len(forecast.Hourly.Time) == 0 {
		return weatheragent.Weather{}, fmt.Errorf("getForecast(lat, lon string): no forecast data")
	}

	timeIdx, err := op.findNearestTime(tm, forecast.Hourly.Time)
	if err != nil {
		return weatheragent.Weather{}, fmt.Errorf("getForecast(lat, lon string): %v", err)
	}

	return weatheragent.Weather{
		LocalTime:      forecast.Hourly.Time[timeIdx].Time,
		Temperature:    forecast.Hourly.Temp_2m[timeIdx],
		WindSpeed:      forecast.Hourly.Wind_10m[timeIdx],
		WindGusts:      forecast.Hourly.Gusts_10m[timeIdx],
		WindDirection:  forecast.Hourly.WindDir_10m[timeIdx],
		CloudCoverLow:  forecast.Hourly.CloudCoverLow[timeIdx],
		CloadCoverMid:  forecast.Hourly.CloudCoverMid[timeIdx],
		CloadCoverHigh: forecast.Hourly.CloudCoverHigh[timeIdx],
	}, nil
}
