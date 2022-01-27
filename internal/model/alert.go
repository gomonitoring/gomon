package model

import "gorm.io/gorm"

type Alert struct {
	gorm.Model
	Time  int64
	Url   Url
	UrlID uint
}
