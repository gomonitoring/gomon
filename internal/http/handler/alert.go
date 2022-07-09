package handler

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gomonitoring/http-server/internal/http/request"
	"github.com/gomonitoring/http-server/internal/storage"
)

type Alert struct {
	Storage storage.Alert
}

func (a *Alert) GetAlert(c *fiber.Ctx) error {
	// extract username
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	username := claims["username"].(string)

	req := new(request.Alert)
	if err := c.BodyParser(req); err != nil {
		log.Printf("cannot load alert data", err)

		return fiber.ErrBadRequest
	}
	if err := req.Validate(); err != nil {
		log.Printf("cannot validate alert data", err)

		return fiber.ErrBadRequest
	}

	alerts, err := a.Storage.GetAlerts(c.Context(), req.UrlName, username)
	if err != nil {
		log.Printf("cannot load alerts", err)
		return fiber.ErrInternalServerError
	}
	return c.Status(http.StatusOK).JSON(alerts)
}

func (a *Alert) Register(g fiber.Router) {
	g.Post("/alerts", a.GetAlert)
}
