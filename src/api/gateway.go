package api

import (
	"github.com/lowl11/lazy-gateway/gateway"
	"github.com/lowl11/lazy-gateway/lazyway"
	"wkey-core/src/definition"
)

func setGateway(router *gateway.Client) {
	config := definition.Config

	stockClientRoute := lazyway.Route("/api/v1/stock").SetHosts(config.Hosts...).SetPort(":8083")
	stockAdminRoute := lazyway.Route("/admin/api/v1/stock").SetHosts(config.Hosts...).SetPort(":8083")

	router.Route(stockClientRoute)
	router.Route(stockAdminRoute)
}
