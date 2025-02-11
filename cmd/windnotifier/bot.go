package main

import "github.com/MikeMalyshev/WindNotifier/internal/bot"

func main() {
	bot, err := bot.Create(nil)
	if err != nil {
		panic(err)
	}
	bot.Start()
}
