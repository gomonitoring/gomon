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
		ResetTime: time.Now().Unix(),
	}
	er := p.db.Create(&url).Error
	if er != nil {
		return model.Url{}, err
	}
	user.UrlCount += 1
	p.db.Save(&user)
	return url, nil
}

func (p PostgresDB) GetUserUrls(ctx context.Context, username string) ([]model.Url, error) {
	var urls []model.Url
	p.db.Preload("User").Where("user_id = (?)", p.db.Table("users").Select("id").Where("username = ?", username)).Find(&urls)
	return urls, nil
}

func (p PostgresDB) GetUrlStats(ctx context.Context, urlName string, username string) ([]model.Call, error) {
	//test
	call := mockCall()
	p.db.Create(&call)
	//
	now := time.Now().Unix()
	yesterday := time.Now().Add(-24 * time.Hour).Unix()
	var calls []model.Call
	p.db.Preload("Url").Preload("Url.User").Where("url_id = (?) AND time BETWEEN ? AND ?", p.db.Table("urls").Select("id").
		Where("user_id = (?) AND name = ?", p.db.Table("users").Select("id").
			Where("username = ?", username), urlName), yesterday, now).Find(&calls)
	return calls, nil
}

func (p PostgresDB) GetAlerts(ctx context.Context, urlName string, username string) ([]model.Alert, error) {
	//test
	alert := mockAlert()
	p.db.Create(&alert)
	//
	var alerts []model.Alert
	p.db.Preload("Url").Preload("Url.User").Where("url_id = (?)", p.db.Table("urls").Select("id").
		Where("user_id = (?) AND name = ?", p.db.Table("users").Select("id").
			Where("username = ?", username), urlName)).Find(&alerts)
	return alerts, nil
}

func mockCall() []model.Call {
	now := time.Now().Unix()
	calls := []model.Call{
		{Time: now - 1, UrlID: 1},
		{Time: now - 2, UrlID: 1},
		{Time: now - 3, UrlID: 1},
		{Time: 300000000, UrlID: 1},
	}
	return calls
}

func mockAlert() []model.Alert {
	now := time.Now().Unix()
	alerts := []model.Alert{
		{Time: now - 1, UrlID: 1},
		{Time: now - 2, UrlID: 1},
		{Time: now - 3, UrlID: 1},
		{Time: 300000000, UrlID: 1},
	}
	return alerts
}
