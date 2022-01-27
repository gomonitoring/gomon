package model

import (
	"gorm.io/gorm"
)

type Url struct {
	gorm.Model
	Name      string
	Url       string
	Threshold int
	User      User
	UserID    int
	ResetTime int
}
