package main

import (
	"github.com/MikeMalyshev/WindNotifier/internal/weatheragent"
)

func main() {
	// Weather agent будет позже переписан на использование OpenMeteo
	wa := weatheragent.New(nil, nil)
	wa.Start()
}
