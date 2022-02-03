package tasks

import (
	"github.com/gomonitoring/http-server/internal/database"
	"github.com/gomonitoring/http-server/internal/model"
)

func CreateAlert(id uint, time int64) error {
	db, _ := database.NewDB()
	db.Create(&model.Alert{
		Time:  time,
		UrlID: id,
	})
	return nil
}
