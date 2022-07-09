package main

import (
	"context"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"

	jwtware "github.com/gofiber/jwt/v3"
	"github.com/gomonitoring/http-server/internal/database"
	"github.com/gomonitoring/http-server/internal/http/handler"
	"github.com/gomonitoring/http-server/internal/http/request"
	"github.com/gomonitoring/http-server/internal/machinery/tasks"
	"github.com/gomonitoring/http-server/internal/machinery/worker"
	"github.com/gomonitoring/http-server/internal/settings"
	"github.com/gomonitoring/http-server/internal/storage"
	"github.com/gomonitoring/http-server/internal/utils"

	"github.com/RichardKnop/machinery/v1"
	"github.com/gofiber/fiber/v2"
	"github.com/urfave/cli"
)

var (
	app        *cli.App
	taskserver *machinery.Server
)

type root struct {
	storage *storage.PostgresDB
}

func (r *root) registerRootUser(username string, pass string) {
	req := request.User{
		Username: username,
		Password: pass,
	}
	_, err := r.storage.SaveUser(context.Background(), &req)
	if err != nil {
		log.Errorln("can not register the root user")
	}
}

func (r *root) registerRootUrls(count int, username string) {
	var err error
	for i := 0; i < count; i++ {
		name := utils.RandSeq(6)
		req := request.Url{
			Name:      name,
			Url:       fmt.Sprintf("http://%s.com", name),
			Threshold: "2",
		}
		_, err = r.storage.SaveUrl(context.Background(), &req, username)
		if err != nil {
			log.Errorln("can not register root urls")
		}
	}
}

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
					log.Fatalf("database connection failed", err)
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
				app.Use(jwtware.New(jwtware.Config{
					SigningKey: []byte(settings.JWTSecret),
				}))

				urlg := app.Group("/url")
				hurl.Register(urlg)

				alertg := app.Group("/alert")
				ha.Register(alertg)

				if settings.AutoUrlPopulateCount > 0 {
					rt := root{
						storage: storage.NewPostgresDBStorage(db),
					}
					rt.registerRootUser("root", "12345678")
					rt.registerRootUrls(settings.AutoUrlPopulateCount, "root")
				}

				if err := app.Listen(":8080"); err != nil {
					log.Infoln("cannot start the server")
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
