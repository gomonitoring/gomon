package main

import (
	"os"

	log "github.com/sirupsen/logrus"

	jwtware "github.com/gofiber/jwt/v3"
	"github.com/gomonitoring/http-server/internal/database"
	"github.com/gomonitoring/http-server/internal/http/handler"
	"github.com/gomonitoring/http-server/internal/machinery/tasks"
	"github.com/gomonitoring/http-server/internal/machinery/worker"
	"github.com/gomonitoring/http-server/internal/settings"
	"github.com/gomonitoring/http-server/internal/storage"

	"github.com/RichardKnop/machinery/v1"
	"github.com/gofiber/fiber/v2"
	"github.com/urfave/cli"
)

var (
	app        *cli.App
	taskserver *machinery.Server
)

func init() {
	app = cli.NewApp()
	taskserver = tasks.GetMachineryServer()
}

func main() {
	app.Commands = []cli.Command{
		{
			Name:  "server",
			Usage: "Start http server.",
			Action: func(c *cli.Context) {
				// create the db and pass it to the handler
				db, err := database.NewDB()
				if err != nil {
					log.Fatalf("database connection failed %s", err)
				}

				hu := handler.User{
					Storage: storage.NewPostgresDBStorage(db),
				}

				hurl := handler.Url{
					Storage: storage.NewPostgresDBStorage(db),
				}

				ha := handler.Alert{
					Storage: storage.NewPostgresDBStorage(db),
				}

				app := fiber.New()
				userg := app.Group("/user")
				hu.Register(userg)
				// read secret from configmap
				app.Use(jwtware.New(jwtware.Config{
					SigningKey: []byte(settings.JWTSecret),
				}))

				urlg := app.Group("/url")
				hurl.Register(urlg)

				alertg := app.Group("/alert")
				ha.Register(alertg)

				// read port from configmap, ":8080"
				if err := app.Listen(":8080"); err != nil {
					log.Println("cannot start the server")
				}
			},
		}, {
			Name:  "local_worker",
			Usage: "Start local machinery queue.",
			Action: func(c *cli.Context) {
				worker.StartLocalWorker(taskserver)
			},
		}, {
			Name:  "monitoring_worker",
			Usage: "Start monitoring machinery worker.",
			Action: func(c *cli.Context) {
				worker.StartMonitoringWorker(taskserver)
			},
		},
	}
	app.Run(os.Args)
}
