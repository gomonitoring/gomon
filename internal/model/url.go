package model

import (
	"gorm.io/gorm"
)

type Url struct {
	gorm.Model
	Name      string
	Url       string
	Threshold string
	User      User
	UserID    uint
	ResetTime int64
}
