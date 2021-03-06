package storage

import (
	"context"
	"time"

	"github.com/gomonitoring/gomon/internal/http/request"
	"github.com/gomonitoring/gomon/internal/model"
	"github.com/gomonitoring/gomon/internal/settings"
	"github.com/gomonitoring/gomon/internal/utils"
	"gorm.io/gorm"
)

type PostgresDB struct {
	db *gorm.DB
}

func NewPostgresDBStorage(db *gorm.DB) *PostgresDB {
	return &PostgresDB{
		db: db,
	}
}

func (p *PostgresDB) SaveUser(ctx context.Context, req *request.User) (*model.User, error) {
	hashedPass, err := utils.GetSha256(req.Password)
	if err != nil {
		return &model.User{}, err
	}
	u := model.User{
		Username: req.Username,
		Password: hashedPass,
	}
	er := p.db.Create(&u).Error
	if er != nil {
		return &model.User{}, ErrorUserDuplicate
	}

	return &u, nil
}

func (p *PostgresDB) LoadByUserPass(ctx context.Context, username string, password string) (*model.User, error) {
	hashedPass, err := utils.GetSha256(password)
	if err != nil {
		return &model.User{}, err
	}

	var user model.User

	p.db.Find(&user, "username = ? AND password = ?", username, hashedPass)

	if user.Username == "" {
		return &model.User{}, ErrorUserNotFound
	}

	return &user, nil
}

func (p *PostgresDB) SaveUrl(ctx context.Context, req *request.Url, username string) (*model.Url, error) {
	var user model.User
	p.db.Where("username = ?", username).Find(&user)
	maxUrl := settings.MaxUrlPerUser
	if user.UrlCount+1 > maxUrl {
		return &model.Url{}, ErrorMaxUrlCount
	}

	resetTime, e := time.ParseDuration(settings.DefaultResetTime)
	if e != nil {
		return &model.Url{}, e
	}

	url := model.Url{
		Name:      req.Name,
		Url:       req.Url,
		Threshold: req.Threshold,
		User:      user,
		ResetTime: resetTime.Nanoseconds(),
	}
	er := p.db.Create(&url).Error
	if er != nil {
		return &model.Url{}, er
	}
	user.UrlCount += 1
	p.db.Save(&user)
	return &url, nil
}

func (p *PostgresDB) GetUserUrls(ctx context.Context, username string) ([]model.Url, error) {
	var urls []model.Url
	p.db.Preload("User").Where("user_id = (?)", p.db.Table("users").Select("id").Where("username = ?", username)).Find(&urls)
	return urls, nil
}

func (p *PostgresDB) GetUrlStats(ctx context.Context, urlName string, username string) ([]model.Call, error) {
	now := time.Now().Unix()
	yesterday := time.Now().Add(-24 * time.Hour).Unix()
	var calls []model.Call
	p.db.Preload("Url").Preload("Url.User").Where("url_id = (?) AND time BETWEEN ? AND ?", p.db.Table("urls").Select("id").
		Where("user_id = (?) AND name = ?", p.db.Table("users").Select("id").
			Where("username = ?", username), urlName), yesterday, now).Find(&calls)
	return calls, nil
}

func (p *PostgresDB) GetAlerts(ctx context.Context, urlName string, username string) ([]model.Alert, error) {
	var alerts []model.Alert
	p.db.Preload("Url").Preload("Url.User").Where("url_id = (?)", p.db.Table("urls").Select("id").
		Where("user_id = (?) AND name = ?", p.db.Table("users").Select("id").
			Where("username = ?", username), urlName)).Find(&alerts)
	return alerts, nil
}

func (p *PostgresDB) GetUrlsToCall() ([]model.Url, error) {
	var urls []model.Url
	p.db.Find(&urls)
	return urls, nil
}

func (p *PostgresDB) SaveAlert(time int64, urlId uint) error {
	err := p.db.Create(&model.Alert{
		Time:  time,
		UrlID: urlId,
	}).Error
	return err
}

func (p *PostgresDB) SaveCallResults(callResults []model.CallUrlResult) error {
	calls := make([]model.Call, len(callResults))
	for i, c := range callResults {
		calls[i] = model.Call{
			Time:       c.Time,
			StatusCode: c.StatusCode,
			UrlID:      c.Id,
			Successful: c.StatusCode < 500 && c.StatusCode > 99,
		}
	}
	return p.db.Create(&calls).Error
}
