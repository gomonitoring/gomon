package tasks

import (
	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	"github.com/gomonitoring/gomon/internal/settings"
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
		"call_url":          CallUrl,
		"collect_results":   CollectResults,
		"create_alert":      CreateAlert,
		"find_urls_to_call": FindUrlsToCall,
	})

	return taskserver
}
