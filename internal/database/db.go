package database

import (
	"fmt"

	"github.com/gomonitoring/http-server/internal/model"
	"github.com/gomonitoring/http-server/internal/settings"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() (*gorm.DB, error) {
	// Connection URL to connect to Postgres Database, read from configmap
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

	fmt.Println("Connection Opened to Database")
	db.AutoMigrate(&model.User{}, &model.Url{}, &model.Call{}, &model.Alert{})
	fmt.Println("Database Migrated")

	return db, nil
}
