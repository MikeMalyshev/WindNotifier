package bot

import (
	"fmt"

	tele "gopkg.in/telebot.v4"
)

type EndPoint string
type WaitingInfo int

const (
	WaNothing WaitingInfo = iota
	WaNewLocation
	WaNewLocationConfirm
)

const (
	CurrentWind     EndPoint = "/wind"
	CurrentLocation EndPoint = "/loc"
	SetNewLocation  EndPoint = "/newloc"

	Confirm EndPoint = "/yes"
	Decline EndPoint = "/no"
)

type User struct {
	ID      TelegramID
	Name    string
	State   WaitingInfo
	Loc     Location
	LastMsg string
}

var UserList map[int64]User = make(map[int64]User, 0)

func (w *WindNotifier) locationHandler(c tele.Context) error {
	user := UserList[c.Sender().ID]
	switch user.State {
	case WaNewLocation:
		loc, err := FindLocationByCoords([2]float32{c.Message().Location.Lat, c.Message().Location.Lng})
		if err != nil {
			c.Send(fmt.Sprintf("Неизвестная локация"))
		}
		w.storage.UpdateUser(user)
		c.Send(fmt.Sprintf("Установлено новое местоположение: %s", loc))

	default:
		user.State = WaNewLocationConfirm
		c.Send("Вы хотите изменить текущее местоположение?")
	}
	return nil
}

func (bot *WindNotifier) defaultTextHandler(c tele.Context) error {
	user := UserList[c.Sender().ID]
	switch EndPoint(c.Text()) {
	case CurrentWind:
		c.Send("Пока ветра нет, ждем ...")

	case CurrentLocation:
		err := c.Send(&tele.Location{Lat: 55, Lng: 55})
		if err != nil {
			fmt.Println(err)
		}
		c.Send(fmt.Sprintf("Текущее местоположение: %s", user.Loc))

	case SetNewLocation:
		c.Send("Введите новое местоположение")
		user.State = WaNewLocation
		UserList[c.Sender().ID] = user

	default:
		c.Send("Неизвестная команда")
	}
	return nil
}

func (w *WindNotifier) textHandler(c tele.Context) error {
	user := UserList[c.Sender().ID]
	switch user.State {
	case WaNothing:
		return w.defaultTextHandler(c)

	case WaNewLocation:
		loc, err := FindLocationByName(c.Text())
		if err == nil {
			user.Loc = loc
			user.State = WaNothing
			UserList[c.Sender().ID] = user
			c.Send(fmt.Sprintf("Местоположение было изменено на %s", loc))
		} else {
			c.Send("Неверное местоположение")
		}
	case WaNewLocationConfirm:
		switch EndPoint(c.Text()) {
		case Confirm:
			user.State = WaNothing
		case Decline:
			user.State = WaNothing
		default:
			c.Send("Ожидаем /yes or /no")
		}
	default:
		c.Send("Какой-то неверный статус пользователя")
	}

	return nil
}
