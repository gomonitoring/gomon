package worker

import (
	log "github.com/sirupsen/logrus"

	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/tasks"
	"github.com/gomonitoring/gomon/internal/database"
	"github.com/gomonitoring/gomon/internal/settings"
	"github.com/gomonitoring/gomon/internal/storage"
)

var postgresStorage storage.LocalWorker

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
		Name:       "find_urls_to_call",
		RoutingKey: "local",
	}
	err := taskserver.RegisterPeriodicTask(settings.CallUrlsSchedule, "find_urls_to_call", signature)
	return err
}

func GetLocalWorkerStorage() storage.LocalWorker {
	if postgresStorage == nil {
		db, err := database.NewDB()
		if err != nil {
			log.Fatalf("database connection failed", err)
		}
		postgresStorage = storage.NewPostgresDBStorage(db)
	}
	return postgresStorage
}
