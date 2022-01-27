package server

import (
	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	"github.com/gomonitoring/http-server/internal/settings"
)

func GetMachineryServer() *machinery.Server {
	taskserver, err := machinery.NewServer(&config.Config{
		Broker:        settings.MachineryBroker,
		ResultBackend: settings.MachineryResultBackend,
		DefaultQueue:  "local",
	})
	if err != nil {
		panic(err.Error())
	}

	taskserver.RegisterTasks(map[string]interface{}{
		// TODO register tasks
	})

	return taskserver
}
