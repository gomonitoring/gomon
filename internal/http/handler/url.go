package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gomonitoring/http-server/internal/storage"
)

type Url struct {
	Storage storage.Url
}

func (u Url) RegisterUrl(c *fiber.Ctx) error {
	return nil
}

func (u Url) GetUrls(c *fiber.Ctx) error {
	return nil
}

func (u Url) GetUrlStats(c *fiber.Ctx) error {
	return nil
}

func (u Url) Register(g fiber.Router) {
	g.Post("/register-url", u.RegisterUrl)
	g.Get("/urls", u.GetUrls)
	g.Post("/stats", u.GetUrlStats)
}
