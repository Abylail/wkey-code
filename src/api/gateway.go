package api

import (
	"github.com/lowl11/lazy-gateway/gateway"
	"github.com/lowl11/lazy-gateway/lazyway"
	"wkey-core/src/definition"
	"wkey-core/src/middlewares"
)

func setGateway(router *gateway.Client) {
	setAuth(router)
	setStock(router)
}

func setAuth(router *gateway.Client) {
	config := definition.Config

	authClientRoute := lazyway.Route("/api/v1/auth").SetHosts(config.Hosts...).SetPort(":8084")
	authAdminRoute := lazyway.Route("/admin/api/v1/auth").SetHosts(config.Hosts...).SetPort(":8084")

	router.Route(authClientRoute)
	router.Route(authAdminRoute)
}

func setStock(router *gateway.Client) {
	config := definition.Config

	stockClientRoute := lazyway.Route("/api/v1/stock").SetHosts(config.Hosts...).SetPort(":8083")
	stockAdminRoute := lazyway.Route("/admin/api/v1/stock").SetHosts(config.Hosts...).SetPort(":8083").SetMiddleware(middlewares.AdminAuth)

	router.Route(stockClientRoute)
	router.Route(stockAdminRoute)
}
