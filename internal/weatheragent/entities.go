package weatheragent

import (
	"time"

	"github.com/labstack/echo"
)

type WeatherAgent struct {
	echoInstance     *echo.Echo
	actualProvider   ActualWeatherProvider
	forecastProvider ForecastWeatherProvider
}

type AgentResponse struct {
	Actual   Weather `json:"current"`
	Forecast Weather `json:"forecast"`
}

type Weather struct {
	LocalTime      time.Time `json:"local_time"`
	Temperature    float64   `json:"temperature"`
	WindSpeed      float64   `json:"wind_speed"`
	WindGusts      float64   `json:"wind_gusts"`
	WindDirection  float64   `json:"wind_direction"`
	CloudCoverLow  float64   `json:"cloud_cover_low"`
	CloadCoverMid  float64   `json:"cloud_cover_mid"`
	CloadCoverHigh float64   `json:"cloud_cover_high"`
}

type ForecastWeatherProvider interface {
	GetForecast(lon, lat string, t time.Time) (Weather, error)
}

type ActualWeatherProvider interface {
	GetActual(lon, lat string, t time.Time) (Weather, error)
}
