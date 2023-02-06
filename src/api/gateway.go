package api

import (
	"github.com/lowl11/lazy-gateway/gateway"
	"github.com/lowl11/lazy-gateway/lazyway"
	"wkey-core/src/definition"
)

func setGateway(router *gateway.Client) {
	config := definition.Config

	stockRoute := lazyway.Route("/api/v1/stock").SetHosts(config.Hosts...)

	router.Route(stockRoute)
}
