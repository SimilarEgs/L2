// HTTP-сервер

// Реализовать HTTP-сервер для работы с календарем. В рамках задания необходимо работать строго со стандартной HTTP-библиотекой.

// В рамках задания необходимо:
// Реализовать вспомогательные функции для сериализации объектов доменной области в JSON.
// Реализовать вспомогательные функции для парсинга и валидации параметров методов /create_event и /update_event.
// Реализовать HTTP обработчики для каждого из методов API, используя вспомогательные функции и объекты доменной области.
// Реализовать middleware для логирования запросов

// Методы API:
// POST /create_event
// POST /update_event
// POST /delete_event
// GET /events_for_day
// GET /events_for_week
// GET /events_for_month

// Параметры передаются в виде www-url-form-encoded (т.е. обычные user_id=3&date=2019-09-09). В GET методах параметры передаются через queryString, в POST через тело запроса.
// В результате каждого запроса должен возвращаться JSON-документ содержащий либо {"result": "..."} в случае успешного выполнения метода, либо {"error": "..."} в случае ошибки бизнес-логики.

// В рамках задачи необходимо:
// Реализовать все методы.
// Бизнес логика НЕ должна зависеть от кода HTTP сервера.
// В случае ошибки бизнес-логики сервер должен возвращать HTTP 503.
// В случае ошибки входных данных (невалидный int например) сервер должен возвращать HTTP 400.
// В случае остальных ошибок сервер должен возвращать HTTP 500.
// Web-сервер должен запускаться на порту указанном в конфиге и выводить в лог каждый обработанный запрос.

// 1. Создадай модели бизнеслогики - готово
// 2. Валидация моедлей - готово
// 3. Создать функции API - в процессе
// 4. Мидлвейр для логирования запросов

package main

import (
	"dev11/config"
	"dev11/internal/app"
	"dev11/internal/repository"
	"log"
	"os"
)

func main() {

	// setup log file
	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)

	config, err := config.ParseConfig()
	if err != nil {
		log.Fatal(err)
	}

	events := repository.NewEventStorage()
	app.InitRoutes(events)

	server := app.NewServer(config)

	log.Println("[Info] starting server...")

	err = server.Run()
	if err != nil {
		log.Fatalf("[Error] %v\n", err)
	}

}
