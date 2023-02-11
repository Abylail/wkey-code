package job_server

import (
	"github.com/lowl11/lazylog/layers"
	"wkey-core/src/definition"
)

func (server *Server) tasks() {
	logger := definition.Logger
	_ = logger

	//if _, err := server.scheduler.Cron("30 1 * * *").Do(server.migrate.CategoriesJob); err != nil {
	//	logger.Error(err, "Call categories migration error", layers.Job)
	//}

	// !!! uncomment next lines for localhost
	//if _, err := server.scheduler.Every(1).Hours().Do(server.migrate.CategoriesJob); err != nil {
	//	logger.Error(err, "Call categories migration error", layers.Job)
	//}

	if _, err := server.scheduler.Cron("35 1 * * *").Do(server.migrate.ProductsJob); err != nil {
		logger.Error(err, "Call products migration error", layers.Job)
	}

	if _, err := server.scheduler.Cron("00 15 * * *").Do(server.migrate.ProductsJob); err != nil {
		logger.Error(err, "Call products migration error", layers.Job)
	}

	// !!! uncomment next lines for localhost
	//if _, err := server.scheduler.Every(1).Hours().Do(server.migrate.ProductsJob); err != nil {
	//	logger.Error(err, "Call products migration error", layers.Job)
	//}
}
