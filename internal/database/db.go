package database

import (
	"fmt"

	"github.com/gomonitoring/http-server/internal/model"
	"github.com/gomonitoring/http-server/internal/settings"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		settings.DBHost,
		settings.DBPort,
		settings.DBUser,
		settings.DBPass,
		settings.DBName)

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic("failed to connect database")
	}

	log.Println("connection opened to database")
	db.AutoMigrate(&model.User{}, &model.Url{}, &model.Call{}, &model.Alert{})
	log.Println("database migrated")

	return db, nil
}
