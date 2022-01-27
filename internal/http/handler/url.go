package handler

import (
	"errors"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gomonitoring/http-server/internal/http/request"
	"github.com/gomonitoring/http-server/internal/storage"
)

type Url struct {
	Storage storage.Url
}

func (u Url) RegisterUrl(c *fiber.Ctx) error {
	// extract username
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	username := claims["username"].(string)

	req := new(request.Url)
	if err := c.BodyParser(req); err != nil {
		log.Printf("cannot load url data %s", err)

		return fiber.ErrBadRequest
	}
	if err := req.Validate(); err != nil {
		log.Printf("cannot validate url data %s", err)

		return fiber.ErrBadRequest
	}

	url, err := u.Storage.SaveUrl(c.Context(), *req, username)
	if err != nil {
		if errors.Is(err, storage.ErrorMaxUrlCount) {
			return fiber.NewError(http.StatusBadRequest, "user reached max url count")
		}
		log.Printf("cannot save url %s", err)
		return fiber.ErrInternalServerError
	}

	return c.Status(http.StatusCreated).JSON(url)
}

func (u Url) GetUrls(c *fiber.Ctx) error {
	// extract username
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	username := claims["username"].(string)

	urls, err := u.Storage.GetUserUrls(c.Context(), username)
	if err != nil {
		log.Printf("cannot load urls %s", err)
		return fiber.ErrInternalServerError
	}

	return c.Status(http.StatusOK).JSON(urls)
}

func (u Url) GetUrlStats(c *fiber.Ctx) error {
	// extract username
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	username := claims["username"].(string)

	req := new(request.Stats)
	if err := c.BodyParser(req); err != nil {
		log.Printf("cannot load stats data %s", err)

		return fiber.ErrBadRequest
	}
	if err := req.Validate(); err != nil {
		log.Printf("cannot validate stats data %s", err)

		return fiber.ErrBadRequest
	}

	calls, err := u.Storage.GetUrlStats(c.Context(), req.Name, username)
	if err != nil {
		log.Printf("cannot load stats %s", err)
		return fiber.ErrInternalServerError
	}
	stats := map[string]int{"successes": 0, "failures": 0}
	for _, call := range calls {
		if call.Successful {
			stats["successes"] += 1
		} else {
			stats["failures"] += 1
		}
	}
	return c.Status(http.StatusOK).JSON(stats)
}

func (u Url) Register(g fiber.Router) {
	g.Post("/register-url", u.RegisterUrl)
	g.Get("/urls", u.GetUrls)
	g.Post("/stats", u.GetUrlStats)
}
