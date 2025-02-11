package main

import (
	"github.com/MikeMalyshev/WindNotifier/internal/openmeteo"
	"github.com/MikeMalyshev/WindNotifier/internal/weatheragent"
)

func main() {
	openMeteoProvider := openmeteo.New()

	wa := weatheragent.New(nil, openMeteoProvider)
	wa.Start()
}
