package definition

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lowl11/lazy-gateway/gateway"
	"github.com/lowl11/lazy-gateway/lazyway"
	"wkey-core/src/middlewares"
)

var Server *echo.Echo
var Gateway *gateway.Client

// initServer создание объекта сервера
func initServer() {
	Server = echo.New()
	Gateway = lazyway.Echo(Server)

	// общие миддлы
	Server.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))
	Server.Use(middleware.Secure())
	Server.Use(middleware.RecoverWithConfig(middleware.DefaultRecoverConfig))
	Server.Use(middlewares.Timeout())
}
