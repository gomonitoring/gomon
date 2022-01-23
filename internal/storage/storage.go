package storage

import (
	"context"

	"github.com/gomonitoring/http-server/internal/model"
)

type Storage interface {
	Save(context.Context, model.User) error
	LoadUserPass(context.Context, string, string) (model.User, error)
}
