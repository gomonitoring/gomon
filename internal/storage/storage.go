package storage

import (
	"context"
	"errors"

	"github.com/gomonitoring/gomon/internal/http/request"
	"github.com/gomonitoring/gomon/internal/model"
)

var (
	ErrorUserNotFound  = errors.New("user with given username doesn't exist")
	ErrorUserDuplicate = errors.New("user with given username already exists")

	ErrorMaxUrlCount = errors.New("user reached max url count")
)

type User interface {
	SaveUser(context.Context, *request.User) (*model.User, error)
	LoadByUserPass(context.Context, string, string) (*model.User, error)
}

type Url interface {
	SaveUrl(context.Context, *request.Url, string) (*model.Url, error)
	GetUserUrls(context.Context, string) ([]model.Url, error)
	GetUrlStats(context.Context, string, string) ([]model.Call, error)
}

type Alert interface {
	GetAlerts(context.Context, string, string) ([]model.Alert, error)
}

type LocalWorker interface {
	GetUrlsToCall() ([]model.Url, error)
	SaveAlert(int64, uint) error
	SaveCallResults([]model.CallUrlResult) error
}
