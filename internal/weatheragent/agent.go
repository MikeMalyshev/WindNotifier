package weatheragent

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type WeatherAgent struct {
	echoInstance    *echo.Echo
	weatherProvider WeatherProvider
}

func New(wp WeatherProvider) *WeatherAgent {
	return &WeatherAgent{weatherProvider: wp}
}

func (wa *WeatherAgent) defaultHandler(c echo.Context) error {
	lon := c.FormValue("lon")
	lat := c.FormValue("lat")

	w, err := wa.weatherProvider.GetCurrent(lon, lat)
	if err != nil {
		return c.String(500, err.Error())
	}

	return c.JSON(200, w)
}

func (wa *WeatherAgent) Start() {
	wa.echoInstance = echo.New()
	wa.echoInstance.Use(middleware.Logger())

	wa.echoInstance.GET("/", wa.defaultHandler)
	wa.echoInstance.Start(":8181")
}
