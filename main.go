package main

import (
	"log"

	jwtware "github.com/gofiber/jwt/v3"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// create the db and pass it to the handler

	app := fiber.New()
	// read secret from configmap
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))
	userg := app.Group("/user")
	hs.Register(userg)

	// read port from configmap
	if err := app.Listen(":8080"); err != nil {
		log.Println("cannot start the server")
	}
}
