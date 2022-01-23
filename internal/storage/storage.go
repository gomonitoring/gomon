package storage

import (
	"context"
	"errors"

	"github.com/gomonitoring/http-server/internal/model"
)

var (
	ErrorUserNotFound  = errors.New("user with given username doesn't exist")
	ErrorUserDuplicate = errors.New("student with given username already exists")
)

type User interface {
	Save(context.Context, model.User) error
	LoadByUserPass(context.Context, string, string) (model.User, error)
}
