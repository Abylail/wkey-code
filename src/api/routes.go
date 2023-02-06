package api

import (
	"github.com/labstack/echo/v4"
	"wkey-core/src/controllers"
	"wkey-core/src/events"
)

func setRoutes(server *echo.Echo, apiControllers *controllers.ApiControllers, apiEvents *events.ApiEvents) {
	// статичные методы
	server.GET("/health", apiControllers.Static.Health)
	server.RouteNotFound("*", apiControllers.Static.RouteNotFound)

	// эндпоинты
	//
}
