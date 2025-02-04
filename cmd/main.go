package main

import "github.com/MikeMalyshev/WindNotifier/internal/bot"

func main() {
	bot, err := bot.Create()
	if err != nil {
		panic(err)
	}
	bot.StartBot()
}
