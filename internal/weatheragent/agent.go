package weatheragent

import (
	"time"

	"github.com/labstack/echo"
)

func New(awp ActualWeatherProvider, fwp ForecastWeatherProvider) *WeatherAgent {
	return &WeatherAgent{
		echoInstance:     echo.New(),
		actualProvider:   awp,
		forecastProvider: fwp,
	}
}

func (wa *WeatherAgent) defaultHandler(c echo.Context) error {
	lon := c.FormValue("lon")
	lat := c.FormValue("lat")

	if lon == "" || lat == "" {
		return c.String(400, "lon and lat must be provided")
	}

	resp := AgentResponse{
		Actual:   Weather{},
		Forecast: Weather{},
	}

	if c.FormValue("actual") == "true" {
		if wa.actualProvider == nil {
			return c.String(500, "currentProvider is not set")
		}
		actual, err := wa.actualProvider.GetActual(lon, lat, time.Now())
		if err != nil {
			return c.String(500, err.Error())
		}
		resp.Actual = actual
	}

	if c.FormValue("forecast") == "true" {
		if wa.forecastProvider == nil {
			return c.String(500, "forecastProvider is not set")
		}
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
