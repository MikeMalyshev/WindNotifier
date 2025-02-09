package main

import (
	"github.com/MikeMalyshev/WindNotifier/internal/weatheragent"
)

func main() {
	wa := weatheragent.New( /* TODO */ )
	wa.Start()
}
