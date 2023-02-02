package job_server

import (
	"github.com/go-co-op/gocron"
	"time"
	"wkey-core/src/controllers"
	"wkey-core/src/controllers/migrate_controller"
)

type Server struct {
	scheduler *gocron.Scheduler

	migrate *migrate_controller.Controller
}

func Create(apiControllers *controllers.ApiControllers) *Server {
	return &Server{
		scheduler: gocron.NewScheduler(time.UTC),
		migrate:   apiControllers.Migrate,
	}
}
