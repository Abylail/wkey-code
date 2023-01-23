package static_controller

import "wkey-core/src/controllers/controller"

type Controller struct {
	controller.Base
}

func Create() *Controller {
	return &Controller{}
}
