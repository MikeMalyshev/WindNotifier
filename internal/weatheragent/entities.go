package weatheragent

import "time"

type Weather struct {
	LocalTime     time.Time
	Temperature   float64
	WindSpeed     float64
	WindDirection float64
}

type WeatherProvider interface {
	GetForecast(lon, lat float64, t time.Time) (Weather, error)
	GetCurrent(lon, lat string) (Weather, error)
}
