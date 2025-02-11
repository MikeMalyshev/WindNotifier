package openmeteo

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDoRequest(t *testing.T) {
	op := New()
	forecast, err := op.doRequest("51.5074", "0.1278")
	assert.NoError(t, err)
	fmt.Println(forecast)
}

func TestGetForecast(t *testing.T) {
	op := New()
	forecast, err := op.GetForecast("51.5074", "0.1278", time.Now())
	assert.NoError(t, err)
	fmt.Println(forecast)
}
