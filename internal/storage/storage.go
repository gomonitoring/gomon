package storage

import (
	"context"
	"errors"

	"github.com/gomonitoring/http-server/internal/http/request"
	"github.com/gomonitoring/http-server/internal/model"
)

var (
	ErrorUserNotFound  = errors.New("user with given username doesn't exist")
	ErrorUserDuplicate = errors.New("user with given username already exists")

	ErrorMaxUrlCount = errors.New("user reached max url count")
)

type User interface {
	SaveUser(context.Context, request.User) (model.User, error)
	LoadByUserPass(context.Context, string, string) (model.User, error)
}

type Url interface {
	SaveUrl(context.Context, request.Url, string) (model.Url, error)
	GetUrl(context.Context, string) (model.Url, error)
	GetUserUrls(context.Context) ([]model.Url, error)
	GetUrlStats(context.Context, model.Url) ([]model.Call, error)
}
