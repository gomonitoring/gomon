package worker

import (
	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/tasks"
	"github.com/gomonitoring/http-server/internal/settings"
)

func StartMonitoringWorker(taskserver *machinery.Server) {
	worker := taskserver.NewCustomQueueWorker("monitoring_worker", settings.MonitoringWorkerConcurrency, "monitoring")
	if err := worker.Launch(); err != nil {
		panic(err.Error())
	}
}

func StartLocalWorker(taskserver *machinery.Server) {
	if err := registerPeriodicTasks(taskserver); err != nil {
		panic(err.Error())
	}
	worker := taskserver.NewCustomQueueWorker("local_worker", settings.LocalWorkerConcurrency, "local")
	if err := worker.Launch(); err != nil {
		panic(err.Error())
	}
}

func registerPeriodicTasks(taskserver *machinery.Server) error {
	signature := &tasks.Signature{
		Name: "find_urls_to_call",
	}
	err := taskserver.RegisterPeriodicTask("* * * * *", "find_urls_to_call", signature)
	return err
}
