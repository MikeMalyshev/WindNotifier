package weatheragent

import (
	"time"

	"github.com/labstack/echo"
)

func New(cwp CurrentWeatherProvider, fwp ForecastWeatherProvider) *WeatherAgent {
	return &WeatherAgent{
		echoInstance:     echo.New(),
		currentProvider:  cwp,
		forecastProvider: fwp,
	}
}

func (wa *WeatherAgent) defaultHandler(c echo.Context) error {
	lon := c.FormValue("lon")
	lat := c.FormValue("lat")

	if lon == "" || lat == "" {
		return c.String(400, "lon and lat must be provided")
	}

	if wa.currentProvider == nil {
		return c.String(500, "currentProvider is not set")
	}

	if wa.forecastProvider == nil {
		return c.String(500, "forecastProvider is not set")
	}

	resp := AgentResponse{
		Current:  Weather{},
		Forecast: Weather{},
	}

	if c.FormValue("current") == "true" {
		current, err := wa.currentProvider.GetCurrent(lon, lat)
		if err != nil {
			return c.String(500, err.Error())
		}
		resp.Current = current
	}

	if c.FormValue("forecast") == "true" {
		var (
			forecast Weather
			err      error
		)

		ts := c.FormValue("time")
		if len(ts) == 0 {
			forecast, err = wa.forecastProvider.GetForecast(lon, lat, time.Now())
			if err != nil {
				return c.String(500, err.Error())
			}
		} else {
			tm, err := time.Parse(time.RFC3339, ts)
			if err != nil {
				return c.String(400, "time must be in RFC3339 format")
			}
			forecast, err = wa.forecastProvider.GetForecast(lon, lat, tm)
		}

		if err != nil {
			return c.String(500, err.Error())
		}
		resp.Forecast = forecast
	}
	return c.JSON(200, resp)
}

func (wa *WeatherAgent) Start() {
	wa.echoInstance.GET("/", wa.defaultHandler)
	wa.echoInstance.Start(":8181")
}
