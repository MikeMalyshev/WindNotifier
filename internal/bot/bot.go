package bot

import (
	"log"
	"os"
	"time"

	tele "gopkg.in/telebot.v4"
)

type TelegramID int64

type Storage interface {
	AddUser(u User) error
	UpdateUser(u User) error
	GetUser(id TelegramID) (User, error)
}

type WeatherInfo interface {
	Forecast(loc Location) (string, error)
	Current(loc Location) (string, error)
	CurrentAndForecast(loc Location) (string, string, error)
}

type WindNotifier struct {
	bot         *tele.Bot
	storage     Storage
	weatherInfo WeatherInfo
}

func Create(stor Storage) (WindNotifier, error) {
	settings := tele.Settings{
		Token:  os.Getenv("windnotifier_token"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}
	bot, err := tele.NewBot(settings)
	if err != nil {
		log.Fatal(err)
		return WindNotifier{}, err
	}

	w := WindNotifier{bot: bot, storage: stor}
	bot.Handle(tele.OnText, w.textHandler)
	bot.Handle(tele.OnLocation, w.locationHandler)

	return w, nil
}

func (w *WindNotifier) Start() {
	w.bot.Start()
}
