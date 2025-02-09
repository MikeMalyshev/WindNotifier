package openmeteo

import (
	"fmt"
	"time"
)

type TimeISO8601 struct {
	time.Time
}
type ForecastObject struct {
	Time []TimeISO8601 `json:"time"`
	Data []float64     `json:"temperature_2m"`
}

type Forecast struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`

	Elevation float64 `json:"elevation"`
	GenTime   float64 `json:"generationtime_ms"`
	UtcOffset int     `json:"utc_offset_seconds"`
	Timezone  string  `json:"timezone"`

	Hourly      ForecastObject    `json:"hourly"`
	HourlyUnits map[string]string `json:"hourly_units"`
	Daily       ForecastObject    `json:"daily"`
	DailyUnits  map[string]string `json:"daily_units"`
}

type ErrorResp struct {
	Error  bool   `json:"error"`
	Reason string `json:"reason"`
}

func (t *TimeISO8601) UnmarshalJSON(data []byte) error {
	// Remove the quotes from the JSON string
	str := string(data)
	str = str[1 : len(str)-1]

	tt, err := time.Parse("2006-01-02T15:04", str)
	if err != nil {
		return fmt.Errorf("UnmarshalJSON: %v", err)
	}
	t.Time = tt
	return nil
}

func (t *TimeISO8601) String() string {
	return t.Format("2006-01-02T15:04")
}
