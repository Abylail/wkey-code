package job

import (
	"wkey-core/src/controllers"
	"wkey-core/src/job/job_server"
)

func RunAsync(apiControllers *controllers.ApiControllers) {
	client := job_server.Create(apiControllers)
	client.RunAsync()
}

func Run(apiControllers *controllers.ApiControllers) {
	client := job_server.Create(apiControllers)
	client.Run()
}
