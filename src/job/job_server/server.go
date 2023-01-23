package job_server

import (
	"github.com/go-co-op/gocron"
	"time"
	"wkey-core/src/controllers"
)

type Server struct {
	scheduler *gocron.Scheduler
}

func Create(apiControllers *controllers.ApiControllers) *Server {
	return &Server{
		scheduler: gocron.NewScheduler(time.UTC),
	}
}
