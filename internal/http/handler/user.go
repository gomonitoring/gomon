package handler

import (
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
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
		log.Printf("cannot load user data %s", err)

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

func (u User) Login(c *fiber.Ctx) error {
	req := new(request.User)

	if err := c.BodyParser(req); err != nil {
		log.Printf("cannot load user data %s", err)

		return fiber.ErrBadRequest
	}

	if err := req.Validate(); err != nil {
		log.Printf("cannot validate user data %s", err)

		return fiber.ErrBadRequest
	}

	user, err := u.Storage.LoadByUserPass(c.Context(), req.Username, req.Password)
	if err != nil {
		if errors.Is(err, storage.ErrorUserNotFound) {
			return fiber.ErrNotFound
		}

		log.Printf("cannot load user %s", err)

		return fiber.ErrInternalServerError
	}

	claims := jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(), // read expire from configmap
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// read the secret from configmap
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"token": t})
}

func (u User) Register(g fiber.Router) {
	g.Post("/login", u.Login)
	g.Post("/signup", u.SignUp)
}
