package storage

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gomonitoring/http-server/internal/http/request"
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

func (p PostgresDB) SaveUser(ctx context.Context, u model.User) error {
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

func (p PostgresDB) SaveUrl(ctx context.Context, req request.Url, username string) (model.Url, error) {
	var user model.User
	p.db.Where("username = ?", username).Find(&user)
	maxUserUrlCount := os.Getenv("MAX_URL_COUNT")
	count, err := strconv.Atoi(maxUserUrlCount)
	if err != nil {
		return model.Url{}, err
	}
	if user.UrlCount+1 > count {
		return model.Url{}, fmt.Errorf("User reached max url count")
	}

	url := model.Url{
		Name:      req.Name,
		Url:       req.Url,
		Threshold: req.Threshold,
		Username:  username,
		ResetTime: int(time.Now().Unix()),
	}
	er := p.db.Create(&url).Error
	if er != nil {
		return model.Url{}, err
	}
	user.UrlCount += 1
	p.db.Save(&user)
	return url, nil
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
