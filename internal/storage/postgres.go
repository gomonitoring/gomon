package storage

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/gomonitoring/http-server/internal/http/request"
	"github.com/gomonitoring/http-server/internal/model"
	"github.com/gomonitoring/http-server/internal/utils"
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

func (p PostgresDB) SaveUser(ctx context.Context, req request.User) (model.User, error) {
	hashedPass, err := utils.GetSha256(req.Password)
	if err != nil {
		return model.User{}, err
	}
	u := model.User{
		Username: req.Username,
		Password: hashedPass,
	}
	er := p.db.Create(&u).Error
	if er != nil {
		return model.User{}, ErrorUserDuplicate
	}

	return u, nil
}

func (p PostgresDB) LoadByUserPass(ctx context.Context, username string, password string) (model.User, error) {
	hashedPass, err := utils.GetSha256(password)
	if err != nil {
		return model.User{}, err
	}

	var user model.User

	p.db.Find(&user, "username = ? AND password = ?", username, hashedPass)

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
		return model.Url{}, ErrorMaxUrlCount
	}

	url := model.Url{
		Name:      req.Name,
		Url:       req.Url,
		Threshold: req.Threshold,
		User:      user,
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

func (p PostgresDB) GetUserUrls(ctx context.Context, username string) ([]model.Url, error) {
	var urls []model.Url
	p.db.Preload("User").Where("user_id = (?)", p.db.Table("users").Select("id").Where("username = ?", username)).Find(&urls)
	return urls, nil
}

func (p PostgresDB) GetUrlStats(context.Context, model.Url) ([]model.Call, error) {
	return []model.Call{}, nil
}
