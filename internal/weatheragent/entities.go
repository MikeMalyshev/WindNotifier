package weatheragent

import (
	"time"

	"github.com/labstack/echo"
)

type WeatherAgent struct {
	echoInstance     *echo.Echo
	currentProvider  CurrentWeatherProvider
	forecastProvider ForecastWeatherProvider
}

type AgentResponse struct {
	Current  Weather `json:"current"`
	Forecast Weather `json:"forecast"`
}

type Weather struct {
	LocalTime     time.Time `json:"local_time"`
	Temperature   float64   `json:"temperature"`
	WindSpeed     float64   `json:"wind_speed"`
	WindDirection float64   `json:"wind_direction"`
}

type ForecastWeatherProvider interface {
	GetForecast(lon, lat string, t time.Time) (Weather, error)
}

type CurrentWeatherProvider interface {
	GetCurrent(lon, lat string) (Weather, error)
}
