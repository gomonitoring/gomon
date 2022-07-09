package tasks

import (
	"github.com/gomonitoring/http-server/internal/machinery/worker"
	"github.com/gomonitoring/http-server/internal/storage"
	log "github.com/sirupsen/logrus"
)

func CreateAlert(id uint, time int64) error {
	var st storage.LocalWorker = worker.GetLocalWorkerStorage()
	err := st.SaveAlert(time, id)
	if err != nil {
		log.Fatalln("could not create alert", err)
	}
	log.Infoln("an alert created")
	return nil
}
