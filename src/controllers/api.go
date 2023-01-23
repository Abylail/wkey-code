package controllers

import (
	"wkey-core/src/controllers/static_controller"
	"wkey-core/src/events"
)

type ApiControllers struct {
	Static *static_controller.Controller
}

func Get(apiEvents *events.ApiEvents) *ApiControllers {
	return &ApiControllers{
		Static: static_controller.Create(),
	}
}
