package handler

import (
	"errors"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gomonitoring/http-server/internal/http/request"
	"github.com/gomonitoring/http-server/internal/model"
	"github.com/gomonitoring/http-server/internal/storage"
)

type User struct {
	Storage storage.User
}

func (u User) SignUp(c *fiber.Ctx) error {
	req := new(request.User)

	if err := c.BodyParser(req); err != nil {
		log.Printf("cannot user student data %s", err)

		return fiber.ErrBadRequest
	}

	if err := req.Validate(); err != nil {
		log.Printf("cannot validate user data %s", err)

		return fiber.ErrBadRequest
	}

	user := model.User{
		Username: req.Username,
		Password: req.Password,
	}

	if err := u.Storage.Save(c.Context(), user); err != nil {
		if errors.Is(err, storage.ErrorUserDuplicate) {
			return fiber.NewError(http.StatusBadRequest, "user already exists")
		}

		log.Printf("cannot save user %s", err)

		return fiber.ErrInternalServerError
	}

	return c.Status(http.StatusCreated).JSON(user)
}
