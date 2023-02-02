package api

import (
	"wkey-core/src/controllers"
	"wkey-core/src/definition"
	"wkey-core/src/events"
)

func StartServer(apiControllers *controllers.ApiControllers, apiEvents *events.ApiEvents) {
	server := definition.Server
	gateway := definition.Gateway
	config := definition.Config.Server

	// проставлять роуты
	setRoutes(server, gateway, apiControllers, apiEvents)

	// проставлять миддлвейры
	setMiddlewares(server)

	// запуск сервера
	var port string
	if definition.Config.Primary {
		port = config.Port.Primary
	} else {
		port = config.Port.Secondary
	}
	server.Logger.Fatal(server.Start(port))
}
