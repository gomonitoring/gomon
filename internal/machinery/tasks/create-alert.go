package tasks

import (
	"github.com/gomonitoring/http-server/internal/database"
	"github.com/gomonitoring/http-server/internal/model"
	log "github.com/sirupsen/logrus"
)

func CreateAlert(id uint, time int64) error {
	db, _ := database.NewDB()
	err := db.Create(&model.Alert{
		Time:  time,
		UrlID: id,
	}).Error
	if err != nil {
		log.Fatalln("Could not create alert", err)
	}
	log.Infoln("An alert created")
	return nil
}
