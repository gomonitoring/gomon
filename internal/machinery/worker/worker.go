package worker

import (
	"github.com/RichardKnop/machinery/v1"
	"github.com/gomonitoring/http-server/internal/settings"
)

func StartMonitoringWorker(taskserver *machinery.Server) error {

	worker := taskserver.NewCustomQueueWorker("monitoring_worker", settings.MonitoringWorkerConcurrency, "monitoring")
	if err := worker.Launch(); err != nil {
		return err
	}
	return nil
}

func StartLocalWorker(taskserver *machinery.Server) error {

	worker := taskserver.NewCustomQueueWorker("local_worker", settings.LocalWorkerConcurrency, "local")
	if err := worker.Launch(); err != nil {
		return err
	}
	return nil
}
