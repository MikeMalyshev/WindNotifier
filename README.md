# WindNotifier
Back-end for telegram bot WindNotifier

## ТЗ:
Приложение должно уведомлять пользователя о наличии ветра на заданном споте, определяемом по широте и долготе:
1. По прогнозу за сутки до ожидаемой даты;
2. По прогнозу и по показаниям ближайших метеостанций на текущий момент времени.

### Возможности пользователя:
1. Задавать положение интересующего спота поиском по названию или по ближайшему к указанной геопозиции;
2. Получать положение заданного спота;
3. Получать прогноз по споту и актуальное состояния ветра по ближайшим метеостанциям;
4. Изменять время получения автоматических уведомлений;
5. Получать описание спота;
6. Вносить предложения по редактированию описания;
Дополнительно можно организовать чаты по отдельным спотам.

### Ограничительные условия:
Необходимо обойтись существующими бесплатными api по получению прогноза и информаций с метеостанций,
поэтому запросы по необходимым локациям буду отправляться строго определенное количество раз в день,
храниться в базе данных и уже оттуда отправляться при необходимости пользователю.

### Компоненты:
- Telegram Api (сначала teleBot, потом попробую сделать свою реализацию);
- Weather Api, OpenMeteo Api, WeatherStack Api (наверное сделаю отдельным сервисом);
- PostgreSql для хранения информации о польхователях и информации, полученной от погодных Api:
  - Таблица пользователей и их настроек;
  - Таблица локаций;
  - Таблица прогнозов;
- Docker / DockerCompose.
