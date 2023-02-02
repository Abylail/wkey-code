package migrate_controller

import (
	"wkey-core/src/controllers/controller"
	"wkey-core/src/events"
	"wkey-core/src/events/migrate_event"
)

type Controller struct {
	controller.Base
	migrate *migrate_event.Event
}

func Create(apiEvents *events.ApiEvents) *Controller {
	return &Controller{
		migrate: apiEvents.Migrate,
	}
}
