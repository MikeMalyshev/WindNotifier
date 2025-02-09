package openmeteo

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetWeather(t *testing.T) {
	forecast, err := getWeather("51.5074", "0.1278")
	assert.NoError(t, err)
	fmt.Println(forecast)
}
