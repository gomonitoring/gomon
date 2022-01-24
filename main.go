package main

import (
	"log"
	"os"

	jwtware "github.com/gofiber/jwt/v3"
	"github.com/gomonitoring/http-server/internal/database"
	"github.com/gomonitoring/http-server/internal/http/handler"
	"github.com/gomonitoring/http-server/internal/storage"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// create the db and pass it to the handler
	db, err := database.NewDB()
	if err != nil {
		log.Fatalf("database connection failed %s", err)
	}

	hu := handler.User{
		Storage: storage.NewPostgresDBStorage(db),
	}

	app := fiber.New()
	// read secret from configmap
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))
	userg := app.Group("/user")
	hu.Register(userg)

	// read port from configmap, ":8080"
	if err := app.Listen(os.Getenv("LISTEN_PORT")); err != nil {
		log.Println("cannot start the server")
	}
}
