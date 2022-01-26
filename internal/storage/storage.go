package storage

import (
	"context"
	"errors"

	"github.com/gomonitoring/http-server/internal/model"
)

var (
	ErrorUserNotFound  = errors.New("user with given username doesn't exist")
	ErrorUserDuplicate = errors.New("user with given username already exists")
)

type User interface {
	Save(context.Context, model.User) error
	LoadByUserPass(context.Context, string, string) (model.User, error)
}

type Url interface {
	SaveUrl(context.Context, model.Url) error
	GetUrl(context.Context, string) (model.Url, error)
	GetUserUrls(context.Context) ([]model.Url, error)
	GetUrlStats(context.Context, model.Url) ([]model.Call, error)
}
