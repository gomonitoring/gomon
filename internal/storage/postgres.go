package storage

import (
	"context"

	"github.com/gomonitoring/http-server/internal/model"
	"gorm.io/gorm"
)

type PostgresDB struct {
	db *gorm.DB
}

func NewPostgresDBStorage(db *gorm.DB) PostgresDB {
	return PostgresDB{
		db: db,
	}
}

func (p PostgresDB) Save(ctx context.Context, u model.User) error {
	err := p.db.Create(&u).Error
	if err != nil {
		return ErrorUserDuplicate
	}

	return nil
}

func (p PostgresDB) LoadByUserPass(ctx context.Context, username string, password string) (model.User, error) {
	var user model.User

	p.db.Find(&user, "username = ? AND password = ?", username, password)

	if user.Username == "" {
		return model.User{}, ErrorUserNotFound
	}

	return user, nil
}

func (p PostgresDB) SaveUrl(context.Context, model.Url) error {
	return nil
}

func (p PostgresDB) GetUrl(context.Context, string) (model.Url, error) {
	return model.Url{}, nil
}

func (p PostgresDB) GetUserUrls(context.Context) ([]model.Url, error) {
	return []model.Url{}, nil
}

func (p PostgresDB) GetUrlStats(context.Context, model.Url) ([]model.Call, error) {
	return []model.Call{}, nil
}
