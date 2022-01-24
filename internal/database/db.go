package database

import (
	"fmt"
	"os"

	"github.com/gomonitoring/http-server/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() (*gorm.DB, error) {
	// Connection URL to connect to Postgres Database, read from configmap
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"))

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")
	db.AutoMigrate(&model.User{})
	fmt.Println("Database Migrated")

	return db, nil
}
