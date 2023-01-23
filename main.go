package main

import (
	"wkey-core/src/api"
	"wkey-core/src/controllers"
	"wkey-core/src/definition"
	"wkey-core/src/events"
	"wkey-core/src/job"
)

func main() {
	definition.Init()
	logger := definition.Logger

	// инициализация ивентов
	apiEvents, err := events.Get()
	if err != nil {
		logger.Fatal(err, "Initializing events error", "Application")
	}

	// контроллеры
	apiControllers := controllers.Get(apiEvents)

	// запуск джобов
	job.RunAsync(apiControllers)

	// запуск сервера
	api.StartServer(apiControllers, apiEvents)
}
