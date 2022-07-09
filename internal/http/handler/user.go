package handler

import (
	"errors"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gomonitoring/http-server/internal/http/request"
	"github.com/gomonitoring/http-server/internal/settings"
	"github.com/gomonitoring/http-server/internal/storage"
)

type User struct {
	Storage storage.User
}

func (u *User) SignUp(c *fiber.Ctx) error {
	req := new(request.User)

	if err := c.BodyParser(req); err != nil {
		log.Infoln("cannot load user data ", err)

		return fiber.ErrBadRequest
	}

	if err := req.Validate(); err != nil {
		log.Infoln("cannot validate user data ", err)

		return fiber.ErrBadRequest
	}

	user, err := u.Storage.SaveUser(c.Context(), req)
	if err != nil {
		if errors.Is(err, storage.ErrorUserDuplicate) {
			return fiber.NewError(http.StatusBadRequest, "user already exists")
		}

		log.Infoln("cannot save user ", err)

		return fiber.ErrInternalServerError
	}

	return c.Status(http.StatusCreated).JSON(user)
}

func (u *User) Login(c *fiber.Ctx) error {
	req := new(request.User)

	if err := c.BodyParser(req); err != nil {
		log.Infoln("cannot load user data ", err)

		return fiber.ErrBadRequest
	}

	if err := req.Validate(); err != nil {
		log.Infoln("cannot validate user data ", err)

		return fiber.ErrBadRequest
	}

	user, err := u.Storage.LoadByUserPass(c.Context(), req.Username, req.Password)
	if err != nil {
		if errors.Is(err, storage.ErrorUserNotFound) {
			return fiber.ErrNotFound
		}

		log.Infoln("cannot load user ", err)

		return fiber.ErrInternalServerError
	}

	claims := jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(settings.JWTSecret))
	if err != nil {
		log.Errorln("can not generate jwt token ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"token": t})
}

func (u *User) Register(g fiber.Router) {
	g.Post("/login", u.Login)
	g.Post("/signup", u.SignUp)
}
