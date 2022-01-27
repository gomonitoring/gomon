package model

import (
	"gorm.io/gorm"
)

type Url struct {
	gorm.Model
	Name      string
	Url       string
	Threshold string
	Username  string
	ResetTime int
}
