package storage

import (
	"context"
	"fmt"

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
		return fmt.Errorf("postgres insert failed %w", err)
	}

	return nil
}

func (p PostgresDB) LoadByUserPass(ctx context.Context, username string, password string) (model.User, error) {
	var user model.User

	p.db.Find(&user, "username = ? AND password = ?", username, password)

	if user.ID == "" {
		return model.User{}, fmt.Errorf("reading from postgres failed")
	}

	return user, nil
}
